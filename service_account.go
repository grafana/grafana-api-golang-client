package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CreateServiceAccountTokenRequest struct {
	Name             string `json:"name"`
	ServiceAccountID int64  `json:"-"`
	SecondsToLive    int64  `json:"secondsToLive,omitempty"`
}

type CreateServiceAccountTokenResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type GetServiceAccountTokensResponse struct {
	ID                     int64      `json:"id"`
	Name                   string     `json:"name"`
	Created                time.Time  `json:"created,omitempty"`
	Expiration             *time.Time `json:"expiration,omitempty"`
	SecondsUntilExpiration *float64   `json:"secondsUntilExpiration,omitempty"`
	HasExpired             bool       `json:"hasExpired,omitempty"`
}

type DeleteServiceAccountTokenResponse struct {
	Message string `json:"message"`
}

// CreateServiceAccountToken creates a new Grafana service account token.
func (c *Client) CreateServiceAccountToken(request CreateServiceAccountTokenRequest) (CreateServiceAccountTokenResponse, error) {
	response := CreateServiceAccountTokenResponse{}

	data, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	err = c.request(http.MethodPost,
		fmt.Sprintf("/api/serviceaccounts/%d/tokens", request.ServiceAccountID),
		nil, bytes.NewBuffer(data), &response)
	return response, err
}

// GetServiceAccountTokens retrieves a list of all service account tokens for a specific service account.
func (c *Client) GetServiceAccountTokens(serviceAccountID int) ([]GetServiceAccountTokensResponse, error) {
	response := make([]GetServiceAccountTokensResponse, 0)

	err := c.request(http.MethodGet,
		fmt.Sprintf("/api/serviceaccounts/%d/tokens", serviceAccountID),
		nil, nil, &response)
	return response, err
}

// DeleteServiceAccountToken deletes the Grafana service account token with the specified ID.
func (c *Client) DeleteServiceAccountToken(serviceAccountID, tokenID int64) (DeleteServiceAccountTokenResponse, error) {
	response := DeleteServiceAccountTokenResponse{}

	path := fmt.Sprintf("/api/serviceaccounts/%d/tokens/%d", serviceAccountID, tokenID)
	err := c.request(http.MethodDelete, path, nil, nil, &response)
	return response, err
}
