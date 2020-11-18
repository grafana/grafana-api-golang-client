package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type CreateApiKeyRequest struct {
	Name          string `json:"name"`
	Role          string `json:"role"`
	SecondsToLive int64  `json:"secondsToLive,omitempty"`
}

type CreateApiKeyResponse struct {
	// Id field only returned after Grafana v7.
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type GetApiKeysResponse struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Role       string    `json:"role"`
	Expiration time.Time `json:"expiration,omitempty"`
}

type DeleteApiKeyResponse struct {
	Message string `json:"message"`
}

// CreateApiKey creates a new Grafana API key.
func (c *Client) CreateApiKey(request CreateApiKeyRequest) (CreateApiKeyResponse, error) {
	response := CreateApiKeyResponse{}

	data, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	err = c.request("POST", "/api/auth/keys", nil, bytes.NewBuffer(data), &response)
	return response, err
}

// GetApiKeys retrieves a list of all API keys.
func (c *Client) GetApiKeys(includeExpired bool) ([]*GetApiKeysResponse, error) {
	response := make([]*GetApiKeysResponse, 0)

	query := url.Values{}
	query.Add("includeExpired", strconv.FormatBool(includeExpired))

	err := c.request("GET", "/api/auth/keys", query, nil, &response)
	return response, err
}

// DeleteApiKey deletes the Grafana API key with the specified ID.
func (c *Client) DeleteApiKey(id int64) (DeleteApiKeyResponse, error) {
	response := DeleteApiKeyResponse{}

	path := fmt.Sprintf("/api/auth/keys/%d", id)
	err := c.request("DELETE", path, nil, nil, &response)
	return response, err
}
