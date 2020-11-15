package gapi

import (
	"github.com/gobs/pretty"
	"testing"
)

const (
	createApiKeyJSON = `{"name":"key-name", "key":"mock-api-key"}`
	deleteApiKeyJSON = `{"message":"API key deleted"}`

	getApiKeysJSON = `[
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
	server, client := gapiTestTools(t, 200, createApiKeyJSON)
	defer server.Close()

	req := CreateApiKeyRequest{
		Name:          "key-name",
		Role:          "Viewer",
		SecondsToLive: 0,
	}

	res, err := client.CreateApiKey(req)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestDeleteApiKey(t *testing.T) {
	server, client := gapiTestTools(t, 200, deleteApiKeyJSON)
	defer server.Close()

	res, err := client.DeleteApiKey(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetApiKeys(t *testing.T) {
	server, client := gapiTestTools(t, 200, getApiKeysJSON)
	defer server.Close()

	res, err := client.GetApiKeys(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
