package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateCloudAPIKeyInput struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type ListCloudAPIKeysOutput struct {
	Items []*CloudAPIKey
}

type CloudAPIKey struct {
	ID         int
	Name       string
	Role       string
	Token      string
	Expiration string
}

// This function creates a API key inside the Grafana instance running in stack `stack`. It's used in order
// to provision API keys inside Grafana while just having access to a Grafana Cloud API key.
//
// See https://grafana.com/docs/grafana-cloud/api/#create-grafana-api-keys for more information.
func (c *Client) CreateGrafanaAPIKeyFromCloud(stack string, input *CreateAPIKeyRequest) (*CreateAPIKeyResponse, error) {
	resp := CreateAPIKeyResponse{}
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.request("POST", fmt.Sprintf("/api/instances/%s/api/auth/keys", stack), nil, bytes.NewBuffer(data), &resp)
	return &resp, err
}

func (c *Client) CreateCloudAPIKey(org string, input *CreateCloudAPIKeyInput) (*CloudAPIKey, error) {
	resp := CloudAPIKey{}
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	err = c.request("POST", fmt.Sprintf("/api/orgs/%s/api-keys", org), nil, bytes.NewBuffer(data), &resp)
	return &resp, err
}

func (c *Client) ListCloudAPIKeys(org string) (*ListCloudAPIKeysOutput, error) {
	resp := &ListCloudAPIKeysOutput{}
	err := c.request("GET", fmt.Sprintf("/api/orgs/%s/api-keys", org), nil, nil, &resp)
	return resp, err
}

func (c *Client) DeleteCloudAPIKey(org string, keyName string) error {
	return c.request("DELETE", fmt.Sprintf("/api/orgs/%s/api-keys/%s", org, keyName), nil, nil, nil)
}
