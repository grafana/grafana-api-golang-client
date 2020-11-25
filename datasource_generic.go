package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DataSourceGeneric struct {
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Access string `json:"access"`

	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	// Deprecated: Use secureJsonData.password instead.
	Password string `json:"password,omitempty"`

	OrgID     int64 `json:"orgId,omitempty"`
	IsDefault bool  `json:"isDefault"`

	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser,omitempty"`
	// Deprecated: Use secureJsonData.basicAuthPassword instead.
	BasicAuthPassword string `json:"basicAuthPassword,omitempty"`

	JSONData       JSONDataRaw `json:"jsonData,omitempty"`
	SecureJSONData JSONDataRaw `json:"secureJsonData,omitempty"`
}

type JSONDataRaw map[string]interface{}

func (c *Client) NewDataSourceGeneric(s *DataSourceGeneric) (int64, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}

	result := struct {
		ID int64 `json:"id"`
	}{}

	err = c.request("POST", "/api/datasources", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return 0, err
	}

	return result.ID, err
}

func (c *Client) UpdateDataSourceGeneric(s *DataSourceGeneric) error {
	path := fmt.Sprintf("/api/datasources/%d", s.ID)
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return c.request("PUT", path, nil, bytes.NewBuffer(data), nil)
}

func (c *Client) DataSourceGeneric(id int64) (*DataSourceGeneric, error) {
	path := fmt.Sprintf("/api/datasources/%d", id)
	result := &DataSourceGeneric{}
	err := c.request("GET", path, nil, nil, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) DeleteDataSourceGeneric(id int64) error {
	path := fmt.Sprintf("/api/datasources/%d", id)

	return c.request("DELETE", path, nil, nil, nil)
}
