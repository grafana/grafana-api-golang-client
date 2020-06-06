package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Org struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// Orgs fetches and returns the Grafana orgs.
func (c *Client) Orgs() ([]Org, error) {
	orgs := make([]Org, 0)

	resp, err := c.request("GET", "/api/orgs/", nil, nil)
	if err != nil {
		return orgs, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orgs, err
	}
	err = json.Unmarshal(data, &orgs)
	return orgs, err
}

// OrgByName fetches and returns the org whose name it's passed.
func (c *Client) OrgByName(name string) (Org, error) {
	org := Org{}
	resp, err := c.request("GET", fmt.Sprintf("/api/orgs/name/%s", name), nil, nil)
	if err != nil {
		return org, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	return org, err
}

// Org fetches and returns the org whose ID it's passed.
func (c *Client) Org(id int64) (Org, error) {
	org := Org{}
	resp, err := c.request("GET", fmt.Sprintf("/api/orgs/%d", id), nil, nil)
	if err != nil {
		return org, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return org, err
	}
	err = json.Unmarshal(data, &org)
	return org, err
}

// NewOrg creates a new Grafana org.
func (c *Client) NewOrg(name string) (int64, error) {
	id := int64(0)

	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return id, err
	}
	resp, err := c.request("POST", "/api/orgs", nil, bytes.NewBuffer(data))
	if err != nil {
		return id, err
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return id, err
	}
	tmp := struct {
		Id int64 `json:"orgId"`
	}{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return id, err
	}
	id = tmp.Id
	return id, err
}

// UpdateOrg updates a Grafana org.
func (c *Client) UpdateOrg(id int64, name string) error {
	dataMap := map[string]string{
		"name": name,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}
	_, err = c.request("PUT", fmt.Sprintf("/api/orgs/%d", id), nil, bytes.NewBuffer(data))

	return err
}

// DeleteOrg deletes the Grafana org whose ID it's passed.
func (c *Client) DeleteOrg(id int64) error {
	_, err := c.request("DELETE", fmt.Sprintf("/api/orgs/%d", id), nil, nil)

	return err
}
