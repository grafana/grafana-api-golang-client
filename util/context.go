package util

import (
	"context"

	"github.com/grafana/grafana-api-golang-client/goclient"
)

func GetContextWithBasicAuth() context.Context {
	return context.WithValue(context.Background(), goclient.ContextBasicAuth, goclient.BasicAuth{
		UserName: "admin",
		Password: "admin",
	})
}
