package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/grafana/grafana-api-golang-client/goclient/client"
	"github.com/grafana/grafana-api-golang-client/goclient/client/api_keys"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
)

// This function creates a API key inside the Grafana instance running in stack `stack`. It's used in order
// to provision API keys inside Grafana while just having access to a Grafana Cloud API key.
//
// See https://grafana.com/docs/grafana-cloud/api/#create-grafana-api-keys for more information.
func (c *CloudClient) CreateGrafanaAPIKeyFromCloud(stack string, input *models.AddAPIKeyCommand) (*models.NewAPIKeyResult, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp := &models.NewAPIKeyResult{}
	err = c.request("POST", fmt.Sprintf("/api/instances/%s/api/auth/keys", stack), nil, bytes.NewBuffer(data), resp)
	return resp, err
}

// The Grafana Cloud API is disconnected from the Grafana API on the stacks unfortunately. That's why we can't use
// the Grafana Cloud API key to fully manage API keys on the Grafana API. The only thing we can do is to create
// a temporary Admin key, and create a Grafana API client with that.
func (c *CloudClient) CreateTemporaryStackGrafanaClient(stackSlug, tempKeyPrefix string, tempKeyDuration time.Duration) (tempClient *client.GrafanaHTTPAPI, cleanup func() error, err error) {
	stack, err := c.StackBySlug(stackSlug)
	if err != nil {
		return nil, nil, err
	}

	name := fmt.Sprintf("%s-%d", tempKeyPrefix, time.Now().UnixNano())
	req := &models.AddAPIKeyCommand{
		Name:          name,
		Role:          "Admin",
		SecondsToLive: int64(tempKeyDuration.Seconds()),
	}

	apiKey, err := c.CreateGrafanaAPIKeyFromCloud(stackSlug, req)
	if err != nil {
		return nil, nil, err
	}

	client, err := GetClient(stack.URL)
	if err != nil {
		return nil, nil, err
	}

	cleanup = func() error {
		_, err = client.APIKeys.DeleteAPIkey(api_keys.NewDeleteAPIkeyParams().WithID(apiKey.ID), nil)
		return err
	}

	return client, cleanup, nil
}
