package gapi

import (
	"fmt"
	"net/url"

	jsonx "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

// Example: {From: "now-1h", To: "now"}
type DashboardRelativeTimeRange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// DashboardMeta represents Grafana dashboard meta.
type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Slug      string `json:"slug"`
	Folder    int64  `json:"folderId"`
	FolderUID string `json:"folderUid"`
	URL       string `json:"url"`
}

// DashboardSaveResponse represents the Grafana API response to creating or saving a dashboard.
type DashboardSaveResponse struct {
	Slug    string `json:"slug"`
	ID      int64  `json:"id"`
	UID     string `json:"uid"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

type DashboardPanel struct {
	Targets []struct {
		Expr         string `json:"expr"`
		Interval     string `json:"interval"`
		LegendFormat string `json:"legendFormat"`
		RefID        string `json:"refId"`
	} `json:"targets"`
	Title        string         `json:"title"`
	Type         string         `json:"type"`
	SnapshotData []SnapshotData `json:"snapshotData,omitempty"`
	Unknown      jsontext.Value `json:",unknown"`
}

type DashboardModel struct {
	ID      int              `json:"id"`
	Panels  []DashboardPanel `json:"panels"`
	Time    QueryTimeRange   `json:"time"`
	Title   string           `json:"title"`
	UID     string           `json:"uid"`
	Unknown jsontext.Value   `json:",unknown"`
}

// Dashboard represents a Grafana dashboard.
type Dashboard struct {
	Model    DashboardModel `json:"dashboard"`
	FolderID int64          `json:"folderId"`

	// This field is read-only. It is not used when creating a new dashboard.
	Meta DashboardMeta `json:"meta"`

	// These fields are only used when creating a new dashboard, they will always be empty when getting a dashboard.
	Overwrite bool   `json:"overwrite,omitempty"`
	Message   string `json:"message,omitempty"`
	FolderUID string `json:"folderUid,omitempty"`
}

// SaveDashboard is a deprecated method for saving a Grafana dashboard. Use NewDashboard.
// Deprecated: Use NewDashboard instead.
func (c *Client) SaveDashboard(model map[string]interface{}, overwrite bool) (*DashboardSaveResponse, error) {
	wrapper := map[string]interface{}{
		"dashboard": model,
		"overwrite": overwrite,
	}
	data, err := jsonx.Marshal(wrapper, defaultJSONOptions()...)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = c.request("POST", "/api/dashboards/db", nil, data, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// NewDashboard creates a new Grafana dashboard.
func (c *Client) NewDashboard(dashboard Dashboard) (*DashboardSaveResponse, error) {
	data, err := jsonx.Marshal(dashboard, defaultJSONOptions()...)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = c.request("POST", "/api/dashboards/db", nil, data, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Dashboards fetches and returns all dashboards.
func (c *Client) Dashboards() ([]FolderDashboardSearchResponse, error) {
	const limit = 1000

	var (
		page          = 0
		newDashboards []FolderDashboardSearchResponse
		dashboards    []FolderDashboardSearchResponse
		query         = make(url.Values)
	)

	query.Set("type", "dash-db")
	query.Set("limit", fmt.Sprint(limit))

	for {
		page++
		query.Set("page", fmt.Sprint(page))

		if err := c.request("GET", "/api/search", query, nil, &newDashboards); err != nil {
			return nil, err
		}

		dashboards = append(dashboards, newDashboards...)

		if len(newDashboards) < limit {
			return dashboards, nil
		}
	}
}

// Dashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DashboardByUID instead.
func (c *Client) Dashboard(slug string) (*Dashboard, error) {
	return c.dashboard(fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DashboardByUID gets a dashboard by UID.
func (c *Client) DashboardByUID(uid string) (*Dashboard, error) {
	return c.dashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

// DashboardsByIDs uses the folder and dashboard search endpoint to find
// dashboards by list of dashboard IDs.
func (c *Client) DashboardsByIDs(ids []int64) ([]FolderDashboardSearchResponse, error) {
	dashboardIdsJSON, err := jsonx.Marshal(ids, defaultJSONOptions()...)
	if err != nil {
		return nil, err
	}

	params := url.Values{
		"type":         {"dash-db"},
		"dashboardIds": {string(dashboardIdsJSON)},
	}
	return c.FolderDashboardSearch(params)
}

func (c *Client) dashboard(path string) (*Dashboard, error) {
	result := &Dashboard{}
	err := c.request("GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	result.FolderID = result.Meta.Folder

	return result, err
}

// DeleteDashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DeleteDashboardByUID instead.
func (c *Client) DeleteDashboard(slug string) error {
	return c.deleteDashboard(fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DeleteDashboardByUID deletes a dashboard by UID.
func (c *Client) DeleteDashboardByUID(uid string) error {
	return c.deleteDashboard(fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

func (c *Client) deleteDashboard(path string) error {
	return c.request("DELETE", path, nil, nil, nil)
}
