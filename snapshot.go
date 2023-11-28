package gapi

import (
	jsonx "github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
)

type SnapshotField struct {
	Field
	Values []interface{} `json:"values"`
}

type SnapshotData struct {
	Fields []SnapshotField        `json:"fields"`
	Meta   map[string]interface{} `json:"meta"`

	Name    string         `json:"name"`
	RefID   string         `json:"refId"`
	Unknown jsontext.Value `json:",unknown"`
}

// Snapshot represents a Grafana snapshot.
type Snapshot struct {
	DashboardModel DashboardModel `json:"dashboard"`
	Name           string         `json:"name"`
	Expires        int64          `json:"expires"`
	External       bool           `json:"external"`
}

// SnapshotResponse represents the Grafana API response to creating a dashboard.
type SnapshotCreateResponse struct {
	DeleteKey string `json:"deleteKey"`
	DeleteURL string `json:"deleteUrl"`
	Key       string `json:"key"`
	URL       string `json:"url"`
	ID        int64  `json:"id"`
}

func (c *Client) SnapshotFieldToSchemaField(snapshotField SnapshotField) SchemaField {
	return SchemaField{
		Field: Field{
			Config: snapshotField.Config,
			Labels: snapshotField.Labels,
			Name:   snapshotField.Name,
			Type:   snapshotField.Type,
		},
		TypeInfo: map[string]string{},
	}
}

// NewSnapshot creates a new Grafana snapshot.
func (c *Client) NewSnapshot(snapshot Snapshot) (*SnapshotCreateResponse, error) {
	data, err := jsonx.Marshal(snapshot, defaultJSONOptions()...)
	if err != nil {
		return nil, err
	}

	result := &SnapshotCreateResponse{}
	err = c.request("POST", "/api/snapshots", nil, data, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
