package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

var headerNameRegex = regexp.MustCompile(`^httpHeaderName(\d+)$`)

// DataSource represents a Grafana data source.
type DataSource struct {
	ID     int64  `json:"id,omitempty"`
	UID    string `json:"uid,omitempty"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	URL    string `json:"url"`
	Access string `json:"access"`

	// This is only returned by the API. It can only be set through the `editable` attribute of provisioned data sources.
	ReadOnly bool `json:"readOnly"`

	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	// Deprecated: Use secureJsonData.password instead.
	Password string `json:"password,omitempty"`

	OrgID     int64 `json:"orgId,omitempty"`
	IsDefault bool  `json:"isDefault"`

	BasicAuth     bool   `json:"basicAuth"`
	BasicAuthUser string `json:"basicAuthUser,omitempty"`
	// Deprecated: Use secureJsonData.basicAuthPassword instead.
	BasicAuthPassword string `json:"basicAuthPassword,omitempty"`

	// Helper to read/write http headers
	HTTPHeaders map[string]string `json:"-"`

	JSONData       JSONData       `json:"jsonData,omitempty"`
	SecureJSONData SecureJSONData `json:"secureJsonData,omitempty"`
}

// Required to avoid recursion during (un)marshal
type _DataSource DataSource

// Marshal DataSource
func (ds *DataSource) MarshalJSON() ([]byte, error) {
	dataSource := _DataSource(*ds)
	for name, value := range ds.HTTPHeaders {
		dataSource.JSONData.httpHeaderNames = append(dataSource.JSONData.httpHeaderNames, name)
		dataSource.SecureJSONData.httpHeaderValues = append(dataSource.SecureJSONData.httpHeaderValues, value)
	}

	// Sentry provider expects this value in the JSON data payload,
	// ignoring the url attribute. This hack allows passing the URL as
	// an attribute but then sends it in the payload.
	if ds.Type == "grafana-sentry-datasource" {
		dataSource.JSONData.URL = ds.URL
	}

	return json.Marshal(dataSource)
}

// Unmarshal DataSource
func (ds *DataSource) UnmarshalJSON(b []byte) (err error) {
	dataSource := _DataSource(*ds)
	if err = json.Unmarshal(b, &dataSource); err == nil {
		*ds = DataSource(dataSource)
	}
	ds.HTTPHeaders = make(map[string]string)
	for _, value := range ds.JSONData.httpHeaderNames {
		ds.HTTPHeaders[value] = "true" // HTTP Headers are not returned by the API
	}
	return err
}

type LokiDerivedField struct {
	Name          string `json:"name"`
	MatcherRegex  string `json:"matcherRegex"`
	URL           string `json:"url"`
	DatasourceUID string `json:"datasourceUid,omitempty"`
}

// JSONData is a representation of the datasource `jsonData` property
type JSONData struct {
	// Used by all datasources
	TLSAuth                bool   `json:"tlsAuth,omitempty"`
	TLSAuthWithCACert      bool   `json:"tlsAuthWithCACert,omitempty"`
	TLSConfigurationMethod string `json:"tlsConfigurationMethod,omitempty"`
	TLSSkipVerify          bool   `json:"tlsSkipVerify,omitempty"`
	httpHeaderNames        []string

	// Used by Athena
	Catalog        string `json:"catalog,omitempty"`
	Database       string `json:"database,omitempty"`
	OutputLocation string `json:"outputLocation,omitempty"`
	Workgroup      string `json:"workgroup,omitempty"`

	// Used by Github
	GitHubURL string `json:"githubUrl,omitempty"`

	// Used by Graphite
	GraphiteVersion string `json:"graphiteVersion,omitempty"`

	// Used by Prometheus, Elasticsearch, InfluxDB, MySQL, PostgreSQL and MSSQL
	TimeInterval string `json:"timeInterval,omitempty"`

	// Used by Elasticsearch
	// From Grafana 8.x esVersion is the semantic version of Elasticsearch.
	EsVersion                  string `json:"esVersion,omitempty"`
	TimeField                  string `json:"timeField,omitempty"`
	Interval                   string `json:"interval,omitempty"`
	LogMessageField            string `json:"logMessageField,omitempty"`
	LogLevelField              string `json:"logLevelField,omitempty"`
	MaxConcurrentShardRequests int64  `json:"maxConcurrentShardRequests,omitempty"`
	XpackEnabled               bool   `json:"xpack"`

	// Used by Cloudwatch
	CustomMetricsNamespaces string `json:"customMetricsNamespaces,omitempty"`
	TracingDatasourceUID    string `json:"tracingDatasourceUid,omitempty"`

	// Used by Cloudwatch, Athena
	AuthType      string `json:"authType,omitempty"`
	AssumeRoleArn string `json:"assumeRoleArn,omitempty"`
	DefaultRegion string `json:"defaultRegion,omitempty"`
	Endpoint      string `json:"endpoint,omitempty"`
	ExternalID    string `json:"externalId,omitempty"`
	Profile       string `json:"profile,omitempty"`

	// Used by Loki
	DerivedFields []LokiDerivedField `json:"derivedFields,omitempty"`
	MaxLines      int                `json:"maxLines,omitempty"`

	// Used by OpenTSDB
	TsdbVersion    int64 `json:"tsdbVersion,omitempty"`
	TsdbResolution int64 `json:"tsdbResolution,omitempty"`

	// Used by MSSQL
	Encrypt string `json:"encrypt,omitempty"`

	// Used by PostgreSQL
	Sslmode         string `json:"sslmode,omitempty"`
	PostgresVersion int64  `json:"postgresVersion,omitempty"`
	Timescaledb     bool   `json:"timescaledb"`

	// Used by MySQL, PostgreSQL and MSSQL
	MaxOpenConns    int64 `json:"maxOpenConns,omitempty"`
	MaxIdleConns    int64 `json:"maxIdleConns,omitempty"`
	ConnMaxLifetime int64 `json:"connMaxLifetime,omitempty"`

	// Used by Prometheus
	HTTPMethod   string `json:"httpMethod,omitempty"`
	QueryTimeout string `json:"queryTimeout,omitempty"`

	// Used by Stackdriver
	AuthenticationType string `json:"authenticationType,omitempty"`
	ClientEmail        string `json:"clientEmail,omitempty"`
	DefaultProject     string `json:"defaultProject,omitempty"`
	TokenURI           string `json:"tokenUri,omitempty"`

	// Used by Prometheus and Elasticsearch
	SigV4AssumeRoleArn string `json:"sigV4AssumeRoleArn,omitempty"`
	SigV4Auth          bool   `json:"sigV4Auth"`
	SigV4AuthType      string `json:"sigV4AuthType,omitempty"`
	SigV4ExternalID    string `json:"sigV4ExternalID,omitempty"`
	SigV4Profile       string `json:"sigV4Profile,omitempty"`
	SigV4Region        string `json:"sigV4Region,omitempty"`

	// Used by Prometheus and Loki
	ManageAlerts    bool   `json:"manageAlerts"`
	AlertmanagerUID string `json:"alertmanagerUid,omitempty"`

	// Used by Alertmanager
	Implementation string `json:"implementation,omitempty"`

	// Used by Sentry
	OrgSlug string `json:"orgSlug,omitempty"`
	URL     string `json:"url,omitempty"` // Sentry is not using the datasource URL attribute

	// Used by InfluxDB
	DefaultBucket string `json:"defaultBucket,omitempty"`
	Organization  string `json:"organization,omitempty"`
	Version       string `json:"version,omitempty"`

	// Used by Azure Monitor
	AzureLogAnalyticsSameAs      bool   `json:"azureLogAnalyticsSameAs"`
	ClientID                     string `json:"clientId,omitempty"`
	CloudName                    string `json:"cloudName,omitempty"`
	LogAnalyticsClientID         string `json:"logAnalyticsClientId,omitempty"`
	LogAnalyticsDefaultWorkspace string `json:"logAnalyticsDefaultWorkspace,omitempty"`
	LogAnalyticsTenantID         string `json:"logAnalyticsTenantId,omitempty"`
	SubscriptionID               string `json:"subscriptionId,omitempty"`
	TenantID                     string `json:"tenantId,omitempty"`
}

// Required to avoid recursion during (un)marshal
type _JSONData JSONData

// Marshal JSONData
func (jd JSONData) MarshalJSON() ([]byte, error) {
	jsonData := _JSONData(jd)
	b, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	fields := make(map[string]interface{})
	if err = json.Unmarshal(b, &fields); err != nil {
		return nil, err
	}
	for index, name := range jd.httpHeaderNames {
		fields[fmt.Sprintf("httpHeaderName%d", index+1)] = name
	}
	return json.Marshal(fields)
}

// Unmarshal JSONData
func (jd *JSONData) UnmarshalJSON(b []byte) (err error) {
	jsonData := _JSONData(*jd)
	if err = json.Unmarshal(b, &jsonData); err == nil {
		*jd = JSONData(jsonData)
	}
	fields := make(map[string]interface{})
	if err = json.Unmarshal(b, &fields); err == nil {
		headerCount := 0
		for name := range fields {
			match := headerNameRegex.FindStringSubmatch(name)
			if len(match) > 0 {
				headerCount++
			}
		}

		jd.httpHeaderNames = make([]string, headerCount)
		for name, value := range fields {
			match := headerNameRegex.FindStringSubmatch(name)
			if len(match) == 2 {
				index, err := strconv.ParseInt(match[1], 10, 64)
				if err != nil {
					return err
				}
				jd.httpHeaderNames[index-1] = value.(string)
			}
		}
	}
	return err
}

// SecureJSONData is a representation of the datasource `secureJsonData` property
type SecureJSONData struct {
	// Used by all datasources
	TLSCACert         string `json:"tlsCACert,omitempty"`
	TLSClientCert     string `json:"tlsClientCert,omitempty"`
	TLSClientKey      string `json:"tlsClientKey,omitempty"`
	Password          string `json:"password,omitempty"`
	BasicAuthPassword string `json:"basicAuthPassword,omitempty"`
	httpHeaderValues  []string

	// Used by Cloudwatch, Athena
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`

	// Used by Stackdriver
	PrivateKey string `json:"privateKey,omitempty"`

	// Used by Prometheus and Elasticsearch
	SigV4AccessKey string `json:"sigV4AccessKey,omitempty"`
	SigV4SecretKey string `json:"sigV4SecretKey,omitempty"`

	// Used by GitHub
	AccessToken string `json:"accessToken,omitempty"`

	// Used by Sentry
	AuthToken string `json:"authToken,omitempty"`

	// Used by Azure Monitor
	ClientSecret string `json:"clientSecret,omitempty"`
}

