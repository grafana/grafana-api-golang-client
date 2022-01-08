package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// LibraryPanelMetaUser represents the Grafana library panel createdBy and updatedBy fields
type LibraryPanelMetaUser struct {
	ID			  int64		`json:"id"`
	Name			string  `json:"name"`
	AvatarURL	string	`json:"folderId"`
}

// LibraryPanelMeta represents Grafana library panel metadata.
type LibraryPanelMeta struct {
	FolderName 					string 								`json:"folderName,,omitempty"`
	FolderUid 					string 								`json:"folderUid,omitempty"`
	ConnectedDashboards	int64  								`json:"connectedDashboards,omitempty"`
	Created							time.Time 						`json:"created,omitempty"`
	Updated							time.Time 						`json:"updated,omitempty"`
	CreatedBy						LibraryPanelMetaUser	`json:"createdBy,omitempty"`
	UpdatedBy						LibraryPanelMetaUser	`json:"updatedBy,omitempty"`
}

// LibraryPanel represents a Grafana library panel.
type LibraryPanel struct {
	Folder			int64  									`json:"folderId"`
	Name 				string 									`json:"name"`
	Model				map[string]interface{}	`json:"model"`
	Description string 									`json:"description,omitempty"`
	ID      		int64  									`json:"id,omitempty"`
	Kind				int64  									`json:"kind,omitempty"`
	OrgID				int64  									`json:"orgId,omitempty"`
	UID     		string 									`json:"uid,omitempty"`
	Version 		int64  									`json:"version,omitempty"`
	Meta 				LibraryPanelMeta				`json:"meta,omitempty"`
}

// LibraryPanelCreateResponse represents the Grafana API response to creating or saving a library panel.
type LibraryPanelCreateResponse struct {
	Result	LibraryPanel	`json:"result"`
}

// LibraryPanelDeleteResponse represents the Grafana API response to creating or saving a library panel.
type LibraryPanelDeleteResponse struct {
	Message string `json:"message"`
	ID 			int64	 `json:"id"`
}

// NewLibraryPanel creates a new Grafana library panel.
func (c *Client) NewLibraryPanel(panel LibraryPanel) (*LibraryPanelCreateResponse, error) {
	panel.Kind := int64(1)  // library element. 1 for library panel or 2 for library variable
	data, err := json.Marshal(panel)
	if err != nil {
		return nil, err
	}

	result := &LibraryPanelCreateResponse{}
	err = c.request("POST", "/api/library-elements", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// LibraryPanelByUID gets a library panel by UID.
func (c *Client) LibraryPanelByUID(uid string) (*LibraryPanel, error) {
	return c.panel(fmt.Sprintf("/api/library-elements/%s", uid))
}

// LibraryPanelByName gets a library panel by name.
func (c *Client) LibraryPanelByName(name string) (*LibraryPanel, error) {
	return c.panel(fmt.Sprintf("/api/library-elements/name/%s", name))
}

func (c *Client) panel(path string) (*LibraryPanel, error) {
	result := &LibraryPanel{}
	err := c.request("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// PatchLibraryPanelByUID updates one or more properties of an existing panel that matches the specified UID.
func (c *Client) PatchLibraryPanelByUID(uid string, p *LibraryPanel) (string, error) {
	path := fmt.Sprintf("/api/library-elements/%s", uid)
	data, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	result := &LibraryPanelCreateResponse{}
	err = c.request("PATCH", path, nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return "", err
	}

	return result, err
}

// DeleteLibraryPanelByUID deletes a panel by UID.
func (c *Client) DeleteLibraryPanelByUID(uid string) error {
	path := fmt.Sprintf("/api/library-elements/%s", uid)

	result := &LibraryPanelDeleteResponse{}
	err := c.request("DELETE", path, nil, bytes.NewBuffer(nil), &result)
	if err != nil {
		return "", err
	}

	return result, err
}
