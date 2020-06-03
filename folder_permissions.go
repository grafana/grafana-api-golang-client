package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// FolderPermission has information such as a folder, user, team, role and permission.
type FolderPermission struct {
	Id        int64  `json:"id"`
	FolderUid string `json:"uid"`
	UserId    int64  `json:"userId"`
	TeamId    int64  `json:"teamId"`
	Role      string `json:"role"`
	IsFolder  bool   `json:"isFolder"`

	// Permission levels are
	// 1 = View
	// 2 = Edit
	// 4 = Admin
	Permission     int64  `json:"permission"`
	PermissionName string `json:"permissionName"`

	// optional fields
	FolderId    int64 `json:"folderId,omitempty"`
	DashboardId int64 `json:"dashboardId,omitempty"`
}

type PermissionItems struct {
	Items []*PermissionItem `json:"items"`
}

type PermissionItem struct {
	// As you can see the docs, each item has a pair of [Role|TeamId|UserId] and Permission.
	// unnecessary fields are omitted.
	Role       string `json:"role,omitempty"`
	TeamId     int64  `json:"teamId,omitempty"`
	UserId     int64  `json:"userId,omitempty"`
	Permission int64  `json:"permission"`
}

func (c *Client) FolderPermissions(fid string) ([]*FolderPermission, error) {
	permissions := make([]*FolderPermission, 0)

	req, err := c.newRequest("GET", fmt.Sprintf("/api/folders/%s/permissions", fid), nil, nil)
	if err != nil {
		return permissions, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return permissions, err
	}
	if resp.StatusCode != 200 {
		return permissions, fmt.Errorf(resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return permissions, err
	}
	if err := json.Unmarshal(data, &permissions); err != nil {
		return permissions, err
	}
	return permissions, nil
}

// UpdateFolderPermissions remove existing permissions if items are not included in the request.
func (c *Client) UpdateFolderPermissions(fid string, items *PermissionItems) error {
	path := fmt.Sprintf("/api/folders/%s/permissions", fid)
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	req, err := c.newRequest("POST", path, nil, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}
	return nil
}
