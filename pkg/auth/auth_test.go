package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

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
			name:     "username exist",
			username: "eran",
		},
		//{
		//	name:     "username empty",
		//	username: "",
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create a handler to use as "next" which will verify the request
			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				val := r.Context().Value(ContextUsernameKey)
				assert.NotNil(t, val)
				valStr, ok := val.(string)
				if !ok {
					t.Error("not string")
				}
				assert.Equal(t, tt.username, valStr)
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
			ree, e := bmclient.Installer.ListClusters(context.TODO(), &client_installer.ListClustersParams{})
			fmt.Println(ree)
			fmt.Println(e)
		})
	}
}
