package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	createdSnapshotResponse = `{
		"deleteKey":"XXXXXXX",
		"deleteUrl":"myurl/api/snapshots-delete/XXXXXXX",
		"key":"YYYYYYY",
		"url":"myurl/dashboard/snapshot/YYYYYYY",
		"id": 1
	}`
)

func TestSnapshotCreate(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdSnapshotResponse)
	defer server.Close()

	snapshot := Snapshot{
		Model: map[string]interface{}{
			"title": "test",
		},
		Expires: 3600,
	}

	resp, err := client.NewSnapshot(snapshot)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.DeleteKey != "XXXXXXX" {
		t.Errorf("Invalid key - %s, Expected %s", resp.DeleteKey, "XXXXXXX")
	}

	for _, code := range []int{400, 401, 403, 412} {
		server.code = code
		_, err = client.NewSnapshot(snapshot)
		if err == nil {
			t.Errorf("%d not detected", code)
		}
	}
}
