//+build !acceptance

package gapi

import "testing"

var (
	cloudRegionResponse = `{
	"id": 1,
	"status": "active",
	"slug": "us",
	"name": "United States",
	"description": "United States",
	"createdAt": "2021-08-20T20:00:27.000Z",
	"updatedAt": "2022-01-18T20:00:51.000Z",
	"visibility": "public",
	"stackStateServiceUrl": "http://apiserver.stackstate.svc.cluster.local",
	"syntheticMonitoringApiUrl": "https://synthetic-monitoring-api.grafana.net",
	"integrationsApiUrl": "https://integrations-api-us-central.grafana.net",
	"hostedExportersApiUrl": "https://hosted-exporters-api-us-central.grafana.net",
	"machineLearningApiUrl": "https://machine-learning-prod-us-central-0.grafana.net/machine-learning",
	"incidentApiUrl": null,
	"hgClusterId": 69,
	"hgClusterSlug": "prod-us-central-0",
	"hgClusterName": "prod-us-central-0",
	"hgClusterUrl": "https://hg-api-prod-us-central-0.grafana.net",
	"hmPromClusterId": 105,
	"hmPromClusterSlug": "prod-10-prod-us-central-0",
	"hmPromClusterName": "cortex-prod-10",
	"hmPromClusterUrl": "https://prometheus-prod-10-prod-us-central-0.grafana.net",
	"hmGraphiteClusterId": 105,
	"hmGraphiteClusterSlug": "prod-10-prod-us-central-0",
	"hmGraphiteClusterName": "cortex-prod-10",
	"hmGraphiteClusterUrl": "https://prometheus-prod-10-prod-us-central-0.grafana.net",
	"hlClusterId": 84,
	"hlClusterSlug": "loki-prod-us-central-0",
	"hlClusterName": "Hosted Logs Cluster (prod-us-central-0)",
	"hlClusterUrl": "https://logs-prod3.grafana.net",
	"amClusterId": 68,
	"amClusterSlug": "alertmanager-us-central1",
	"amClusterName": "alertmanager-us-central1",
	"amClusterUrl": "https://alertmanager-us-central1.grafana.net",
	"htClusterId": 78,
	"htClusterSlug": "tempo-prod-us-central1",
	"htClusterName": "tempo-prod-us-central1",
	"htClusterUrl": "https://tempo-us-central1.grafana.net"
}
`
	cloudRegionsResponse = `{
	"items": [
		` + cloudRegionResponse + `
	]
}
`
	expectedRegion = CloudRegion{ID: 1,
		Status:                    "active",
		Slug:                      "us",
		Name:                      "United States",
		Description:               "United States",
		CreatedAt:                 "2021-08-20T20:00:27.000Z",
		UpdatedAt:                 "2022-01-18T20:00:51.000Z",
		Visibility:                "public",
		StackStateServiceURL:      "http://apiserver.stackstate.svc.cluster.local",
		SyntheticMonitoringAPIURL: "https://synthetic-monitoring-api.grafana.net",
		IntegrationsAPIURL:        "https://integrations-api-us-central.grafana.net",
		HostedExportersAPIURL:     "https://hosted-exporters-api-us-central.grafana.net",
		MachineLearningAPIURL:     "https://machine-learning-prod-us-central-0.grafana.net/machine-learning",
		IncidentsAPIURL:           "",
		HGClusterID:               69,
		HGClusterSlug:             "prod-us-central-0",
		HGClusterName:             "prod-us-central-0",
		HGClusterURL:              "https://hg-api-prod-us-central-0.grafana.net",
		HMPromClusterID:           105,
		HMPromClusterSlug:         "prod-10-prod-us-central-0",
		HMPromClusterName:         "cortex-prod-10",
		HMPromClusterURL:          "https://prometheus-prod-10-prod-us-central-0.grafana.net",
		HMGraphiteClusterID:       105,
		HMGraphiteClusterSlug:     "prod-10-prod-us-central-0",
		HMGraphiteClusterName:     "cortex-prod-10",
		HMGraphiteClusterURL:      "https://prometheus-prod-10-prod-us-central-0.grafana.net",
		HLClusterID:               84,
		HLClusterSlug:             "loki-prod-us-central-0",
		HLClusterName:             "Hosted Logs Cluster (prod-us-central-0)",
		HLClusterURL:              "https://logs-prod3.grafana.net",
		AMClusterID:               68,
		AMClusterSlug:             "alertmanager-us-central1",
		AMClusterName:             "alertmanager-us-central1",
		AMClusterURL:              "https://alertmanager-us-central1.grafana.net",
		HTClusterID:               78,
		HTClusterSlug:             "tempo-prod-us-central1",
		HTClusterName:             "tempo-prod-us-central1",
		HTClusterURL:              "https://tempo-us-central1.grafana.net"}
)

func TestCloudRegions(t *testing.T) {
	server, client := gapiTestTools(t, 200, cloudRegionsResponse)
	defer server.Close()

	regions, err := client.GetCloudRegions()

	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	// check that the number of items is the same
	if got := len(regions.Items); got != 1 {
		t.Errorf("Length of returned regions - Actual regions count: %d, Expected regions count: %d", got, 1)
	}

	if got := regions.Items[0]; got != expectedRegion {
		t.Errorf("Unexpected Region - Got:\n%#v\n, Expected:\n%#v\n", got, expectedRegion)
	}
}

func TestCloudRegionBySlug(t *testing.T) {
	server, client := gapiTestTools(t, 200, cloudRegionResponse)
	defer server.Close()

	resp, err := client.GetCloudRegionBySlug("us")
	if err != nil {
		t.Fatal(err)
	}

	if resp != expectedRegion {
		t.Errorf("Unexpected Region - Got:\n%#v\n, Expected:\n%#v\n", resp, expectedRegion)
	}
}
