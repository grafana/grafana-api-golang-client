package gapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNew_basicAuth(t *testing.T) {
	c, err := New("http://my-grafana.com", Config{BasicAuth: url.UserPassword("user", "pass")})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://user:pass@my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}
}

func TestNew_tokenAuth(t *testing.T) {
	const apiKey = "123"
	c, err := New("http://my-grafana.com", Config{APIKey: apiKey})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}

	if c.config.APIKey != apiKey {
		t.Errorf("expected error: %s; got: %s", apiKey, c.config.APIKey)
	}
}

func TestNew_orgID(t *testing.T) {
	const orgID = 456
	c, err := New("http://my-grafana.com", Config{OrgID: orgID})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}

	if c.config.OrgID != orgID {
		t.Errorf("expected error: %d; got: %d", orgID, c.config.OrgID)
	}
}

func TestNew_HTTPHeaders(t *testing.T) {
	const key = "foo"
	headers := map[string]string{key: "bar"}
	c, err := New("http://my-grafana.com", Config{HTTPHeaders: headers})
	if err != nil {
		t.Fatalf("expected error to be nil; got: %s", err.Error())
	}

	value, ok := c.config.HTTPHeaders[key]
	if !ok {
		t.Errorf("expected error: %v; got: %v", headers, c.config.HTTPHeaders)
	}
	if value != headers[key] {
		t.Errorf("expected error: %s; got: %s", headers[key], value)
	}
}

func TestNew_invalidURL(t *testing.T) {
	_, err := New("://my-grafana.com", Config{APIKey: "123"})

	expected := "parse \"://my-grafana.com\": missing protocol scheme"
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_200(t *testing.T) {
	client := gapiTestTools(t, 200, `{"foo":"bar"}`)

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRequest_201(t *testing.T) {
	client := gapiTestTools(t, 201, `{"foo":"bar"}`)

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestRequest_400(t *testing.T) {
	client := gapiTestTools(t, 400, `{"foo":"bar"}`)

	expected := `status: 400, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_500(t *testing.T) {
	client := gapiTestTools(t, 500, `{"foo":"bar"}`)

	expected := `status: 500, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_badURL(t *testing.T) {
	client := gapiTestTools(t, 200, `{"foo":"bar"}`)
	baseURL, err := url.Parse("bad-url")
	if err != nil {
		t.Fatal(err)
	}
	client.baseURL = *baseURL

	expected := `Get "bad-url/foo": unsupported protocol scheme ""`
	err = client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err)
	}
}

func TestRequest_200Unmarshal(t *testing.T) {
	client := gapiTestTools(t, 200, `{"foo":"bar"}`)

	result := struct {
		Foo string `json:"foo"`
	}{}
	err := client.request("GET", "/foo", url.Values{}, nil, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Foo != "bar" {
		t.Errorf("expected: bar; got: %s", result.Foo)
	}
}

func TestRequest_200UnmarshalPut(t *testing.T) {
	client := gapiTestTools(t, 200, `{"name":"mike"}`)

	u := User{
		Name: "mike",
	}
	data, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err)
	}

	result := struct {
		Name string `json:"name"`
	}{}
	q := url.Values{}
	q.Add("a", "b")
	err = client.request("PUT", "/foo", q, data, &result)
	if err != nil {
		t.Error(err)
	}

	if result.Name != "mike" {
		t.Errorf("expected: name; got: %s", result.Name)
	}
}

func TestClient_requestWithRetries(t *testing.T) {
	// Test that calls to c.client.Do will retry correctly,
	// even if the original request fails prematurely

	body := []byte(`lorem ipsum dolor sit amet`)

	var try int

	// This is our actual test, checking that we do in fact receive a body.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		try++

		got, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("retry %d: unexpected error reading body: %v", try, err)
		}

		if !bytes.Equal(body, got) {
			t.Errorf("retry %d: request body doesn't match body sent by client:\nexp: %v\ngot: %v", try, body, got)
		}

		switch try {
		case 1:
			http.Error(w, `{"error":"waiting for the right time"}`, http.StatusInternalServerError)

		case 2:
			http.Error(w, `{"error":"calm down"}`, http.StatusTooManyRequests)

		case 3:
			w.Write([]byte(`{"foo":"bar"}`)) //nolint:errcheck

		default:
			t.Errorf("unexpected retry %d", try)
		}
	}))
	defer ts.Close()

	// From http.Client.Do documentation: an error is returned if
	// caused by client policy (such as CheckRedirect), or failure to
	// speak HTTP (such as a network connectivity problem). A non-2xx
	// status code doesn't cause an error.
	//
	// For this reason we build a custom http.Client that will fail
	// the first time with an error *before* the request is sent, and
	// succeed afterwards. See customRoundTripper below.
	httpClient := &http.Client{
		Transport: &customRoundTripper{},
	}

	c, err := New(ts.URL, Config{
		NumRetries:   5,
		Client:       httpClient,
		RetryTimeout: 50 * time.Millisecond,
	})
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	type res struct {
		Foo string `json:"foo"`
	}

	var got res

	if err := c.request(http.MethodPost, "/", nil, body, &got); err != nil {
		t.Fatalf("unexpected error sending request: %v", err)
	}

	exp := res{Foo: "bar"}

	if exp != got {
		t.Fatalf("response doesn't match\nexp: %#v\ngot: %#v", exp, got)
	}

	t.Logf("request successful after %d retries", try)
}

func TestClient_CustomRetryStatusCode(t *testing.T) {
	body := []byte(`lorem ipsum dolor sit amet`)
	var try int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		try++

		switch try {
		case 1, 2:
			http.Error(w, `{"error":"weird error"}`, http.StatusUpgradeRequired)
		default:
			http.Error(w, `{"error":"failed"}`, http.StatusInternalServerError)
		}
	}))
	defer ts.Close()

	httpClient := &http.Client{
		Transport: &customRoundTripper{},
	}

	c, err := New(ts.URL, Config{
		NumRetries:       5,
		Client:           httpClient,
		RetryTimeout:     50 * time.Millisecond,
		RetryStatusCodes: []string{strconv.Itoa(http.StatusUpgradeRequired)},
	})
	if err != nil {
		t.Fatalf("unexpected error creating client: %v", err)
	}

	var got interface{}
	err = c.request(http.MethodPost, "/", nil, body, &got)
	expectedErr := "status: 500, body: {\"error\":\"failed\"}" // The 500 is not retried because it's not in RetryStatusCodes
	if strings.TrimSpace(err.Error()) != expectedErr {
		t.Fatalf("expected err: %s, got err: %v", expectedErr, err)
	}

	if try != 3 {
		t.Fatalf("unexpected number of tries: %d", try)
	}
}

type customRoundTripper struct {
	try int
}

func (rt *customRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if rt.try++; rt.try < 2 {
		return nil, errors.New("failure")
	}

	return http.DefaultTransport.RoundTrip(r)
}
