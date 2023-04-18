package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var sloPath string = "/api/plugins/grafana-slo-app/resources/v1/slo"

type Slos struct {
	Slos []Slo `json:"slos"`
}

type Slo struct {
	Uuid                  string        `json:"uuid"`
	Name                  string        `json:"name"`
	Description           string        `json:"description"`
	Service               string        `json:"service,omitempty"`
	Query                 Query         `json:"query"`
	Alerting              *Alerting     `json:"alerting,omitempty"`
	Labels                *[]Label      `json:"labels,omitempty"`
	Objectives            []Objective   `json:"objectives"`
	DrilldownDashboardUid string        `json:"dashboardUid,omitempty"`
	DrilldownDashboardRef *DashboardRef `json:"drillDownDashboardRef,omitempty"`
}

type Alerting struct {
	Name        string         `json:"name"`
	Annotations *[]Label       `json:"annotations,omitempty"`
	Labels      *[]Label       `json:"labels,omitempty"`
	FastBurn    *AlertMetadata `json:"fastBurn,omitempty"`
	SlowBurn    *AlertMetadata `json:"slowBurn,omitempty"`
}

type AlertMetadata struct {
	Annotations *[]Label `json:"annotations,omitempty"`
	Labels      *[]Label `json:"labels,omitempty"`
}

type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Objective struct {
	Value  float64 `json:"value"`
	Window string  `json:"window"`
}

type DashboardRef struct {
	ID  int    `json:"id,omitempty"`
	UID string `json:"uid,omitempty"`
}

type FreeformQuery struct {
	Query string `json:"freeformQuery,omitempty"`
}

type ThresholdQuery struct {
	ThresholdMetric *MetricDef `json:"thresholdMetric,omitempty"`
}

type RatioQuery struct {
	SuccessMetric *MetricDef `json:"successMetric,omitempty"`
	TotalMetric   *MetricDef `json:"totalMetric,omitempty"`
}

type PercentileQuery struct {
	HistogramMetric *MetricDef `json:"histogramMetric,omitempty"`
	Percentile      float64    `json:"percentile,omitempty"`
}

type Threshold struct {
	Value    float64 `json:"value,omitempty"`
	Operator string  `json:"operator,omitempty"`
}

type MetricDef struct {
	PrometheusMetric string `json:"prometheusMetric,omitempty"`
	Type             string `json:"type,omitempty"`
}

type Query struct {
	ThresholdQuery
	RatioQuery
	PercentileQuery
	FreeformQuery
	Threshold     *Threshold `json:"threshold,omitempty"`
	GroupByLabels []string   `json:"groupBy,omitempty"`
}

type CreateSLOResponse struct {
	Message string `json:"message,omitempty"`
	Uuid    string `json:"uuid,omitempty"`
}

// ListSLOs retrieves a list of all SLOs
func (c *Client) ListSLOs() (Slos, error) {
	var slos Slos

	if err := c.request("GET", sloPath, nil, nil, &slos); err != nil {
		return Slos{}, err
	}

	return slos, nil
}

// GetSLO returns a single SLO based on its uuid
func (c *Client) GetSlo(uuid string) (Slo, error) {
	var slo Slo
	path := fmt.Sprintf("%s/%s", sloPath, uuid)

	if err := c.request("GET", path, nil, nil, &slo); err != nil {
		return Slo{}, err
	}

	return slo, nil
}

// CreateSLO creates a single SLO
func (c *Client) CreateSLO(slo Slo) (CreateSLOResponse, error) {
	response := CreateSLOResponse{}

	data, err := json.Marshal(slo)
	if err != nil {
		return response, err
	}

	if err := c.request("POST", sloPath, nil, bytes.NewBuffer(data), &response); err != nil {
		return CreateSLOResponse{}, err
	}

	return response, err
}

// DeleteSLO deletes the SLO with the passed in UUID
func (c *Client) DeleteSLO(uuid string) error {
	var slo Slo
	path := fmt.Sprintf("%s/%s", sloPath, uuid)

	if err := c.request("DELETE", path, nil, nil, &slo); err != nil {
		return err
	}

	return nil
}
