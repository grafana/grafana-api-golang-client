package gapi

import (
	"encoding/json"
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

	var expectedStacks StackItems

	err = UnmarshalJSONToStruct(getStacksJSON, &expectedStacks)
	if err != nil {
		t.Fatal(err)
	}
	actualItemCount := len(stacks.Items)
	expectedItemCount := len(expectedStacks.Items)

	// check that the number of items is the same
	if actualItemCount != expectedItemCount {
		t.Errorf("Length of returned stacks - Actual Stacks Count: %d, Expected Stacks Count: %d", actualItemCount, expectedItemCount)
	}

	// Check ID of the returned Stack is as expected
	actualStackID := stacks.Items[0].ID
	expectedStackID := expectedStacks.Items[0].ID

	if actualStackID != expectedStackID {
		t.Errorf("Unexpected Stack ID - Actual Stack ID: %d, Expected Stack ID: %d", actualStackID, expectedStackID)
	}

	// Check the slug of the returned stack as expected
	actualSlug := stacks.Items[0].Slug
	expectedSlug := "mystack"
	if actualSlug != expectedSlug {
		t.Errorf("Unexpected Stack Slug - Actual Slug: %d, Expected Slug: %d", actualStackID, expectedStackID)
	}
}

func TestCreateStack(t *testing.T) {
	server, client := gapiTestTools(t, 200, createStackJSON)
	defer server.Close()

	actualStackID, err := client.NewStack("mystack", "mystack", "eu")

	if err != nil {
		t.Fatal(err)
	}

	var expectedStack Stack
	err = UnmarshalJSONToStruct(createStackJSON, &expectedStack)
	if err != nil {
		t.Fatal(err)
	}

	if actualStackID != expectedStack.ID {
		t.Errorf("Unexpected Stack ID - Actual: %d, Expected: %d", actualStackID, expectedStack.ID)
	}
}

func TestStackBySlug(t *testing.T) {
	server, client := gapiTestTools(t, 200, getStackJSON)
	defer server.Close()

	expectedStackSlug := "mystack"
	resp, err := client.StackBySlug(expectedStackSlug)
	if err != nil {
		t.Fatal(err)
	}

	actualStackSlug := resp.Slug

	if actualStackSlug != expectedStackSlug {
		t.Errorf("Unexpected Stack Slug - Actual: %s, Expected: %s", actualStackSlug, expectedStackSlug)
	}
}

func TestStackByID(t *testing.T) {
	server, client := gapiTestTools(t, 200, getStackJSON)
	defer server.Close()

	expectedStackID := int64(1)
	resp, err := client.StackByID(expectedStackID)

	if err != nil {
		t.Fatal(err)
	}

	actualStackID := resp.ID

	if actualStackID != expectedStackID {
		t.Errorf("Unexpected Stack ID - Actual: %d, Expected: %d", actualStackID, expectedStackID)
	}
}

func TestUpdateStack(t *testing.T) {
	server, client := gapiTestTools(t, 200, getStacksJSON)
	defer server.Close()

	err := client.UpdateStack(1, "mystack2", "mystack2", "This is a test stack update")
	if err != nil {
		t.Error(err)
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

func UnmarshalJSONToStruct(jsonString string, target interface{}) error {
	err := json.Unmarshal([]byte(jsonString), &target)
	if err != nil {
		return err
	}

	return nil
}
