package gapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func gapiTestTools(code int, body string) (*httptest.Server, *Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, body)
	}))

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: tr}

	url := url.URL{
		Scheme: "http",
		Host:   "my-grafana.com",
	}

	client := &Client{"my-key", url, httpClient}

	return server, client
}
