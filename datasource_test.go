//go:build !integration
// +build !integration

package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	createdDataSourceJSON = `{"id":1,"uid":"myuid0001","message":"Datasource added", "name": "test_datasource"}`
	getDataSourceJSON     = `{"id":1}`
	getDataSourcesJSON    = `[{"id":1,"name":"foo","type":"cloudwatch","url":"http://some-url.com","access":"access","isDefault":true}]`
)

func TestNewDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo",
		Type:      "cloudwatch",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			AssumeRoleArn:           "arn:aws:iam::123:role/some-role",
			AuthType:                "keys",
			CustomMetricsNamespaces: "SomeNamespace",
			DefaultRegion:           "us-east-1",
			TLSSkipVerify:           true,
		},
		SecureJSONData: SecureJSONData{
			AccessKey: "123",
			SecretKey: "456",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			HTTPMethod:   "POST",
			QueryTimeout: "60s",
			TimeInterval: "1m",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusSigV4DataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "sigv4_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			HTTPMethod:    "POST",
			SigV4Auth:     true,
			SigV4AuthType: "keys",
			SigV4Region:   "us-east-1",
		},
		SecureJSONData: SecureJSONData{
			SigV4AccessKey: "123",
			SigV4SecretKey: "456",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewElasticsearchDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_elasticsearch",
		Type:      "elasticsearch",
		URL:       "http://some-url.com",
		IsDefault: true,
		JSONData: JSONData{
			EsVersion:                  "7.0.0",
			TimeField:                  "time",
			Interval:                   "1m",
			LogMessageField:            "message",
			LogLevelField:              "field",
			MaxConcurrentShardRequests: 8,
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewInfluxDBDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_influxdb",
		Type:      "influxdb",
		URL:       "http://some-url.com",
		IsDefault: true,
		JSONData: JSONData{
			DefaultBucket:   "telegraf",
			httpHeaderNames: []string{"Authorization"},
			Organization:    "acme",
			Version:         "Flux",
		},
		SecureJSONData: SecureJSONData{
			httpHeaderValues: []string{"Token alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj=="},
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewOpenTSDBDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_opentsdb",
		Type:      "opentsdb",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			TsdbResolution: 1,
			TsdbVersion:    3,
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewAzureDataSource(t *testing.T) {
	server, client := gapiTestTools(t, 200, createdDataSourceJSON)
	defer server.Close()

	ds := &DataSource{
		Name:      "foo_azure",
		Type:      "grafana-azure-monitor-datasource",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: JSONData{
			ClientID:       "lorem-ipsum",
			CloudName:      "azuremonitor",
			SubscriptionID: "lorem-ipsum",
			TenantID:       "lorem-ipsum",
		},
		SecureJSONData: SecureJSONData{
			ClientSecret: "alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj==",
		},
	}

	created, err := client.NewDataSource(ds)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(created))

	if created != 1 {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestDataSources(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSourcesJSON)
	defer server.Close()

	datasources, err := client.DataSources()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(datasources))

	if len(datasources) != 1 {
		t.Error("Length of returned datasources should be 1")
	}
	if datasources[0].ID != 1 || datasources[0].Name != "foo" {
		t.Error("Not correctly parsing returned datasources.")
	}
}

func TestDataSourceIDByName(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSourceJSON)
	defer server.Close()

	datasourceID, err := client.DataSourceIDByName("foo")
	if err != nil {
		t.Fatal(err)
	}

	if datasourceID != 1 {
		t.Error("Not correctly parsing returned datasources.")
	}
}
