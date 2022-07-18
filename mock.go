package gapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/grafana/grafana-api-golang-client/goclient/client"
)

type mockServer struct {
	code   int
	server *httptest.Server
}

func (m *mockServer) Close() {
	m.server.Close()
}

func gapiTestTools(t *testing.T, code int, body string) (*mockServer, *client.GrafanaHTTPAPI) {
	t.Helper()

	mock := &mockServer{
		code: code,
	}

	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(mock.code)
		fmt.Fprint(w, body)
	}))

	client, err := GetClient(mock.server.URL)
	if err != nil {
		t.Fatal(err)
	}
	return mock, client
}

func getCloudClient(t *testing.T, srv string) *CloudClient {
	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(srv)
		},
	}

	httpClient := &http.Client{Transport: tr}

	client, err := New("http://my-grafana.com", Config{APIKey: "my-key", Client: httpClient})
	if err != nil {
		t.Fatal(err)
	}
	return client
}
