package auth

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type contextKey string

const ContextUsernameKey = contextKey("username")

// Fake auth Middleware handler to add username from headers to request context
func GetUserMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Update the context, as the jwt middleware will update it
		ctx := r.Context()
		username := r.Header.Get("username")
		log.Infof("ERAN auth ctx: %v", ctx)
		log.Infof("ERAN auth username: %v", username)
		// Append the username to the request context
		ctx = context.WithValue(ctx, ContextUsernameKey, username)
		*r = *r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// Authenticate will add username to the request headers
func Authenticate(user string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("username", user)
	})
}
