package gapi

import (
	"testing"
)

const (
	createAndUpdatePlaylistResponse = `  {
		"uid": "1",
		"name": "my playlist",
		"interval": "5m"
	}`

	getPlaylistResponse = `{
		"uid": "2",
		"name": "my playlist",
		"interval": "5m",
		"orgId": "my org",
		"items": [
			{
				"id": 1,
				"playlistId": 1,
				"type": "dashboard_by_id",
				"value": "3",
				"order": 1,
				"title":"my dasboard"
			},
			{
				"id": 1,
				"playlistId": 1,
				"type": "dashboard_by_id",
				"value": "3",
				"order": 1,
				"title":"my dasboard"
			}
		]
	}`
)

func TestPlaylistCreateAndUpdate(t *testing.T) {
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{code: 200, body: createAndUpdatePlaylistResponse},
		{code: 200, body: createAndUpdatePlaylistResponse},
	})

	playlist := Playlist{
		Name:     "my playlist",
		Interval: "5m",
		Items: []PlaylistItem{
			{},
		},
	}

	// create
	id, err := client.NewPlaylist(playlist)
	if err != nil {
		t.Fatal(err)
	}

	if id != "1" {
		t.Errorf("Invalid id - %s, Expected %s", id, "1")
	}

	// update
	playlist.Items = append(playlist.Items, PlaylistItem{
		Type:  "dashboard_by_id",
		Value: "1",
		Order: 1,
		Title: "my dashboard",
	})

	err = client.UpdatePlaylist(playlist)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPlaylist(t *testing.T) {
	client := gapiTestTools(t, 200, getPlaylistResponse)

	playlist, err := client.Playlist("2")
	if err != nil {
		t.Fatal(err)
	}

	if playlist.UID != "2" {
		t.Errorf("Invalid id - %s, Expected %s", playlist.UID, "2")
	}

	if len(playlist.Items) != 2 {
		t.Errorf("Invalid len(items) - %d, Expected %d", len(playlist.Items), 2)
	}
}

func TestDeletePlaylist(t *testing.T) {
	client := gapiTestTools(t, 200, "")

	err := client.DeletePlaylist("1")
	if err != nil {
		t.Fatal(err)
	}
}
