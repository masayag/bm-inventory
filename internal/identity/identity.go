package identity

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/filanov/bm-inventory/pkg/auth"

	"github.com/thoas/go-funk"
)

var ADMIN_LIST = []string{"ercohen", "mfilanov", "rfreiman", "alazar"}

func GetOwner(ctx context.Context, log logrus.FieldLogger) string {
	username := ctx.Value(auth.ContextUsernameKey)
	log.Infof("ERAN GetOwner ctx: %v", ctx)
	log.Infof("ERAN GetOwner username: %v", username)
	if username == nil {
		username = ""
	}
	return username.(string)
}

func GetOwnerFilter(ctx context.Context, log logrus.FieldLogger) string {
	query := ""
	username := GetOwner(ctx, log)
	if username != "" && !funk.ContainsString(ADMIN_LIST, username) {
		query = fmt.Sprintf("owner = '%s'", username)
	}
	return query
}
