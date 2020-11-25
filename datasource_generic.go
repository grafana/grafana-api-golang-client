package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type DataSourceGeneric struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Access string `json:"access"`

	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	// Deprecated in favor of secureJsonData.password
	Password string `json:"password,omitempty"`

	OrgId     int64 `json:"orgId,omitempty"`
	IsDefault bool  `json:"isDefault"`

	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser,omitempty"`
	// Deprecated in favor of secureJsonData.basicAuthPassword
	BasicAuthPassword string `json:"basicAuthPassword,omitempty"`

	JSONData       JsonData `json:"jsonData,omitempty"`
	SecureJSONData JsonData `json:"secureJsonData,omitempty"`
}

type JsonData map[string]interface{}

func (c *Client) NewDataSourceGeneric(s *DataSourceGeneric) (int64, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}
	req, err := c.newRequest("POST", "/api/datasources", nil, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("creating data source failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	result := struct {
		Id int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.Id, err
}

func (c *Client) UpdateDataSourceGeneric(s *DataSourceGeneric) error {
	path := fmt.Sprintf("/api/datasources/%d", s.Id)
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	req, err := c.newRequest("PUT", path, nil, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("updating data source failed: %s", resp.Status)
	}

	return nil
}

func (c *Client) DataSourceGeneric(id int64) (*DataSourceGeneric, error) {
	path := fmt.Sprintf("/api/datasources/%d", id)
	req, err := c.newRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("getting data source failed: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &DataSourceGeneric{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) DeleteDataSourceGeneric(id int64) error {
	path := fmt.Sprintf("/api/datasources/%d", id)
	req, err := c.newRequest("DELETE", path, nil, nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed deleting data source: %s", resp.Status)
	}

	return nil
}
