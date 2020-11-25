package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

func TestNewDataSourceGeneric(t *testing.T) {
	server, client := gapiTestTools(200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSourceGeneric{
		Name:      "foo",
		Type:      "cloudwatch",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JsonData{
			"AssumeRoleArn":           "arn:aws:iam::123:role/some-role",
			"AuthType":                "keys",
			"CustomMetricsNamespaces": "SomeNamespace",
			"DefaultRegion":           "us-east-1",
			"TlsSkipVerify":           true,
		},
		SecureJSONData: JsonData{
			"AccessKey": "123",
			"SecretKey": "456",
		},
	}

	created, err := client.NewDataSourceGeneric(ds)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusDataSourceGeneric(t *testing.T) {
	server, client := gapiTestTools(200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSourceGeneric{
		Name:      "foo_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JsonData{
			"HttpMethod":   "POST",
			"QueryTimeout": "60s",
			"TimeInterval": "1m",
		},
	}

	created, err := client.NewDataSourceGeneric(ds)
	if err != nil {
		t.Error(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("data source creation response should return the created data source ID")
	}
}