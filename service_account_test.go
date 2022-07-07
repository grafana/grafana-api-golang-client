package gapi

import (
	"net/http"
	"testing"

	"github.com/gobs/pretty"
)

const (
	serviceAccountJSON = `{
	"id": 8,
	"name": "newSA",
	"login": "sa-newsa",
	"orgId": 1,
	"isDisabled": false,
	"role": "",
	"tokens": 0,
	"avatarUrl": ""
}`
	searchServiceAccountsJSON = `{
	"totalCount": 2,
	"serviceAccounts": [
		{
			"id": 8,
			"name": "newSA",
			"login": "sa-newsa",
			"orgId": 1,
			"isDisabled": false,
			"role": "",
			"tokens": 1,
			"avatarUrl": "/avatar/0e94f33c929884a5163d953582f27fec"
		},
		{
			"id": 9,
			"name": "newnewSA",
			"login": "sa-newnewsa",
			"orgId": 1,
			"isDisabled": true,
			"role": "Admin",
			"tokens": 2,
			"avatarUrl": "/avatar/0e29f33c929824a5163d953582e83abe"
		}
	],
	"page": 1,
	"perPage": 1000
}`
	createServiceAccountTokenJSON = `{"name":"key-name", "key":"mock-api-key"}`   //#nosec
	deleteServiceAccountTokenJSON = `{"message":"Service account token deleted"}` //#nosec
	deleteServiceAccountJSON      = `{"message":"service account deleted"}`

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
	]`  //#nosec
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

func TestCreateServiceAccount(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, serviceAccountJSON)
	defer server.Close()

	isDisabled := true
	req := CreateServiceAccountRequest{
		Name:       "newSA",
		Role:       "Admin",
		IsDisabled: &isDisabled,
	}

	res, err := client.CreateServiceAccount(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestUpdateServiceAccount(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, serviceAccountJSON)
	defer server.Close()

	isDisabled := false
	req := UpdateServiceAccountRequest{
		Name:       "",
		Role:       "Admin",
		IsDisabled: &isDisabled,
	}

	res, err := client.UpdateServiceAccount(5, req)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteServiceAccount(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, deleteServiceAccountJSON)
	defer server.Close()

	res, err := client.DeleteServiceAccount(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteServiceAccountToken(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, deleteServiceAccountTokenJSON)
	defer server.Close()

	res, err := client.DeleteServiceAccountToken(int64(1), int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetServiceAccounts(t *testing.T) {
	server, client := gapiTestTools(t, http.StatusOK, searchServiceAccountsJSON)
	defer server.Close()

	res, err := client.GetServiceAccounts()
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
