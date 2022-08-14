// +build integration

package gapi

import (
	"log"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	client *Client
)

func TestMain(m *testing.M) {
	auth := os.Getenv("GRAFANA_AUTH")
	authParts := strings.Split(auth, ":")
	var err error
	client, err = New(os.Getenv("GRAFANA_URL"), Config{
		BasicAuth: url.UserPassword(authParts[0], authParts[1]),
	})
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}
