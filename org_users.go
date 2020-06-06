package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type OrgUser struct {
	OrgId  int64  `json:"orgId"`
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	Login  string `json:"login"`
	Role   string `json:"role"`
}

// OrgUsers fetches and returns the users for the org whose ID it's passed.
func (c *Client) OrgUsers(orgId int64) ([]OrgUser, error) {
	users := make([]OrgUser, 0)
	err := c.request("GET", fmt.Sprintf("/api/orgs/%d/users", orgId), nil, nil, &users)
	if err != nil {
		return users, err
	}

	return users, err
}

// AddOrgUser adds a user to an org with the specified role.
func (c *Client) AddOrgUser(orgId int64, user, role string) error {
	dataMap := map[string]string{
		"loginOrEmail": user,
		"role":         role,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request("POST", fmt.Sprintf("/api/orgs/%d/users", orgId), nil, bytes.NewBuffer(data), nil)
}

// UpdateOrgUser updates and org user.
func (c *Client) UpdateOrgUser(orgId, userId int64, role string) error {
	dataMap := map[string]string{
		"role": role,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request("PATCH", fmt.Sprintf("/api/orgs/%d/users/%d", orgId, userId), nil, bytes.NewBuffer(data), nil)
}

// RemoveOrgUser removes a user from an org.
func (c *Client) RemoveOrgUser(orgId, userId int64) error {
	return c.request("DELETE", fmt.Sprintf("/api/orgs/%d/users/%d", orgId, userId), nil, nil, nil)
}
