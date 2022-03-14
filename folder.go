package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Folder represents a Grafana folder.
type Folder struct {
	ID    int64  `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Folders fetches and returns Grafana folders.
func (c *Client) Folders() ([]Folder, error) {
	folders := make([]Folder, 0)
	err := c.request("GET", "/api/folders/", nil, nil, &folders)
	if err != nil {
		return folders, err
	}

	return folders, err
}

// Folder fetches and returns the Grafana folder whose ID it's passed.
func (c *Client) Folder(id int64) (*Folder, error) {
	folder := &Folder{}
	err := c.request("GET", fmt.Sprintf("/api/folders/id/%d", id), nil, nil, folder)
	if err != nil {
		return folder, err
	}

	return folder, err
}

// Folder fetches and returns the Grafana folder whose UID it's passed.
func (c *Client) FolderByUID(uid string) (*Folder, error) {
	folder := &Folder{}
	err := c.request("GET", fmt.Sprintf("/api/folders/%s", uid), nil, nil, folder)
	if err != nil {
		return folder, err
	}

	return folder, err
}

// NewFolder creates a new Grafana folder.
func (c *Client) NewFolder(title string, uid ...string) (Folder, error) {
	if len(uid) > 1 {
		return Folder{}, fmt.Errorf("too many arguments. Expected 1 or 2")
	}

	folder := Folder{}
	dataMap := map[string]string{
		"title": title,
	}
	if len(uid) == 1 {
		dataMap["uid"] = uid[0]
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return folder, err
	}

	err = c.request("POST", "/api/folders", nil, bytes.NewBuffer(data), &folder)
	if err != nil {
		return folder, err
	}

	return folder, err
}

// UpdateFolder updates the folder whose ID it's passed.
func (c *Client) UpdateFolder(id string, name string) error {
	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request("PUT", fmt.Sprintf("/api/folders/%s", id), nil, bytes.NewBuffer(data), nil)
}

// DeleteFolder deletes the folder whose ID it's passed.
func (c *Client) DeleteFolder(id string) error {
	return c.request("DELETE", fmt.Sprintf("/api/folders/%s", id), nil, nil, nil)
}
