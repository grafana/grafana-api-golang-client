package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	getDatasourceCacheJSON = `
	{
		"message": "Data source cache settings loaded",
		"dataSourceID": 1,
		"dataSourceUID": "jZrmlLCGka",
		"enabled": true,
		"useDefaultTTL": false,
		"ttlQueriesMs": 60000,
		"ttlResourcesMs": 300000,
		"defaultTTLMs": 300000,
		"created": "2023-04-21T11:49:22-04:00",
		"updated": "2023-04-24T17:03:40-04:00"
	}`
	updateDatasourceCacheJSON = `
	{
		"message": "Data source cache settings updated",
		"dataSourceID": 1,
		"dataSourceUID": "jZrmlLCGka",
		"enabled": true,
		"useDefaultTTL": false,
		"ttlQueriesMs": 60000,
		"ttlResourcesMs": 300000,
		"defaultTTLMs": 300000,
		"created": "2023-04-21T11:49:22-04:00",
		"updated": "2023-04-24T17:03:40-04:00"
 	}`
)

func TestDatasourceCache(t *testing.T) {
	client := gapiTestTools(t, 200, getDatasourceCacheJSON)
	resp, err := client.DatasourceCache(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	expects := DatasourceCache{
		Message:        "Data source cache settings loaded",
		DatasourceID:   1,
		DatasourceUID:  "jZrmlLCGka",
		Enabled:        true,
		UseDefaultTLS:  false,
		TTLQueriesMs:   60000,
		TTLResourcesMs: 300000,
		DefaultTTLMs:   300000,
		Created:        "2023-04-21T11:49:22-04:00",
		Updated:        "2023-04-24T17:03:40-04:00",
	}

	if resp.Enabled != expects.Enabled ||
		resp.DatasourceUID != expects.DatasourceUID ||
		resp.UseDefaultTLS != expects.UseDefaultTLS ||
		resp.TTLQueriesMs != expects.TTLQueriesMs ||
		resp.TTLResourcesMs != expects.TTLResourcesMs ||
		resp.DefaultTTLMs != expects.DefaultTTLMs {
		t.Error("Not correctly parsing returned datasource cache")
	}
}

func TestUpdateDatasourceCache(t *testing.T) {
	client := gapiTestTools(t, 200, updateDatasourceCacheJSON)
	payload := &DatasourceCachePayload{
		DatasourceID:   1,
		DatasourceUID:  "jZrmlLCGka",
		Enabled:        true,
		UseDefaultTLS:  true,
		TTLQueriesMs:   6000,
		TTLResourcesMs: 30000,
	}
	err := client.UpdateDatasourceCache(1, payload)
	if err != nil {
		t.Error(err)
	}
}