// Required to avoid recursion during unmarshal
type _SecureJSONData SecureJSONData

// Marshal SecureJSONData
func (sjd SecureJSONData) MarshalJSON() ([]byte, error) {
	secureJSONData := _SecureJSONData(sjd)
	b, err := json.Marshal(secureJSONData)
	if err != nil {
		return nil, err
	}
	fields := make(map[string]interface{})
	if err = json.Unmarshal(b, &fields); err != nil {
		return nil, err
	}
	for index, value := range sjd.httpHeaderValues {
		fields[fmt.Sprintf("httpHeaderValue%d", index+1)] = value
	}
	return json.Marshal(fields)
}

// NewDataSource creates a new Grafana data source.
func (c *Client) NewDataSource(s *DataSource) (int64, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return 0, err
	}

	result := struct {
		ID int64 `json:"id"`
	}{}

	err = c.request("POST", "/api/datasources", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return 0, err
	}

	return result.ID, err
}

// UpdateDataSource updates a Grafana data source.
func (c *Client) UpdateDataSource(s *DataSource) error {
	path := fmt.Sprintf("/api/datasources/%d", s.ID)
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return c.request("PUT", path, nil, bytes.NewBuffer(data), nil)
}

// DataSource fetches and returns the Grafana data source whose ID it's passed.
func (c *Client) DataSource(id int64) (*DataSource, error) {
	path := fmt.Sprintf("/api/datasources/%d", id)
	result := &DataSource{}
	err := c.request("GET", path, nil, nil, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// DataSourceByUID fetches and returns the Grafana data source whose UID is passed.
func (c *Client) DataSourceByUID(uid string) (*DataSource, error) {
	path := fmt.Sprintf("/api/datasources/uid/%s", uid)
	result := &DataSource{}
	err := c.request("GET", path, nil, nil, result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// DataSourceIDByName returns the Grafana data source ID by name.
func (c *Client) DataSourceIDByName(name string) (int64, error) {
	path := fmt.Sprintf("/api/datasources/id/%s", name)

	result := struct {
		ID int64 `json:"id"`
	}{}

	err := c.request("GET", path, nil, nil, &result)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}

// DataSources returns all data sources as defined in Grafana.
func (c *Client) DataSources() ([]*DataSource, error) {
	result := make([]*DataSource, 0)
	err := c.request("GET", "/api/datasources", nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDataSource deletes the Grafana data source whose ID it's passed.
func (c *Client) DeleteDataSource(id int64) error {
	path := fmt.Sprintf("/api/datasources/%d", id)

	return c.request("DELETE", path, nil, nil, nil)
}
