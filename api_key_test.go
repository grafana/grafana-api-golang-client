package gapi

import (
	"testing"

	"github.com/gobs/pretty"
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
	]` //#nosec
)

func TestCreateAPIKey(t *testing.T) {
	client := gapiTestTools(t, 200, createAPIKeyJSON)

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

func TestDeleteAPIKey(t *testing.T) {
	client := gapiTestTools(t, 200, deleteAPIKeyJSON)

	res, err := client.DeleteAPIKey(int64(1))
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}

func TestGetAPIKeys(t *testing.T) {
	client := gapiTestTools(t, 200, getAPIKeysJSON)

	res, err := client.GetAPIKeys(true)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(res))
}
