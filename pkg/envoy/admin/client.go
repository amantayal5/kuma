package admin

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/kumahq/kuma/pkg/core/ca"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	util_tls "github.com/kumahq/kuma/pkg/tls"
	xds_tls "github.com/kumahq/kuma/pkg/xds/envoy/tls"
)

type EnvoyAdminClient interface {
	PostQuit(dataplane *core_mesh.DataplaneResource) error
}

type AdminAddresser interface {
	AdminAddress(defaultAdminPort uint32) string
}

type envoyAdminClient struct {
	rm               manager.ResourceManager
	caManagers       ca.Managers
	clientCert       tls.Certificate
	defaultAdminPort uint32
}

func NewEnvoyAdminClient(rm manager.ResourceManager, caManagers ca.Managers, clientCertPath, clientKeyPath string, adminPort uint32) (EnvoyAdminClient, error) {
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	client := &envoyAdminClient{
		rm:               rm,
		caManagers:       caManagers,
		clientCert:       cert,
		defaultAdminPort: adminPort,
	}
	return client, nil
}

// Envoy admin API endpoint is secured in two possible ways
// 1) When mTLS on the mesh is disabled, we pass autogenerated self signed cert just to have TLS in place.
// 2) When mTLS on the mesh is enabled, we are protecting the endpoint with enabled mTLS backend.
//
// Regardless of which CA is used to protect Admin API endpoint, Envoy will always require certs from CP which are the same certs as DP server.
func (a *envoyAdminClient) buildHTTPClient(dataplane *core_mesh.DataplaneResource) (*http.Client, error) {
	caCertPool, err := a.caCertPoolOfMeshMTLS(dataplane.Meta.GetMesh())
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 3 * time.Second,
			TLSClientConfig: &tls.Config{
				VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
					if caCertPool == nil {
						// It means we that admin endpoint is protected with 1) option. We skip extra verification of cert
						return nil
					}
					// verify CA against the Mesh CA
					if err := util_tls.VerifyOnlyCA(caCertPool)(rawCerts, verifiedChains); err != nil {
						return err
					}

					// Verify SPIFFE to see if we are connecting to the right DP
					cert, _ := x509.ParseCertificate(rawCerts[0]) // ignore error because cert was parsed already
					dpSpiffe := xds_tls.ServiceSpiffeID(dataplane.Meta.GetMesh(), dataplane.Spec.GetIdentifyingService())
					for _, uri := range cert.URIs {
						if uri.String() == dpSpiffe {
							return nil
						}
					}
					return errors.Errorf("could not find expected URI SAN %s", dpSpiffe)
				},
				// We disable builtin verification because
				// 1) In first case, we don't have stable self-signed cert between instances of CP and we don't want to operate them.
				// 2) it expects hostname or IP in cert instead of SPIFFE URI SAN, so we cannot use builtin verification
				//
				// Also keep in mind that on this very moment we are not sending sensitive data to the DP.
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{a.clientCert},
			},
		},
		Timeout: 5 * time.Second,
	}
	return c, err
}

func (a *envoyAdminClient) caCertPoolOfMeshMTLS(mesh string) (*x509.CertPool, error) {
	meshRes := core_mesh.NewMeshResource()
	err := a.rm.Get(context.Background(), meshRes, core_store.GetByKey(mesh, model.NoMesh))
	if err != nil {
		return nil, err
	}
	backend := meshRes.GetEnabledCertificateAuthorityBackend()
	if backend == nil {
		return nil, nil
	}
	caManager, ok := a.caManagers[backend.Type]
	if !ok {
		return nil, errors.Errorf("cannot find CA Manager for type %s", backend.Type)
	}
	rootCerts, err := caManager.GetRootCert(context.Background(), mesh, backend)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	for _, certPEM := range rootCerts {
		block, _ := pem.Decode(certPEM)
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		certPool.AddCert(cert)
	}
	return certPool, nil
}

const (
	quitquitquit = "quitquitquit"
)

func (a *envoyAdminClient) PostQuit(dataplane *core_mesh.DataplaneResource) error {
	httpClient, err := a.buildHTTPClient(dataplane)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%s/%s", dataplane.AdminAddress(a.defaultAdminPort), quitquitquit)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	// Envoy will not send back any response, so do we not check the response
	response, err := httpClient.Do(request)
	if errors.Is(err, io.EOF) {
		return nil // Envoy may not respond correctly for this request because it already started the shut-down process.
	}
	if err != nil {
		return errors.Wrapf(err, "unable to send POST to %s", quitquitquit)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.Errorf("envoy response [%d %s] [%s]", response.StatusCode, response.Status, response.Body)
	}

	return nil
}

func ConfigDump(addresser AdminAddresser, defaultAdminAddress uint32) ([]byte, error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("http://%s/%s", addresser.AdminAddress(defaultAdminAddress), "config_dump")
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to send GET to %s", "config_dump")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("envoy response [%d %s] [%s]", response.StatusCode, response.Status, response.Body)
	}

	return io.ReadAll(response.Body)
}
