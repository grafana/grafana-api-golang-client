package gapi

import (
	"net/http"
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/service_accounts"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
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
	searchServiceAccountsJSON = `[
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
	]`
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
	mocksv, client := gapiTestTools(t, http.StatusOK, createServiceAccountTokenJSON)
	defer mocksv.Close()

	req := models.AddServiceAccountTokenCommand{
		Name:          "key-name",
		SecondsToLive: 0,
	}

	res, err := client.ServiceAccounts.CreateToken(
		service_accounts.NewCreateTokenParams().
			WithBody(&req),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestCreateServiceAccount(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, serviceAccountJSON)
	defer mocksrv.Close()

	req := models.CreateServiceAccountForm{
		Name: "newSA",
	}

	res, err := client.ServiceAccounts.CreateServiceAccount(
		service_accounts.NewCreateServiceAccountParams().
			WithBody(&req),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestUpdateServiceAccount(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, serviceAccountJSON)
	defer mocksrv.Close()

	req := models.UpdateServiceAccountForm{
		Name:       "",
		Role:       "Admin",
		IsDisabled: false,
	}

	res, err := client.ServiceAccounts.UpdateServiceAccount(
		service_accounts.NewUpdateServiceAccountParams().
			WithServiceAccountID(5).
			WithBody(&req),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteServiceAccount(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, deleteServiceAccountJSON)
	defer mocksrv.Close()

	res, err := client.ServiceAccounts.DeleteServiceAccount(
		service_accounts.NewDeleteServiceAccountParams().
			WithServiceAccountID(1),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteServiceAccountToken(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, deleteServiceAccountTokenJSON)
	defer mocksrv.Close()

	res, err := client.ServiceAccounts.DeleteToken(
		service_accounts.NewDeleteTokenParams().
			WithServiceAccountID(1).
			WithTokenID(1),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetServiceAccounts(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, searchServiceAccountsJSON)
	defer mocksrv.Close()

	res, err := client.ServiceAccounts.SearchOrgServiceAccountsWithPaging(
		service_accounts.NewSearchOrgServiceAccountsWithPagingParams(),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetServiceAccountTokens(t *testing.T) {
	mocksrv, client := gapiTestTools(t, http.StatusOK, getServiceAccountTokensJSON)
	defer mocksrv.Close()

	res, err := client.ServiceAccounts.ListTokens(
		service_accounts.NewListTokensParams().
			WithServiceAccountID(5),
		nil,
	)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
