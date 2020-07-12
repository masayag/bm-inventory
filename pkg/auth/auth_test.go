package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/filanov/bm-inventory/restapi/operations/installer"
	"github.com/go-openapi/runtime/middleware"

	"github.com/filanov/bm-inventory/restapi"
	"github.com/sirupsen/logrus"

	"github.com/filanov/bm-inventory/client"
	client_installer "github.com/filanov/bm-inventory/client/installer"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		username string
	}{
		{
			name:     "username eran",
			username: "eran",
		},
		{
			name:     "username ronnie",
			username: "ronnie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// auth middleware that can check the username header
			authMiddleware := func() func(http.Handler) http.Handler {
				return func(inner http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						username := r.Header.Get("username")
						fmt.Printf("username %v", username)
						assert.NotNil(t, username)
						assert.Equal(t, tt.username, username)
						inner.ServeHTTP(w, r)
					})
				}
			}

			h, _ := restapi.Handler(restapi.Config{
				InstallerAPI:      fakeInventory{},
				EventsAPI:         nil,
				Logger:            logrus.Printf,
				VersionsAPI:       nil,
				ManagedDomainsAPI: nil,
				InnerMiddleware:   authMiddleware(),
			})
			bmclient := client.New(client.Config{
				URL: &url.URL{
					Scheme: client.DefaultSchemes[0],
					Host:   "localhost:8081",
					Path:   client.DefaultBasePath,
				},
				AuthInfo: Authenticate(tt.username),
			})
			go http.ListenAndServe("localhost:8081", h)
			bmclient.Installer.ListClusters(context.TODO(), &client_installer.ListClustersParams{})
		})
	}
}

type fakeInventory struct{}

func (f fakeInventory) CancelInstallation(ctx context.Context, params installer.CancelInstallationParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DeregisterCluster(ctx context.Context, params installer.DeregisterClusterParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DeregisterHost(ctx context.Context, params installer.DeregisterHostParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DisableHost(ctx context.Context, params installer.DisableHostParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DownloadClusterFiles(ctx context.Context, params installer.DownloadClusterFilesParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DownloadClusterISO(ctx context.Context, params installer.DownloadClusterISOParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) DownloadClusterKubeconfig(ctx context.Context, params installer.DownloadClusterKubeconfigParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) EnableHost(ctx context.Context, params installer.EnableHostParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) GenerateClusterISO(ctx context.Context, params installer.GenerateClusterISOParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) GetCluster(ctx context.Context, params installer.GetClusterParams) middleware.Responder {
	return installer.NewGetClusterOK()
}

func (f fakeInventory) GetCredentials(ctx context.Context, params installer.GetCredentialsParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) GetFreeAddresses(ctx context.Context, params installer.GetFreeAddressesParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) GetHost(ctx context.Context, params installer.GetHostParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) GetNextSteps(ctx context.Context, params installer.GetNextStepsParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) InstallCluster(ctx context.Context, params installer.InstallClusterParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) ListClusters(ctx context.Context, params installer.ListClustersParams) middleware.Responder {
	return installer.NewListClustersOK()
}

func (f fakeInventory) ListHosts(ctx context.Context, params installer.ListHostsParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) PostStepReply(ctx context.Context, params installer.PostStepReplyParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) RegisterCluster(ctx context.Context, params installer.RegisterClusterParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) RegisterHost(ctx context.Context, params installer.RegisterHostParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) ResetCluster(ctx context.Context, params installer.ResetClusterParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) SetDebugStep(ctx context.Context, params installer.SetDebugStepParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) UpdateCluster(ctx context.Context, params installer.UpdateClusterParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) UpdateHostInstallProgress(ctx context.Context, params installer.UpdateHostInstallProgressParams) middleware.Responder {
	panic("implement me")
}

func (f fakeInventory) UploadClusterIngressCert(ctx context.Context, params installer.UploadClusterIngressCertParams) middleware.Responder {
	panic("implement me")
}

var _ restapi.InstallerAPI = fakeInventory{}
