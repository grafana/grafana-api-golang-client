package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/grafana/grafana/pkg/api/dtos"
)

func (c *Client) CreateUserForm(settings dtos.AdminCreateUserForm) (int64, error) {
	id := int64(0)
	data, err := json.Marshal(settings)
	req, err := c.newRequest("POST", "/api/admin/users", bytes.NewBuffer(data))
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
	user := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		return id, err
	}
	return user.Id, err
}

func (c *Client) CreateUser(email, login, name, password string) (int64, error) {
	return c.CreateUserForm(dtos.AdminCreateUserForm{email, login, name, password})
}

func (c *Client) DeleteUser(id int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/admin/users/%d", id), nil)
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
