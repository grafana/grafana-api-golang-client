package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateUser creates a Grafana user.
func (c *Client) CreateUser(user User) (int64, error) {
	id := int64(0)
	data, err := json.Marshal(user)
	if err != nil {
		return id, err
	}

	resp, err := c.request("POST", "/api/admin/users", nil, bytes.NewBuffer(data))
	if err != nil {
		return id, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return id, err
	}
	created := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &created)
	if err != nil {
		return id, err
	}
	return created.Id, err
}

// DeleteUser deletes a Grafana user.
func (c *Client) DeleteUser(id int64) error {
	_, err := c.request("DELETE", fmt.Sprintf("/api/admin/users/%d", id), nil, nil)

	return err
}
