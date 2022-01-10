package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// LibraryPanelMetaUser represents the Grafana library panel createdBy and updatedBy fields
type LibraryPanelMetaUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"folderId"`
}

// LibraryPanelMeta represents Grafana library panel metadata.
type LibraryPanelMeta struct {
	FolderName          string               `json:"folderName,,omitempty"`
	FolderUID           string               `json:"folderUid,omitempty"`
	ConnectedDashboards int64                `json:"connectedDashboards,omitempty"`
	Created             time.Time            `json:"created,omitempty"`
	Updated             time.Time            `json:"updated,omitempty"`
	CreatedBy           LibraryPanelMetaUser `json:"createdBy,omitempty"`
	UpdatedBy           LibraryPanelMetaUser `json:"updatedBy,omitempty"`
}

// LibraryPanel represents a Grafana library panel.
type LibraryPanel struct {
	Folder      int64                  `json:"folderId"`
	Name        string                 `json:"name"`
	Model       map[string]interface{} `json:"model"`
	Description string                 `json:"description,omitempty"`
	ID          int64                  `json:"id,omitempty"`
	Kind        int64                  `json:"kind,omitempty"`
	OrgID       int64                  `json:"orgId,omitempty"`
	UID         string                 `json:"uid,omitempty"`
	Version     int64                  `json:"version,omitempty"`
	Meta        LibraryPanelMeta       `json:"meta,omitempty"`
}

// LibraryPanelCreateResponse represents the Grafana API response to creating or saving a library panel.
type LibraryPanelCreateResponse struct {
	Result LibraryPanel `json:"result"`
}

// LibraryPanelDeleteResponse represents the Grafana API response to creating or saving a library panel.
type LibraryPanelDeleteResponse struct {
	Message string `json:"message"`
	ID      int64  `json:"id,omitempty"`
}

// LibraryPanelConnection represents Grafana library panel connections to dashboard.
type LibraryPanelConnection struct {
	ID          int64                `json:"id"`
	Kind        int64                `json:"kind"`
	PanelID     int64                `json:"elementId"`
	DashboardID int64                `json:"connectionId"`
	Created     time.Time            `json:"created"`
	CreatedBy   LibraryPanelMetaUser `json:"createdBy"`
}

// NewLibraryPanel creates a new Grafana library panel.
func (c *Client) NewLibraryPanel(panel LibraryPanel) (*LibraryPanel, error) {
	panel.Kind = int64(1)
	data, err := json.Marshal(panel)
	if err != nil {
		return nil, err
	}

	resp := &LibraryPanelCreateResponse{}
	err = c.request("POST", "/api/library-elements", nil, bytes.NewBuffer(data), &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
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
	resp := &LibraryPanelCreateResponse{}
	err := c.request("GET", path, nil, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
}

// PatchLibraryPanel updates one or more properties of an existing panel that matches the specified UID.
func (c *Client) PatchLibraryPanel(uid string, p *LibraryPanel) (*LibraryPanel, error) {
	path := fmt.Sprintf("/api/library-elements/%s", uid)
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	resp := &LibraryPanelCreateResponse{}
	err = c.request("PATCH", path, nil, bytes.NewBuffer(data), &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
}

// DeleteLibraryPanel deletes a panel by UID.
func (c *Client) DeleteLibraryPanel(uid string) (*LibraryPanelDeleteResponse, error) {
	path := fmt.Sprintf("/api/library-elements/%s", uid)

	resp := &LibraryPanelDeleteResponse{}
	err := c.request("DELETE", path, nil, bytes.NewBuffer(nil), &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// LibraryPanelConnections gets a library panel by name.
func (c *Client) LibraryPanelConnections(uid string) (*[]LibraryPanelConnection, error) {
	path := fmt.Sprintf("/api/library-elements/%s/connections", uid)

	resp := struct {
		Result []LibraryPanelConnection `json:"result"`
	}{}

	// resp := &LibraryPanelConnectionResponse{}
	err := c.request("POST", path, nil, bytes.NewBuffer(nil), &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Result, err
}

// LibraryPanelConnectedDashboards gets a library panel by UID.
func (c *Client) LibraryPanelConnectedDashboards(uid string) (dashboards []*Dashboard, err error) {
	connections, err := c.LibraryPanelConnections(uid)
	if err != nil {
		return nil, err
	}

	var dashboardIds []int64
	for _, this_connection := range *connections {
		dashboardIds = append(dashboardIds, this_connection.DashboardID)
	}

	dashboards, err = c.DashboardsByIDs(dashboardIds)
	if err != nil {
		return nil, err
	}

	return dashboards, err
}
