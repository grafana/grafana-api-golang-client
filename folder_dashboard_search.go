package gapi

import (
	"net/url"
)

// FolderDashboardSearchResponse represents the Grafana API dashboard search response.
type FolderDashboardSearchResponse struct {
	ID          uint     `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URI         string   `json:"uri"`
	URL         string   `json:"url"`
	Slug        string   `json:"slug"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	FolderID    uint     `json:"folderId"`
	FolderUID   string   `json:"folderUid"`
	FolderTitle string   `json:"folderTitle"`
	FolderURL   string   `json:"folderUrl"`
}

// FolderDashboardSearch uses the folder and dashboard search endpoint to find
// dashboards based on the params passed in.
func (c *Client) FolderDashboardSearch(params map[string]interface{}) (resp []FolderDashboardSearchResponse, err error) {
	query := url.Values{}
	for p, v := range params {
		switch c := v.(type) {
		case []string:
			for _, listElement := range c {
				query.Add(p, listElement)
			}
		case string:
			query.Add(p, c)
		}
	}
	err = c.request("GET", "/api/search", query, nil, &resp)
	return
}
