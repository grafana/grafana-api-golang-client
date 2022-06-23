package gapi

import (
	"net/url"

	"github.com/grafana/grafana-api-golang-client/goclient/client"
)

func GetClient(serverURL string) (*client.GrafanaHTTPAPI, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := client.NewHTTPClientWithConfig(
		nil,
		client.DefaultTransportConfig().
		WithHost(u.Host).
		WithSchemes([]string{"http"}),
	)
	return c, nil
}
