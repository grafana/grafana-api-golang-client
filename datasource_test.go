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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		AssumeRoleArn:           "arn:aws:iam::123:role/some-role",
		AuthType:                "keys",
		CustomMetricsNamespaces: "SomeNamespace",
		DefaultRegion:           "us-east-1",
		TLSSkipVerify:           true,
	}.Map()
	if err != nil {
		t.Fatal(err)
	}
	sjd, err := SecureJSONData{
		AccessKey: "123",
		SecretKey: "456",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:           "foo",
		Type:           "cloudwatch",
		URL:            "http://some-url.com",
		Access:         "access",
		IsDefault:      true,
		JSONData:       jd,
		SecureJSONData: sjd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		HTTPMethod:   "POST",
		QueryTimeout: "60s",
		TimeInterval: "1m",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:      "foo_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData:  jd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		HTTPMethod:    "POST",
		SigV4Auth:     true,
		SigV4AuthType: "keys",
		SigV4Region:   "us-east-1",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}
	sjd, err := SecureJSONData{
		SigV4AccessKey: "123",
		SigV4SecretKey: "456",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:           "sigv4_prometheus",
		Type:           "prometheus",
		URL:            "http://some-url.com",
		Access:         "access",
		IsDefault:      true,
		JSONData:       jd,
		SecureJSONData: sjd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		EsVersion:                  "7.0.0",
		TimeField:                  "time",
		Interval:                   "1m",
		LogMessageField:            "message",
		LogLevelField:              "field",
		MaxConcurrentShardRequests: 8,
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:      "foo_elasticsearch",
		Type:      "elasticsearch",
		URL:       "http://some-url.com",
		IsDefault: true,
		JSONData:  jd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		DefaultBucket: "telegraf",
		Organization:  "acme",
		Version:       "Flux",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}
	sjd, err := SecureJSONData{}.Map()
	if err != nil {
		t.Fatal(err)
	}
	jd, sjd = JSONDataWithHeaders(jd, sjd, map[string]string{
		"Authorization": "Token alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj==",
	})

	ds := &DataSource{
		Name:           "foo_influxdb",
		Type:           "influxdb",
		URL:            "http://some-url.com",
		IsDefault:      true,
		JSONData:       jd,
		SecureJSONData: sjd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		TsdbResolution: 1,
		TsdbVersion:    3,
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:      "foo_opentsdb",
		Type:      "opentsdb",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData:  jd,
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
	client := gapiTestTools(t, 200, createdDataSourceJSON)

	jd, err := JSONData{
		ClientID:       "lorem-ipsum",
		CloudName:      "azuremonitor",
		SubscriptionID: "lorem-ipsum",
		TenantID:       "lorem-ipsum",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}
	sjd, err := SecureJSONData{
		ClientSecret: "alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj==",
	}.Map()
	if err != nil {
		t.Fatal(err)
	}

	ds := &DataSource{
		Name:           "foo_azure",
		Type:           "grafana-azure-monitor-datasource",
		URL:            "http://some-url.com",
		Access:         "access",
		IsDefault:      true,
		JSONData:       jd,
		SecureJSONData: sjd,
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
	client := gapiTestTools(t, 200, getDataSourcesJSON)

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
	client := gapiTestTools(t, 200, getDataSourceJSON)

	datasourceID, err := client.DataSourceIDByName("foo")
	if err != nil {
		t.Fatal(err)
	}

	if datasourceID != 1 {
		t.Error("Not correctly parsing returned datasources.")
	}
}

func TestDeleteDataSourceByName(t *testing.T) {
	client := gapiTestTools(t, 200, "")

	err := client.DeleteDataSourceByName("foo")
	if err != nil {
		t.Fatal(err)
	}
}
