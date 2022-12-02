package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DashboardVersion is the data returned from the DashboardVersion* methods.
type DashboardVersion struct {
	ID            int64  `json:"id"`
	DashboardID   int64  `json:"dashboardId"`
	UID           string `json:"uid"`
	ParentVersion int64  `json:"parentVersion"`
	RestoredFrom  int64  `json:"restoredFrom"`
	Version       int64  `json:"version"`
	Created       string `json:"created"`
	CreatedBy     string `json:"createdBy"`
	Message       string `json:"message"`
	// Data is the dashboard model, and only gets filled when a single version is requested.
	Data map[string]interface{} `json:"data"`
}

// DashboardVersionRestore is the returned data from restoring a dashboard.
// slug = The URL friendly slug of the dashboard’s title,
// status = Whether the restoration was successful or not,
// version = The new dashboard version, following the restoration.
type DashboardVersionRestore struct {
	ID      int64  `json:"id"`      // 70
	Slug    string `json:"slug"`    // my-dashboard
	Status  string `json:"status"`  // success
	UID     string `json:"uid"`     // QA7wKklGz
	URL     string `json:"url"`     // /d/QA7wKklGz/my-dashboard
	Version int64  `json:"version"` // 3
}

// CompareDashboardsInput is the required input when comparing dashboards.
type CompareDashboardsInput struct {
	BaseDashboardID      int64
	BaseDashboardVersion int64
	NewDashboardID       int64
	NewDashboardVersion  int64
	// DiffType may be 'json' or 'basic'. Default is 'basic'.
	DiffType string
}

// DashboardVersions returns all dashboard versions for a specific dashboard UID.
// limit = Maximum number of results to return,
// start = Version to start from when returning queries.
func (c *Client) DashboardVersions(dashboardUID string, limit, start int64) ([]*DashboardVersion, error) {
	var (
		params = make(url.Values)
		result = []*DashboardVersion{}
		path   = fmt.Sprintf("/api/dashboards/uid/%s/versions", dashboardUID)
	)

	params.Set("limit", fmt.Sprint(limit))
	params.Set("start", fmt.Sprint(start))

	if err := c.request(http.MethodGet, path, params, nil, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DashboardVersion returns a single dashboard version for a specific dashboard UID.
// version = The version number to return. Empty data is returned if it does not exist.
func (c *Client) DashboardVersion(dashboardUID string, version int64) (*DashboardVersion, error) {
	var (
		params = make(url.Values)
		result = DashboardVersion{}
		path   = fmt.Sprintf("/api/dashboards/uid/%s/versions/%d", dashboardUID, version)
	)

	if err := c.request(http.MethodGet, path, params, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RestoreDashboardVersion restores a dashboard version for a specific dashboard UID.
// version = The version number to restore.
func (c *Client) RestoreDashboardVersion(dashboardUID string, version int64) (*DashboardVersionRestore, error) {
	var (
		params = make(url.Values)
		result = DashboardVersionRestore{}
		path   = fmt.Sprintf("/api/dashboards/uid/%s/restore", dashboardUID)
		body   = bytes.NewBufferString(fmt.Sprintf(`{"version":%d}`, version))
	)

	if err := c.request(http.MethodPost, path, params, body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CompareDashboardVersions compares two different dashboard versions. Can even compare different dashboards IDs.
// Returns "text" (formatted html) not json. You can wrap the []byte in string() to convert it.
func (c *Client) CompareDashboardVersions(compareDashboards CompareDashboardsInput) ([]byte, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(compareDashboards); err != nil {
		return nil, fmt.Errorf("json encoding input: %w", err)
	}

	req, err := c.newRequest(http.MethodPost, "/api/dashboards/calculate-diff", url.Values{}, &body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body: %w", err)
	}

	// check status code.
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("%w: %d, body: %v", ErrInvalidStatus, resp.StatusCode, string(bodyContents))
	}

	return bodyContents, nil
}

// MarshalJSON turns a CompareDashboardsInput into the json Grafana requires.
func (c CompareDashboardsInput) MarshalJSON() ([]byte, error) {
	type dashboard struct {
		ID      int64 `json:"dashboardId"`
		Version int64 `json:"version"`
	}

	diffType := c.DiffType
	if diffType == "" {
		diffType = "basic"
	}

	return json.Marshal(&struct {
		Base dashboard `json:"base"`
		New  dashboard `json:"new"`
		Type string    `json:"diffType"`
	}{
		Base: dashboard{ID: c.BaseDashboardID, Version: c.BaseDashboardVersion},
		New:  dashboard{ID: c.NewDashboardID, Version: c.NewDashboardVersion},
		Type: diffType,
	})
}
