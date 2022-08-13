// +build acceptance

package gapi

import (
	"net/url"
	"testing"
)

func TestAcceptance(t *testing.T) {
	client, err := New("http://grafana:3000", Config{
		BasicAuth: url.UserPassword("admin", "admin"),
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	_, err = client.AddTeam("foo", "foo@bar.com")
	if err != nil {
		t.Fatalf("failed to add team: %v", err)
	}
}
