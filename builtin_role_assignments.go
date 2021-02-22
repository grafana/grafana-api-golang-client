package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const baseUrl = "/api/access-control/builtin-roles"

type BuiltInRoleAssignment struct {
	BuiltinRole string `json:"builtInRole"`
	RoleUID     string `json:"roleUid"`
	Global      bool   `json:"global"`
}

// GetBuiltInRoleAssignments gets all built-in role assignments. Available only in Grafana Enterprise 8.+.
func (c *Client) GetBuiltInRoleAssignments() (map[string][]*Role, error) {
	br := make(map[string][]*Role, 0)
	err := c.request("GET", baseUrl, nil, nil, &br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

// NewBuiltInRoleAssignment creates a new built-in role assignment. Available only in Grafana Enterprise 8.+.
func (c *Client) NewBuiltInRoleAssignment(builtInRoleAssignment BuiltInRoleAssignment) (*BuiltInRoleAssignment, error) {
	body, err := json.Marshal(builtInRoleAssignment)
	if err != nil {
		return nil, err
	}

	br := &BuiltInRoleAssignment{}

	err = c.request("POST", baseUrl, nil, bytes.NewBuffer(body), &br)
	if err != nil {
		return nil, err
	}

	return br, err
}

// DeleteBuiltInRoleAssignment remove the built-in role assignments. Available only in Grafana Enterprise 8.+.
func (c *Client) DeleteBuiltInRoleAssignment(builtInRole BuiltInRoleAssignment) error {
	data, err := json.Marshal(builtInRole)
	if err != nil {
		return err
	}

	qp := map[string][]string{
		"global": {fmt.Sprint(builtInRole.Global)},
	}
	url := fmt.Sprintf("%s/%s/roles/%s", baseUrl, builtInRole.BuiltinRole, builtInRole.RoleUID)
	err = c.request("DELETE", url, qp, bytes.NewBuffer(data), nil)

	return err
}
