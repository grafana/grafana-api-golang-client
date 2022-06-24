package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/api_keys"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
	"github.com/stretchr/testify/require"
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

func TestCreateAPIKey(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createAPIKeyJSON)
	defer mocksrv.Close()

	params := api_keys.AddAPIkeyParams{
		Body: &models.AddAPIKeyCommand{
			Name:          "key-name",
			Role:          "Viewer",
			SecondsToLive: 0,
		},
	}
	
	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	resp, err := client.APIKeys.AddAPIkey(&params, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(pretty.PrettyFormat(resp))
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
