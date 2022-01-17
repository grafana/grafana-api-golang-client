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

// JSONData is a representation of the datasource `jsonData` property
type JSONData struct {
	// Used by all datasources
	TLSAuth           bool `json:"tlsAuth,omitempty"`
	TLSAuthWithCACert bool `json:"tlsAuthWithCACert,omitempty"`
	TLSSkipVerify     bool `json:"tlsSkipVerify,omitempty"`
	httpHeaderNames   []string

	// Used by Github
	GitHubUrl string `json:"githubUrl,omitempty"`

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

	// Used by Cloudwatch
	AuthType                string `json:"authType,omitempty"`
	AssumeRoleArn           string `json:"assumeRoleArn,omitempty"`
	DefaultRegion           string `json:"defaultRegion,omitempty"`
	CustomMetricsNamespaces string `json:"customMetricsNamespaces,omitempty"`
	Profile                 string `json:"profile,omitempty"`

	// Used by OpenTSDB
	TsdbVersion    string `json:"tsdbVersion,omitempty"`
	TsdbResolution string `json:"tsdbResolution,omitempty"`

	// Used by MSSQL
	Encrypt string `json:"encrypt,omitempty"`

	// Used by PostgreSQL
	Sslmode         string `json:"sslmode,omitempty"`
	PostgresVersion int64  `json:"postgresVersion,omitempty"`
	Timescaledb     bool   `json:"timescaledb,omitempty"`

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
	SigV4Auth          bool   `json:"sigV4Auth,omitempty"`
	SigV4AuthType      string `json:"sigV4AuthType,omitempty"`
	SigV4ExternalID    string `json:"sigV4ExternalID,omitempty"`
	SigV4Profile       string `json:"sigV4Profile,omitempty"`
	SigV4Region        string `json:"sigV4Region,omitempty"`

	// Used by Prometheus and Loki
	ManageAlerts    bool   `json:"manageAlerts,omitempty"`
	AlertmanagerUID string `json:"alertmanagerUid,omitempty"`

	// Used by Alertmanager
	Implementation string `json:"implementation,omitempty"`
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

	// Used by Cloudwatch
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`

	// Used by Stackdriver
	PrivateKey string `json:"privateKey,omitempty"`

	// Used by Prometheus and Elasticsearch
	SigV4AccessKey string `json:"sigV4AccessKey,omitempty"`
	SigV4SecretKey string `json:"sigV4SecretKey,omitempty"`
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

// DeleteDataSource deletes the Grafana data source whose ID it's passed.
func (c *Client) DeleteDataSource(id int64) error {
	path := fmt.Sprintf("/api/datasources/%d", id)

	return c.request("DELETE", path, nil, nil, nil)
}
