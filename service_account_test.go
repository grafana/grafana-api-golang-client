package gapi

import (
	"net/http"
	"testing"

	"github.com/gobs/pretty"
)

const (
	createServiceAccountTokenJSON = `{"name":"key-name", "key":"mock-api-key"}`
	deleteServiceAccountTokenJSON = `{"message":"Service account token deleted"}`

	getServiceAccountTokensJSON = `[
		{
			"id": 4,
			"name": "testToken",
			"created": "2022-06-15T15:19:00+02:00",
			"expiration": "2022-06-15T16:17:20+02:00",
			"secondsUntilExpiration": 3412.443626017,
			"hasExpired": false
		},
		{
			"id": 1,
			"name": "testToken2",
			"created": "2022-01-15T15:19:00+02:00",
			"expiration": "2022-02-15T16:17:20+02:00",
			"secondsUntilExpiration": 0,
			"hasExpired": true
		},
		{
			"id": 6,
			"name": "testTokenzx",
			"created": "2022-06-15T15:39:54+02:00",
			"expiration": null,
			"secondsUntilExpiration": 0,
			"hasExpired": false
		}
	]`
)

func TestCreateServiceAccountToken(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, createServiceAccountTokenJSON)
	defer server.Close()

	req := CreateServiceAccountTokenRequest{
		Name:          "key-name",
		SecondsToLive: 0,
	}

	res, err := client.CreateServiceAccountToken(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteServiceAccountToken(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, deleteServiceAccountTokenJSON)
	defer server.Close()

	res, err := client.DeleteServiceAccountToken(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetServiceAccountTokens(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, getServiceAccountTokensJSON)
	defer server.Close()

	res, err := client.GetServiceAccountTokens(5)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
