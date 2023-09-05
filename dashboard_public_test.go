package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	createPublicDashboard = `{
	"uid": "fdc8b8fd-72cb-45d2-927a-75900e4f6e70",
    "dashboardUid": "nErXDvCkzz",
    "isEnabled": true,
    "share": "public"
}`
	updatePublicDashboard = `{
    "timeSelectionEnabled": true,
    "isEnabled": true,
    "annotationsEnabled": true
}`
	publicDashboardByUID = `{
    "uid": "cd56d9fd-f3d4-486d-afba-a21760e2acbe",
    "dashboardUid": "xCpsVuc4z",
    "accessToken": "5c948bf96e6a4b13bd91975f9a2028b7",
    "createdBy": 1,
    "updatedBy": 1,
    "createdAt": "2023-09-05T11:41:14-03:00",
    "updatedAt": "2023-09-05T11:41:14-03:00",
    "timeSelectionEnabled": false,
    "isEnabled": true,
    "annotationsEnabled": false,
    "share": "public"
}`
	publicDashboardList = `{
    "publicDashboards": [
        {
            "uid": "e9f29a3c-fcc3-4fc5-a690-ae39c97d24ba",
            "accessToken": "6c13ec1997ba48c5af8c9c5079049692",
            "title": "A Datasource not found query",
            "dashboardUid": "d2f21d0a-76c7-47ec-b5f3-9dda16e5a996",
            "isEnabled": true
        },
        {
            "uid": "a174f604-6fe7-47de-97b4-48b7e401b540",
            "accessToken": "d1fcff345c0f45e8a78c096c9696034a",
            "title": "A Issue heatmap bargauge panel",
            "dashboardUid": "51DiOw0Vz",
            "isEnabled": true
        }
    ],
    "totalCount": 2,
    "page": 1,
    "perPage": 1000
}`
)

func TestNewPublicDashboard(t *testing.T) {
	const dashboardUID = "nErXDvCkzz"

	isEnabled := true

	client := gapiTestTools(t, 200, createPublicDashboard)

	publicDashboard := PublicDashboardPayload{
		UID:         "fdc8b8fd-72cb-45d2-927a-75900e4f6e70",
		AccessToken: "b1d5f3f534d84375a897f3969b6157f3",
		IsEnabled:   &isEnabled,
		Share:       "public",
	}

	resp, err := client.NewPublicDashboard(dashboardUID, publicDashboard)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.UID != "fdc8b8fd-72cb-45d2-927a-75900e4f6e70" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "fdc8b8fd-72cb-45d2-927a-75900e4f6e70")
	}

	if resp.DashboardUID != dashboardUID {
		t.Errorf("Invalid dashboard uid - %s, Expected %s", resp.DashboardUID, dashboardUID)
	}
}

func TestDeletePublicDashboard(t *testing.T) {
	client := gapiTestTools(t, 200, "")

	err := client.DeletePublicDashboard("nErXDvCkza", "fdc8b8fd-72cb-45d2-927a-75900e4f6e70")
	if err != nil {
		t.Error(err)
	}
}

func TestPublicDashboards(t *testing.T) {
	client := gapiTestTools(t, 200, publicDashboardList)

	resp, err := client.PublicDashboards()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if len(resp.PublicDashboards) != 2 || resp.TotalCount != 2 {
		t.Error("Length of returned public dashboards should be 2")
	}
	if resp.PublicDashboards[0].UID != "e9f29a3c-fcc3-4fc5-a690-ae39c97d24ba" || resp.PublicDashboards[0].AccessToken != "6c13ec1997ba48c5af8c9c5079049692" {
		t.Error("Not correctly parsing returned public dashboards.")
	}
}

func TestPublicDashboardByUID(t *testing.T) {
	client := gapiTestTools(t, 200, publicDashboardByUID)

	resp, err := client.PublicDashboardbyUID("xCpsVuc4z")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp.UID != "cd56d9fd-f3d4-486d-afba-a21760e2acbe" {
		t.Errorf("Invalid uid - %s, Expected %s", resp.UID, "cd56d9fd-f3d4-486d-afba-a21760e2acbe")
	}

	if resp.DashboardUID != "xCpsVuc4z" {
		t.Errorf("Invalid dashboard uid - %s, Expected %s", resp.DashboardUID, "xCpsVuc4z")
	}
}

func TestUpdatePublicDashboard(t *testing.T) {
	client := gapiTestTools(t, 200, updatePublicDashboard)
	trueVal := true

	publicDashboard := PublicDashboardPayload{
		IsEnabled:            &trueVal,
		TimeSelectionEnabled: &trueVal,
		AnnotationsEnabled:   &trueVal,
	}

	resp, err := client.UpdatePublicDashboard("xCpsVuc4z", "cd56d9fd-f3d4-486d-afba-a21760e2acbe", publicDashboard)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if !resp.IsEnabled {
		t.Errorf("Invalid IsEnabled - %t, Expected %t", resp.IsEnabled, trueVal)
	}

	if !resp.TimeSelectionEnabled {
		t.Errorf("Invalid TimeSelectionEnabled - %t, Expected %t", resp.TimeSelectionEnabled, trueVal)
	}

	if !resp.AnnotationsEnabled {
		t.Errorf("Invalid AnnotationsEnabled - %t, Expected %t", resp.AnnotationsEnabled, trueVal)
	}
}
