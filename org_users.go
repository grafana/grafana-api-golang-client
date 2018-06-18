package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type OrgUser struct {
	OrgId   int64
	UserId  int64
	Email   string
	Login   string
	Role    string
}

func (c *Client) OrgUsers(id int64) ([]OrgUser, error) {
	users := make([]OrgUser, 0)
	req, err := c.newRequest("GET", fmt.Sprintf("/api/orgs/%d/users", id), nil)
	if err != nil {
		return users, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return users, err
	}
	if resp.StatusCode != 200 {
		return users, errors.New(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return users, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return users, err
	}
	return users, err
}

func (c *Client) AddOrgUser(orgId int64, user, role string) error {
	dataMap := map[string]string{
		"loginOrEmail": user,
		"role": role,
	}
	data, err := json.Marshal(dataMap)
	req, err := c.newRequest("POST", fmt.Sprintf("/api/orgs/%d/users", orgId), bytes.NewBuffer(data))
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

func (c *Client) UpdateOrgUser(orgId, userId int64, role string) error {
	dataMap := map[string]string{
		"role": role,
	}
	data, err := json.Marshal(dataMap)
	req, err := c.newRequest("PATCH", fmt.Sprintf("/api/orgs/%d/users/%d", orgId, userId), bytes.NewBuffer(data))
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

func (c *Client) RemoveOrgUser(orgId, userId int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/orgs/%d/users/%d", orgId, userId), nil)
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
