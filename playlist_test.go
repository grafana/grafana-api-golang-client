package gapi

import (
	"encoding/json"
	"net/url"
	"reflect"
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

	listOfPlaylists = `[
		{
			"id": 1,
			"uid": "BmMFcuVVz",
			"name": "screen",
			"interval": "1m"
		},
		{
			"id": 2,
			"uid": "uEH1YqVVz",
			"name": "screen2",
			"interval": "1m"
		}
	]`
)

func TestPlaylistCreateAndUpdate(t *testing.T) {
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{200, createAndUpdatePlaylistResponse},
		{200, createAndUpdatePlaylistResponse},
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

func TestGetAllPlaylists(t *testing.T) {
	server, client := gapiTestTools(t, 200, listOfPlaylists)
	defer server.Close()
	existingPlaylists := []Playlist{}
	err := json.Unmarshal([]byte(listOfPlaylists), &existingPlaylists)
	if err != nil {
		t.Errorf("Unable to unmarshal listOfPlaylists - %s", listOfPlaylists)
	}

	playlists, err := client.Playlists(url.Values{})
	if err != nil {
		t.Fatal(err)
	}

	existingID := false
	for _, playlist := range *playlists {
		if playlist.ID == 2 {
			existingID = true
			if !reflect.DeepEqual(playlist, existingPlaylists[1]) {
				t.Errorf("The existing playlists definition don't match the listed item")
			}
		}
	}
	if !existingID {
		t.Errorf("Playlists didn't include any ID == 2 and thus didn't get the correct data")
	}
	if len(*playlists) != 2 {
		t.Errorf("The number of playlists should be 2 but was %v", len(*playlists))
	}
}

func TestDeletePlaylist(t *testing.T) {
	client := gapiTestTools(t, 200, "")

	err := client.DeletePlaylist("1")
	if err != nil {
		t.Fatal(err)
	}
}
