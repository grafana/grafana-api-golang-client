package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// DatasourcePermission has information such as a datasource, user, team, role and permission.
type DatasourcePermission struct {
	DatasourceID int64  `json:"datasourceId"`
	UserID       int64  `json:"userId"`
	UserEmail    string `json:"userEmail"`
	TeamID       int64  `json:"teamId"`

	// Permission levels are
	// 1 = Query
	Permission     int64  `json:"permission"`
	PermissionName string `json:"permissionName"`
}

type DatasourcePermissionsResponse struct {
	DatasourceID int64 `json:"datasourceId"`
	Enabled      bool  `json:"enabled"`
	Permissions  []*DatasourcePermission
}

type DatasourcePermissionAddPayload struct {
	UserID     int64 `json:"userId"`
	TeamID     int64 `json:"teamId"`
	Permission int64 `json:"permission"`
}

// DatasourcePermissions fetches and returns the permissions for the datasource whose ID it's passed.
func (c *Client) DatasourcePermissions(id int64) (*DatasourcePermissionsResponse, error) {
	var out *DatasourcePermissionsResponse
	err := c.request("GET", fmt.Sprintf("/api/datasources/id/%d/permissions", id), nil, nil, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

// AddDatasourcePermission adds the given permission item
func (c *Client) AddDatasourcePermission(id int64, item *DatasourcePermissionAddPayload) error {
	path := fmt.Sprintf("/api/datasources/id/%d/permissions", id)
	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return c.request("POST", path, nil, bytes.NewBuffer(data), nil)
}

// RemoveDatasourcePermission removes the permission with the given id
func (c *Client) RemoveDatasourcePermission(id, permissionID int64) error {
	path := fmt.Sprintf("/api/datasources/id/%d/permissions/%d", id, permissionID)

	return c.request("DELETE", path, nil, nil, nil)
}
