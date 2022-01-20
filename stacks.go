package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Stack struct {
	ID                       int64       `json:"id"`
	OrgID                    int64       `json:"orgId"`
	OrgSlug                  string      `json:"orgSlug"`
	OrgName                  string      `json:"orgName"`
	Name                     string      `json:"name"`
	URL                      string      `json:"url"`
	Slug                     string      `json:"slug"`
	Version                  string      `json:"version"`
	Description              string      `json:"description"`
	Status                   string      `json:"status"`
	Gateway                  string      `json:"gateway"`
	CreatedAt                time.Time   `json:"createdAt"`
	CreatedBy                string      `json:"createdBy"`
	UpdatedAt                interface{} `json:"updatedAt"`
	UpdatedBy                string      `json:"updatedBy"`
	Trial                    int         `json:"trial"`
	TrialExpiresAt           interface{} `json:"trialExpiresAt"`
	ClusterID                int         `json:"clusterId"`
	ClusterSlug              string      `json:"clusterSlug"`
	ClusterName              string      `json:"clusterName"`
	Plan                     string      `json:"plan"`
	PlanName                 string      `json:"planName"`
	BillingStartDate         time.Time   `json:"billingStartDate"`
	BillingEndDate           interface{} `json:"billingEndDate"`
	BillingActiveUsers       int         `json:"billingActiveUsers"`
	CurrentActiveUsers       int         `json:"currentActiveUsers"`
	CurrentActiveAdminUsers  int         `json:"currentActiveAdminUsers"`
	CurrentActiveEditorUsers int         `json:"currentActiveEditorUsers"`
	CurrentActiveViewerUsers int         `json:"currentActiveViewerUsers"`
	DailyUserCnt             int         `json:"dailyUserCnt"`
	DailyAdminCnt            int         `json:"dailyAdminCnt"`
	DailyEditorCnt           int         `json:"dailyEditorCnt"`
	DailyViewerCnt           int         `json:"dailyViewerCnt"`
	BillableUserCnt          int         `json:"billableUserCnt"`
	DashboardCnt             int         `json:"dashboardCnt"`
	DatasourceCnts           struct {
	} `json:"datasourceCnts"`
	UserQuota                         int    `json:"userQuota"`
	DashboardQuota                    int    `json:"dashboardQuota"`
	AlertQuota                        int    `json:"alertQuota"`
	Ssl                               bool   `json:"ssl"`
	CustomAuth                        bool   `json:"customAuth"`
	CustomDomain                      bool   `json:"customDomain"`
	Support                           bool   `json:"support"`
	RunningVersion                    string `json:"runningVersion"`
	MachineLearning                   int    `json:"machineLearning"`
	HmInstancePromID                  int    `json:"hmInstancePromId"`
	HmInstancePromURL                 string `json:"hmInstancePromUrl"`
	HmInstancePromName                string `json:"hmInstancePromName"`
	HmInstancePromStatus              string `json:"hmInstancePromStatus"`
	HmInstancePromCurrentUsage        int    `json:"hmInstancePromCurrentUsage"`
	HmInstancePromCurrentActiveSeries int    `json:"hmInstancePromCurrentActiveSeries"`
	HmInstanceGraphiteID              int    `json:"hmInstanceGraphiteId"`
	HmInstanceGraphiteURL             string `json:"hmInstanceGraphiteUrl"`
	HmInstanceGraphiteName            string `json:"hmInstanceGraphiteName"`
	HmInstanceGraphiteType            string `json:"hmInstanceGraphiteType"`
	HmInstanceGraphiteStatus          string `json:"hmInstanceGraphiteStatus"`
	HmInstanceGraphiteCurrentUsage    int    `json:"hmInstanceGraphiteCurrentUsage"`
	HlInstanceID                      int    `json:"hlInstanceId"`
	HlInstanceURL                     string `json:"hlInstanceUrl"`
	HlInstanceName                    string `json:"hlInstanceName"`
	HlInstanceStatus                  string `json:"hlInstanceStatus"`
	HlInstanceCurrentUsage            int    `json:"hlInstanceCurrentUsage"`
	AmInstanceID                      int    `json:"amInstanceId"`
	AmInstanceName                    string `json:"amInstanceName"`
	AmInstanceURL                     string `json:"amInstanceUrl"`
	AmInstanceStatus                  string `json:"amInstanceStatus"`
	AmInstanceGeneratorURL            string `json:"amInstanceGeneratorUrl"`
	HtInstanceID                      int    `json:"htInstanceId"`
	HtInstanceURL                     string `json:"htInstanceUrl"`
	HtInstanceName                    string `json:"htInstanceName"`
	HtInstanceStatus                  string `json:"htInstanceStatus"`
	RegionID                          int    `json:"regionId"`
	RegionSlug                        string `json:"regionSlug"`
	Links                             []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
}

// StackItems represents Grafana stack items.
type StackItems struct {
	Items []*Stack `json:"items"`
}

// Stacks fetches and returns the Grafana stacks.
func (c *Client) Stacks() (StackItems, error) {
	stacks := StackItems{}
	err := c.request("GET", "/api/instances", nil, nil, &stacks)
	if err != nil {
		return stacks, err
	}

	return stacks, err
}

// StackByName fetches and returns the stack whose slug it's passed.
func (c *Client) StackBySlug(slug string) (Stack, error) {
	stack := Stack{}
	err := c.request("GET", fmt.Sprintf("/api/instances/%s", slug), nil, nil, &stack)

	if err != nil {
		return stack, err
	}

	return stack, err
}

// StackByID fetches and returns the stack whose name it's passed.
// This returns deleted instances as well with `status=deleted`.
func (c *Client) StackByID(id int64) (Stack, error) {
	stack := Stack{}
	err := c.request("GET", fmt.Sprintf("/api/instances/%d", id), nil, nil, &stack)

	if err != nil {
		return stack, err
	}

	// If we are getting a deleted stack then return an empty stack
	if stack.Status == "deleted" {
		return Stack{}, err
	}

	return stack, err
}

// NewStack creates a new Grafana Stack
func (c *Client) NewStack(stackName string, stackSlug string, region string) (int64, error) {
	id := int64(0)

	// Check if this stack already exists
	stack, stackErr := c.StackBySlug(stackSlug)

	if stackErr != nil {
		return id, nil
	}

	// If the stack already exist then return the existing ID
	// There is currently no API for updating a stack so if one exists already we just return its ID
	if stack.ID != 0 {
		return stack.ID, nil
	}

	dataMap := map[string]string{
		"name":   stackName,
		"slug":   stackSlug,
		"region": region,
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return id, err
	}

	result := struct {
		ID int64 `json:"id"`
	}{}

	err = c.request("POST", "/api/instances", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return 0, err
	}

	return result.ID, err
}

// UpdateOrg updates a Grafana stack.
// Only the name can be updated. No other parameters of the stack are updatable
func (c *Client) UpdateStack(id int64, name string, description string) error {
	dataMap := map[string]string{
		"name":        name,
		"description": description,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request("POST", fmt.Sprintf("/api/instances/%d", id), nil, bytes.NewBuffer(data), nil)
}

// DeleteStack deletes the Grafana stack whose slug it passed in.
func (c *Client) DeleteStack(stackSlug string) error {
	return c.request("DELETE", fmt.Sprintf("/api/instances/%s", stackSlug), nil, nil, nil)
}
