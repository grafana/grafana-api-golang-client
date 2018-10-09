package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Folder struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func (c *Client) Folders() ([]Folder, error) {
	folders := make([]Folder, 0)

	req, err := c.newRequest("GET", "/api/folders/", nil, nil)
	if err != nil {
		return folders, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return folders, err
	}
	if resp.StatusCode != 200 {
		return folders, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return folders, err
	}
	err = json.Unmarshal(data, &folders)
	return folders, err
}

func (c *Client) Folder(id int64) (Folder, error) {
	folder := Folder{}
	req, err := c.newRequest("GET", fmt.Sprintf("/api/folders/%d", id), nil, nil)
	if err != nil {
		return folder, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return folder, err
	}
	if resp.StatusCode != 200 {
		return folder, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return folder, err
	}
	err = json.Unmarshal(data, &folder)
	return folder, err
}

func (c *Client) NewFolder(name string) (int64, error) {
	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	id := int64(0)
	req, err := c.newRequest("POST", "/api/folders", nil, bytes.NewBuffer(data))
	if err != nil {
		return id, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return id, err
	}
	if resp.StatusCode != 200 {
		return id, errors.New(resp.Status)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return id, err
	}
	tmp := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return id, err
	}
	id = tmp.Id
	return id, err
}

func (c *Client) UpdateFolder(id int64, name string) error {
	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	req, err := c.newRequest("PUT", fmt.Sprintf("/api/folders/%d", id), nil, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}

func (c *Client) DeleteFolder(id int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/folders/%d", id), nil, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}
