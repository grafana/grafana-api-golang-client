package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	createAPIKeyJSON = `{"name":"key-name", "key":"mock-api-key"}`
	deleteAPIKeyJSON = `{"message":"API key deleted"}`

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
	]`
)

func TestCreateApiKey(t *testing.T) {
	server, client := gapiTestTools(t, 200, createAPIKeyJSON)
	defer server.Close()

	req := CreateAPIKeyRequest{
		Name:          "key-name",
		Role:          "Viewer",
		SecondsToLive: 0,
	}

	res, err := client.CreateAPIKey(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteApiKey(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteAPIKeyJSON)
	defer server.Close()

	res, err := client.DeleteAPIKey(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetApiKeys(t *testing.T) {
	server, client := gapiTestTools(t, 200, getAPIKeysJSON)
	defer server.Close()

	res, err := client.GetAPIKeys(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
