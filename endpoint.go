package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	m "github.com/grafana/grafana/pkg/models"
)

type Endpoint struct {
	Id    int64
	OrgId int64 `json:"org_id"`
	Name  string
	Slug  string
	Tags  []string
}

// note: grafana will override the org-id in your query with the active org based on your login.
func (c *Client) Endpoints(settings m.GetEndpointsQuery) ([]Endpoint, error) {
	data, _ := json.Marshal(settings)
	endpoints := make([]Endpoint, 0)

	req, err := c.newRequest("GET", "/api/endpoints/", bytes.NewBuffer(data))
	if err != nil {
		return endpoints, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return endpoints, err
	}
	if resp.StatusCode != 200 {
		return endpoints, errors.New(resp.Status)
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return endpoints, err
	}
	err = json.Unmarshal(data, &endpoints)
	return endpoints, err
}

func (c *Client) NewEndpoint(settings m.AddEndpointCommand) error {
	data, err := json.Marshal(settings)
	req, err := c.newRequest("PUT", "/api/endpoints", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println("response BODY", string(data))
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return err
}

func (c *Client) DeleteEndpoint(id int64) error {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/endpoints/%d", id), nil)
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
