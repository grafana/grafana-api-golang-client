package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Folder struct {
	Id    int64  `json:"id"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
}

// Folders fetches and returns Grafana folders.
func (c *Client) Folders() ([]Folder, error) {
	folders := make([]Folder, 0)
	resp, err := c.request("GET", "/api/folders/", nil, nil)
	if err != nil {
		return folders, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return folders, err
	}
	err = json.Unmarshal(data, &folders)
	return folders, err
}

// Folder fetches and returns the Grafana folder whose ID it's passed.
func (c *Client) Folder(id int64) (*Folder, error) {
	folder := &Folder{}
	resp, err := c.request("GET", fmt.Sprintf("/api/folders/id/%d", id), nil, nil)
	if err != nil {
		return folder, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return folder, err
	}
	err = json.Unmarshal(data, &folder)
	return folder, err
}

// NewFolder creates a new Grafana folder.
func (c *Client) NewFolder(title string) (Folder, error) {
	folder := Folder{}
	dataMap := map[string]string{
		"title": title,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return folder, err
	}
	resp, err := c.request("POST", "/api/folders", nil, bytes.NewBuffer(data))
	if err != nil {
		return folder, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return folder, err
	}
	err = json.Unmarshal(data, &folder)
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
	_, err = c.request("PUT", fmt.Sprintf("/api/folders/%s", id), nil, bytes.NewBuffer(data))

	return err
}

// DeleteFolder deletes the folder whose ID it's passed.
func (c *Client) DeleteFolder(id string) error {
	_, err := c.request("DELETE", fmt.Sprintf("/api/folders/%s", id), nil, nil)

	return err
}
