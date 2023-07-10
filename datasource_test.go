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

	jd := map[string]interface{}{
		"assumeRoleArn":           "arn:aws:iam::123:role/some-role",
		"authType":                "keys",
		"customMetricsNamespaces": "SomeNamespace",
		"defaultRegion":           "us-east-1",
		"tlsSkipVerify":           true,
	}
	sjd := map[string]interface{}{
		"accessKey": "123",
		"secretKey": "456",
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

	jd := map[string]interface{}{
		"httpMethod":   "POST",
		"queryTimeout": "60s",
		"timeInterval": "1m",
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

	jd := map[string]interface{}{
		"httpMethod":    "POST",
		"sigV4Auth":     true,
		"sigV4AuthType": "keys",
		"sigV4Region":   "us-east-1",
	}
	sjd := map[string]interface{}{
		"sigV4AccessKey": "123",
		"sigV4SecretKey": "456",
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

	jd := map[string]interface{}{
		"esVersion":                  "7.0.0",
		"timeField":                  "time",
		"interval":                   "1m",
		"logMessageField":            "message",
		"logLevelField":              "field",
		"maxConcurrentShardRequests": 8,
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

	jd, sjd := JSONDataWithHeaders(
		map[string]interface{}{
			"defaultBucket": "telegraf",
			"organization":  "acme",
			"version":       "Flux",
		},
		map[string]interface{}{},
		map[string]string{
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

	jd := map[string]interface{}{
		"tsdbResolution": 1,
		"tsdbVersion":    3,
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

	jd := map[string]interface{}{
		"clientId":       "lorem-ipsum",
		"cloudName":      "azuremonitor",
		"subscriptionId": "lorem-ipsum",
		"tenantId":       "lorem-ipsum",
	}
	sjd := map[string]interface{}{
		"clientSecret": "alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj==",
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
