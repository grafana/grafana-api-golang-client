package gapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	recommendationsResponse = `
[
  {
    "metric": "kube_job_status_failed",
    "drop_labels": [
      "instance",
      "job_name"
    ],
    "aggregations": [
      "count",
      "sum"
    ]
  },
  {
    "metric": "kube_pod_status_scheduled",
    "drop_labels": [
      "instance",
      "provider",
      "uid"
    ],
    "aggregations": [
      "min",
      "sum"
    ]
  },
  {
    "metric": "kube_pod_status_reason",
    "drop_labels": [
      "pod",
      "uid"
    ],
    "aggregations": [
      "count",
      "sum"
    ]
  },
  {
    "metric": "kube_pod_status_reason",
    "drop_labels": [
      "pod",
      "uid"
    ],
    "aggregations": [
      "count",
      "sum"
    ],
    "recommended_action": "keep",
    "usages_in_rules": 0,
    "usages_in_queries": 0,
    "usages_in_dashboards": 0
  }
]
`
	recommendationsVerboseResponse = `
[
  {
    "metric": "kube_job_status_failed",
    "drop_labels": [
      "instance",
      "job_name"
    ],
    "aggregations": [
      "count",
      "sum"
    ],
    "recommended_action": "add",
    "usages_in_rules": 0,
    "usages_in_queries": 0,
    "usages_in_dashboards": 0,
    "kept_labels": [
      "__name__",
      "asserts_env",
      "cluster",
      "job",
      "namespace",
      "provider",
      "reason"
    ],
    "total_series_after_aggregation": 99,
    "total_series_before_aggregation": 569
  },
  {
    "metric": "kube_pod_status_scheduled",
    "drop_labels": [
      "instance",
      "provider",
      "uid"
    ],
    "aggregations": [
      "min",
      "sum"
    ],
    "recommended_action": "keep",
    "usages_in_rules": 0,
    "usages_in_queries": 432,
    "usages_in_dashboards": 1
  },
  {
    "metric": "kube_pod_status_reason",
    "drop_labels": [
      "pod",
      "uid"
    ],
    "aggregations": [
      "sum"
    ],
    "recommended_action": "update",
    "usages_in_rules": 0,
    "usages_in_queries": 0,
    "usages_in_dashboards": 0
  },
  {
    "metric": "kube_job_info",
    "drop_labels": [
      "job_name"
    ],
    "aggregations": [
      "count",
      "sum"
    ],
    "recommended_action": "remove",
    "usages_in_rules": 0,
    "usages_in_queries": 2,
    "usages_in_dashboards": 0
  }
]
`
	rules = `
[
  {
    "metric": "kube_pod_status_scheduled",
    "drop_labels": [
      "instance",
      "provider",
      "uid"
    ],
    "aggregations": [
      "min",
      "sum"
    ]
  },
  {
    "metric": "kube_pod_status_reason",
    "drop_labels": [
      "pod",
      "uid"
    ],
    "aggregations": [
      "count",
      "sum"
    ]
  },
  {
    "metric": "kube_job_info",
    "drop_labels": [
      "job_name"
    ],
    "aggregations": [
      "count",
      "sum"
    ]
  }
]
`
)

func TestAggregationRecommendations(t *testing.T) {
	testCases := []struct {
		verbose bool
		resp    string
	}{
		{
			verbose: false,
			resp:    recommendationsResponse,
		},
		{
			verbose: true,
			resp:    recommendationsVerboseResponse,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("verbose=%t", tc.verbose), func(t *testing.T) {
			client := gapiTestTools(t, 200, tc.resp)

			var expected []AggregationRecommendation
			err := UnmarshalJSONToStruct(tc.resp, &expected)
			require.NoError(t, err)

			actual, err := client.AggregationRecommendations(tc.verbose)
			require.NoError(t, err)

			require.Equal(t, expected, actual)
		})
	}
}

func TestAggregationRecommendationsConfig(t *testing.T) {
	client := gapiTestTools(t, 200, `{"keep_labels":["cluster", "namespace"]}`)

	expected := AggregationRecommendationConfiguration{
		KeepLabels: []string{"cluster", "namespace"},
	}

	actual, err := client.AggregationRecommendationsConfig()
	require.NoError(t, err)

	require.Equal(t, expected, actual)
}

func TestUpdateAggregationRecommendationsConfig(t *testing.T) {
	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{
			code:            200,
			expectedReqBody: []byte(`{"keep_labels":["namespace"]}`),
		},
	})
	err := client.UpdateAggregationRecommendationsConfig(AggregationRecommendationConfiguration{
		KeepLabels: []string{"namespace"},
	})
	require.NoError(t, err)
}

func TestAggregationRules(t *testing.T) {
	t.Run("list aggregation rules", func(t *testing.T) {
		header := make(http.Header)
		header.Set("ETag", "fake-etag-value")

		client := gapiTestToolsFromCalls(t, []mockServerCall{
			{
				code:   200,
				header: header,
				body:   rules,
			},
		})

		var expected []AggregationRule
		require.NoError(t, UnmarshalJSONToStruct(rules, &expected))

		actual, etag, err := client.AggregationRules()
		require.NoError(t, err)

		require.Equal(t, expected, actual)
		require.Equal(t, "fake-etag-value", etag)
	})

	t.Run("missing etag in response", func(t *testing.T) {
		client := gapiTestTools(t, 200, rules)

		var expected []AggregationRule
		require.NoError(t, UnmarshalJSONToStruct(rules, &expected))

		_, _, err := client.AggregationRules()
		require.EqualError(t, err, "response from /aggregations/rules endpoint missing etag header")
	})
}

func TestUpdateAggregationRules(t *testing.T) {
	var body []AggregationRule
	require.NoError(t, UnmarshalJSONToStruct(recommendationsResponse, &body))

	expectedHeader := make(http.Header)
	expectedHeader.Set("If-Match", "fake-etag-value")

	// We can't just use recommendationsResponse because it's pretty-printed so we
	// marshal it back to minified json here.
	expectedBody, err := json.Marshal(body)
	require.NoError(t, err)

	client := gapiTestToolsFromCalls(t, []mockServerCall{
		{
			code:              200,
			expectedReqHeader: expectedHeader,
			expectedReqBody:   expectedBody,
		},
	})
	require.NoError(t, client.UpdateAggregationRules(body, "fake-etag-value"))
}
