package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Annotation represents a Grafana API Annotation
type Annotation struct {
	ID          int64    `json:"id,omitempty"`
	AlertID     int64    `json:"alertId,omitempty"`
	DashboardID int64    `json:"dashboardId"`
	PanelID     int64    `json:"panelId"`
	UserID      int64    `json:"userId,omitempty"`
	UserName    string   `json:"userName,omitempty"`
	NewState    string   `json:"newState,omitempty"`
	PrevState   string   `json:"prevState,omitempty"`
	Time        int64    `json:"time"`
	TimeEnd     int64    `json:"timeEnd,omitempty"`
	Text        string   `json:"text"`
	Metric      string   `json:"metric,omitempty"`
	RegionID    int64    `json:"regionId,omitempty"`
	Type        string   `json:"type,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	IsRegion    bool     `json:"isRegion,omitempty"`
}

// Annotations fetches the annotations queried with the params it's passed
func (c *Client) Annotations(params map[string]string) ([]Annotation, error) {
	pathAndQuery := buildPathAndQuery("/api/annotations", params)
	req, err := c.newRequest("GET", pathAndQuery, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := []Annotation{}
	err = json.Unmarshal(data, &result)
	return result, err
}

// NewAnnotation creates a new annotation with the Annotation it is passed
func (c *Client) NewAnnotation(a *Annotation) (int64, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return 0, err
	}
	req, err := c.newRequest("POST", "/api/annotations", bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New(resp.Status)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	result := struct {
		ID int64 `json:"id"`
	}{}
	err = json.Unmarshal(data, &result)
	return result.ID, err
}
