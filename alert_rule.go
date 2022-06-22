package gapi

import (
	"fmt"
	"time"
)

// AlertRule represents a Grafana alert rule.
type AlertRule struct {
	UID          string                   `json:"uid"`
	ID           int64                    `json:"id"`
	Title        string                   `json:"title"`
	OrgID        int64                    `json:"orgID"`
	Updated      time.Time                `json:"updated"`
	RuleGroup    string                   `json:"ruleGroup"`
	FolderUID    string                   `json:"folderUID"`
	For          int64                    `json:"for"`
	NoDataState  string                   `json:"noDataState"`
	ExecErrState string                   `json:"execErrState"`
	Condition    string                   `json:"condition"`
	Data         []map[string]interface{} `json:"data"`
	Labels       map[string]string        `json:"labels"`
	Annotations  map[string]string        `json:"annotations"`
	Provenance   string                   `json:"provenance"`
}

// AlertRule fetches a single rule, identified by its UID.
func (c *Client) AlertRule(UID string) (AlertRule, error) {
	rule := AlertRule{}
	uri := fmt.Sprintf("/api/v1/provisioning/alert-rules/%s", UID)
	err := c.request("GET", uri, nil, nil, &rule)
	if err != nil {
		return AlertRule{}, err
	}
	return rule, nil
}
