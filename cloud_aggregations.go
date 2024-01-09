package gapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type AggregationRule struct {
	Metric    string `json:"metric"`
	MatchType string `json:"match_type,omitempty"`

	Drop       bool     `json:"drop,omitempty"`
	KeepLabels []string `json:"keep_labels,omitempty"`
	DropLabels []string `json:"drop_labels,omitempty"`

	Aggregations []string `json:"aggregations,omitempty"`

	AggregationInterval string `json:"aggregation_interval,omitempty"`
	AggregationDelay    string `json:"aggregation_delay,omitempty"`

	Ingest bool `json:"ingest,omitempty"`
}

type AggregationRecommendation struct {
	AggregationRule

	RecommendedAction  string `json:"recommended_action"`
	UsagesInRules      int    `json:"usages_in_rules"`
	UsagesInQueries    int    `json:"usages_in_queries"`
	UsagesInDashboards int    `json:"usages_in_dashboards"`

	KeptLabels                   []string `json:"kept_labels,omitempty"`
	TotalSeriesAfterAggregation  int      `json:"total_series_after_aggregation,omitempty"`
	TotalSeriesBeforeAggregation int      `json:"total_series_before_aggregation,omitempty"`
}

type AggregationRecommendationConfiguration struct {
	KeepLabels []string `json:"keep_labels,omitempty"`
}

const (
	recommendationsEndpoint       = "/aggregations/recommendations"
	recommendationsConfigEndpoint = "/aggregations/recommendations/config"
	aggregationRulesEndpoint      = "/aggregations/rules"
)

func (c *Client) AggregationRecommendations(verbose bool) ([]AggregationRecommendation, error) {
	query := make(url.Values)
	query.Set("verbose", strconv.FormatBool(verbose))

	var recs []AggregationRecommendation
	err := c.request("GET", recommendationsEndpoint, query, nil, &recs)
	return recs, err
}

func (c *Client) AggregationRecommendationsConfig() (AggregationRecommendationConfiguration, error) {
	config := AggregationRecommendationConfiguration{}
	err := c.request("GET", recommendationsConfigEndpoint, nil, nil, &config)
	return config, err
}

func (c *Client) UpdateAggregationRecommendationsConfig(config AggregationRecommendationConfiguration) error {
	body, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return c.request("POST", recommendationsConfigEndpoint, nil, body, nil)
}

func (c *Client) AggregationRules() ([]AggregationRule, string, error) {
	var rules []AggregationRule
	header, err := c.requestWithHeaders("GET", aggregationRulesEndpoint, nil, nil, nil, &rules)
	if err != nil {
		return rules, "", err
	}

	etag := header.Get("ETag")
	if etag == "" {
		return rules, "", fmt.Errorf("response from %s endpoint missing etag header", aggregationRulesEndpoint)
	}

	return rules, etag, err
}

func (c *Client) UpdateAggregationRules(rules []AggregationRule, etag string) error {
	body, err := json.Marshal(rules)
	if err != nil {
		return err
	}

	header := make(http.Header)
	header.Add("If-Match", etag)

	_, err = c.requestWithHeaders("POST", aggregationRulesEndpoint, nil, header, body, nil)
	if err != nil {
		return err
	}

	return nil
}
