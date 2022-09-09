package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RoleAssignments struct {
	RoleUID         string           `json:"role_uid"`
	Users           []RoleAssignment `json:"users,omitempty"`
	Teams           []RoleAssignment `json:"teams,omitempty"`
	ServiceAccounts []RoleAssignment `json:"service_accounts,omitempty"`
}

type RoleAssignment struct {
	ID     int  `json:"id"`
	Global bool `json:"global"`
}

func (c *Client) GetRoleAssignments(uid string) (*RoleAssignments, error) {
	assignments := &RoleAssignments{}
	url := fmt.Sprintf("/api/access-control/roles/%s/assignments", uid)
	if err := c.request(http.MethodGet, url, nil, nil, assignments); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (c *Client) UpdateRoleAssignments(ra RoleAssignments) (*RoleAssignments, error) {
	response := &RoleAssignments{}

	data, err := json.Marshal(ra)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("/api/access-control/roles/%s/assignments", ra.RoleUID)
	err = c.request(http.MethodPut, url, nil, bytes.NewBuffer(data), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
