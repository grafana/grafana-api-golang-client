package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// AlertRule represents a Grafana Alert Rule.
type AlertRule struct {
	Annotations  map[string]string `json:"annotations,omitempty"`
	Condition    string            `json:"condition"`
	Data         []*AlertQuery     `json:"data"`
	ExecErrState ExecErrState      `json:"execErrState"`
	FolderUID    string            `json:"folderUid"`
	ID           int64             `json:"id,omitempty"`
	Labels       map[string]string `json:"labels,omitempty"`
	NoDataState  NoDataState       `json:"noDataState"`
	OrgID        int64             `json:"orgId"`
	RuleGroup    string            `json:"ruleGroup"`
	Title        string            `json:"title"`
	UID          string            `json:"uid,omitempty"`
	Updated      time.Time         `json:"updated"`
	ForDuration  time.Duration     `json:"for"`
	Provenance   string            `json:"provenance"`
}

// RuleGroup represents a group of rules in Grafana Alerting.
type RuleGroup struct {
	Title     string      `json:"title"`
	FolderUID string      `json:"folderUid"`
	Interval  int64       `json:"interval"`
	Rules     []AlertRule `json:"rules"`
}

// AlertQuery represents a single query stage associated with an alert definition.
type AlertQuery struct {
	DatasourceUID     string            `json:"datasourceUid,omitempty"`
	Model             interface{}       `json:"model"`
	QueryType         string            `json:"queryType,omitempty"`
	RefID             string            `json:"refId,omitempty"`
	RelativeTimeRange RelativeTimeRange `json:"relativeTimeRange"`
}

type ExecErrState string
type NoDataState string

const (
	ErrOK          ExecErrState = "OK"
	ErrError       ExecErrState = "Error"
	ErrAlerting    ExecErrState = "Alerting"
	NoDataOk       NoDataState  = "OK"
	NoData         NoDataState  = "NoData"
	NoDataAlerting NoDataState  = "Alerting"
)

// RelativeTimeRange represents the time range for an alert query.
type RelativeTimeRange struct {
	From time.Duration `json:"from"`
	To   time.Duration `json:"to"`
}

// AlertRule fetches a single alert rule, identified by its UID.
func (c *Client) AlertRule(uid string) (AlertRule, error) {
	path := fmt.Sprintf("/api/v1/provisioning/alert-rules/%s", uid)
	result := AlertRule{}
	err := c.request("GET", path, nil, nil, &result)
	if err != nil {
		return AlertRule{}, err
	}
	return result, err
}

// AlertRuleGroup fetches a group of alert rules, identified by its name and the UID of its folder.
func (c *Client) AlertRuleGroup(folderUID string, name string) (RuleGroup, error) {
	path := fmt.Sprintf("/api/v1/provisioning/folder/%s/rule-groups/%s", folderUID, name)
	result := RuleGroup{}
	err := c.request("GET", path, nil, nil, &result)
	return result, err
}

// NewAlertRule creates a new alert rule and returns its UID.
func (c *Client) NewAlertRule(ar *AlertRule) (string, error) {
	req, err := json.Marshal(ar)
	if err != nil {
		return "", err
	}
	result := AlertRule{}
	err = c.request("POST", "/api/v1/provisioning/alert-rules", nil, bytes.NewBuffer(req), &result)
	if err != nil {
		return "", err
	}
	return result.UID, nil
}

// UpdateAlertRule replaces an alert rule, identified by the alert rule's UID.
func (c *Client) UpdateAlertRule(ar *AlertRule) error {
	uri := fmt.Sprintf("/api/v1/provisioning/alert-rules/%s", ar.UID)
	req, err := json.Marshal(ar)
	if err != nil {
		return err
	}

	return c.request("PUT", uri, nil, bytes.NewBuffer(req), nil)
}

// DeleteAlertRule deletes a alert rule, identified by the alert rule's UID.
func (c *Client) DeleteAlertRule(uid string) error {
	uri := fmt.Sprintf("/api/v1/provisioning/alert-rules/%s", uid)
	return c.request("DELETE", uri, nil, nil, nil)
}
