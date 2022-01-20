package gapi

import (
	"testing"
)

const (
	getStackJSON    = `{"id": 1,"slug": "mystack"}`
	createStackJSON = `{"id": 1,"slug": "mystack"}`
	getStacksJSON   = `
{
	"items": [
		{
			"id": 1,
			"orgId": 1,
			"orgSlug": "myorg",
			"orgName": "MyOrg",
			"name": "mystack.grafana.net",
			"url": "https://mystack.grafana.net",
			"slug": "mystack",
			"version": "stable",
			"description": "My amazing stack",
			"status": "active",
			"gateway": "istio",
			"createdAt": "2021-12-22T14:02:46.000Z",
			"createdBy": "xyz",
			"updatedAt": null,
			"updatedBy": "",
			"trial": 0,
			"trialExpiresAt": null,
			"clusterId": 666,
			"clusterSlug": "prod-eu-west-0",
			"clusterName": "prod-eu-west-0",
			"plan": "gcloud",
			"planName": "Grafana Cloud",
			"billingStartDate": "2021-12-22T14:02:46.000Z",
			"billingEndDate": null,
			"billingActiveUsers": 0,
			"currentActiveUsers": 1,
			"currentActiveAdminUsers": 1,
			"currentActiveEditorUsers": 0,
			"currentActiveViewerUsers": 0,
			"dailyUserCnt": 0,
			"dailyAdminCnt": 0,
			"dailyEditorCnt": 0,
			"dailyViewerCnt": 0,
			"billableUserCnt": 1,
			"dashboardCnt": 6,
			"datasourceCnts": {},
			"userQuota": 10,
			"dashboardQuota": -1,
			"alertQuota": -1,
			"ssl": true,
			"customAuth": true,
			"customDomain": true,
			"support": true,
			"runningVersion": "8.3.3 (commit: 30bb7a93c, branch: HEAD)",
			"machineLearning": 0,
			"hmInstancePromId": 112233,
			"hmInstancePromUrl": "https://prometheus-prod-01-eu-west-0.grafana.net",
			"hmInstancePromName": "mystack-prom",
			"hmInstancePromStatus": "active",
			"hmInstancePromCurrentUsage": 11111,
			"hmInstancePromCurrentActiveSeries": 222222,
			"hmInstanceGraphiteId": 111111,
			"hmInstanceGraphiteUrl": "https://graphite-prod-01-eu-west-0.grafana.net",
			"hmInstanceGraphiteName": "mystack-graphite",
			"hmInstanceGraphiteType": "graphite-v5",
			"hmInstanceGraphiteStatus": "active",
			"hmInstanceGraphiteCurrentUsage": 0,
			"hlInstanceId": 121416,
			"hlInstanceUrl": "https://logs-prod-eu-west-0.grafana.net",
			"hlInstanceName": "mystack-logs",
			"hlInstanceStatus": "active",
			"hlInstanceCurrentUsage": 0,
			"amInstanceId": 654321,
			"amInstanceName": "mystack-alerts",
			"amInstanceUrl": "https://alertmanager-eu-west-0.grafana.net",
			"amInstanceStatus": "active",
			"amInstanceGeneratorUrl": "https://mystack.grafana.net",
			"htInstanceId": 123456,
			"htInstanceUrl": "https://tempo-eu-west-0.grafana.net",
			"htInstanceName": "mystack-traces",
			"htInstanceStatus": "active",
			"regionId": 3,
			"regionSlug": "eu"
			}
	]
}
`
)

func TestStacks(t *testing.T) {

	server, client := gapiTestTools(t, 200, getStacksJSON)
	defer server.Close()

	stacks, err := client.Stacks()

	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	actualItemCound := len(stacks.Items)
	expectedItemCount := 1

	if actualItemCound != expectedItemCount {
		t.Errorf("Length of returned stacks - Actual Stacks Count: %d, Expected Stacks Count: %d", actualItemCound, expectedItemCount)
	}

	actualStackId := stacks.Items[0].ID
	expectedStackId := int64(1)

	if actualStackId != expectedStackId {
		t.Errorf("Unexpected Stack ID - Actual Stack ID: %d, Expected Stack ID: %d", actualStackId, expectedStackId)
	}

	actualSlug := stacks.Items[0].Slug
	expectedSlug := "mystack"
	if actualSlug != expectedSlug {
		t.Errorf("Unexpected Stack Slug - Actual Slug: %d, Expected Slug: %d", actualStackId, expectedStackId)
	}

}

func TestCreateStack(t *testing.T) {

	server, client := gapiTestTools(t, 200, createStackJSON)
	defer server.Close()

	stackId, err := client.NewStack("mystack", "mystack", "eu")

	if err != nil {
		t.Fatal(err)
	}

	expectedStackId := int64(1)
	actualStackId := stackId

	if actualStackId != expectedStackId {
		t.Errorf("Unexpected Stack ID - Actual: %d, Expected: %d", actualStackId, expectedStackId)
	}
}

func TestStackBySlug(t *testing.T) {

	server, client := gapiTestTools(t, 200, getStackJSON)
	defer server.Close()

	stack := "mystack"
	resp, err := client.StackBySlug(stack)
	if err != nil {
		t.Fatal(err)
	}

	expectedStack := stack
	actualStack := resp.Slug

	if actualStack != expectedStack {
		t.Errorf("Unexpected Stack Slug - Actual: %s, Expected: %s", actualStack, expectedStack)
	}
}

func TestStackByID(t *testing.T) {

	server, client := gapiTestTools(t, 200, getStackJSON)
	defer server.Close()

	expectedStackId := int64(1)
	resp, err := client.StackByID(expectedStackId)

	if err != nil {
		t.Fatal(err)
	}

	actualStackId := resp.ID

	if actualStackId != expectedStackId {
		t.Errorf("Unexpected Stack ID - Actual: %d, Expected: %d", actualStackId, expectedStackId)
	}
}

func TestUpdateStack(t *testing.T) {

	server, client := gapiTestTools(t, 200, getStacksJSON)
	defer server.Close()

	errr := client.UpdateStack(1, "mystack-update", "This is a test stack")
	if errr != nil {
		t.Error(errr)
	}
}

func TestDeleteStack(t *testing.T) {

	server, client := gapiTestTools(t, 200, getStacksJSON)
	defer server.Close()

	err := client.DeleteStack("mystack")

	// The DELETE api returns an error so check if there is an error
	if err != nil {
		t.Errorf("Unexpected error - Actual: %s, Expected: nil", err.Error())
	}
}
