package gapi

import (
	"context"
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient"
)

const (
	createAPIKeyJSON = `{"name":"key-name", "key":"mock-api-key"}` //#nosec
	deleteAPIKeyJSON = `{"message":"API key deleted"}`             //#nosec

	getAPIKeysJSON = `[
		{
			"id": 1,
			"name": "key-name-2",
			"role": "Viewer"
		},
		{
			"id": 2,
			"name": "key-name-2",
			"role": "Admin",
			"expiration": "2021-10-30T10:52:03+03:00"
		}
	]`  //#nosec
)

func getContextWithBasicAuth() context.Context {
	return context.WithValue(context.Background(), goclient.ContextBasicAuth, goclient.BasicAuth{
		UserName: "admin",
		Password: "admin",
	})
}

func TestCreateAPIKey(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createAPIKeyJSON)
	defer mocksrv.Close()

	req := goclient.AddApiKeyCommandModel{
		Name:          "key-name",
		Role:          "Viewer",
		SecondsToLive: 0,
	}

	client := getClient(mocksrv.server.URL)

	_, res, err := client.ApiKeysApi.AddAPIkey(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteAPIKey(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteAPIKeyJSON)
	defer server.Close()

	res, err := client.DeleteAPIKey(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetAPIKeys(t *testing.T) {
	server, client := gapiTestTools(t, 200, getAPIKeysJSON)
	defer server.Close()

	res, err := client.GetAPIKeys(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
