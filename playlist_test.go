package gapi

import (
	"testing"
)

const (
	createAndUpdatePlaylistResponse = `  {
		"id": 1,
		"name": "my playlist",
		"interval": "5m"
	}`

	getPlaylistResponse = `{
		"id" : 2,
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
	server, client := gapiTestTools(200, createAndUpdatePlaylistResponse)
	defer server.Close()

	playlist := Playlist{
		Name:     "my playlist",
		Interval: "5m",
		Items: []PlaylistItem{
			PlaylistItem{},
		},
	}

	// create
	id, err := client.NewPlaylist(playlist)
	if err != nil {
		t.Fatal(err)
	}

	if id != 1 {
		t.Errorf("Invalid id - %d, Expected %d", id, 1)
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
	server, client := gapiTestTools(200, getPlaylistResponse)
	defer server.Close()

	playlist, err := client.Playlist(1)
	if err != nil {
		t.Error(err)
	}

	if playlist.Id != 2 {
		t.Errorf("Invalid id - %d, Expected %d", playlist.Id, 1)
	}

	if len(playlist.Items) != 2 {
		t.Errorf("Invalid len(items) - %d, Expected %d", len(playlist.Items), 2)
	}
}

func TestDeletePlaylist(t *testing.T) {
	server, client := gapiTestTools(200, "")
	defer server.Close()

	err := client.DeletePlaylist(1)
	if err != nil {
		t.Error(err)
	}
}
