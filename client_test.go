package gapi

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
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
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_200(t *testing.T) {
	server, client := gapiTestTools(t, 200, `{"foo":"bar"}`)
	defer server.Close()

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRequest_201(t *testing.T) {
	server, client := gapiTestTools(t, 201, `{"foo":"bar"}`)
	defer server.Close()

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRequest_400(t *testing.T) {
	server, client := gapiTestTools(t, 400, `{"foo":"bar"}`)
	defer server.Close()

	expected := `status: 400, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_500(t *testing.T) {
	server, client := gapiTestTools(t, 500, `{"foo":"bar"}`)
	defer server.Close()

	expected := `status: 500, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_200Unmarshal(t *testing.T) {
	server, client := gapiTestTools(t, 200, `{"foo":"bar"}`)
	defer server.Close()

	result := struct {
		Foo string `json:"foo"`
	}{}
	err := client.request("GET", "/foo", url.Values{}, nil, &result)
	if err != nil {
		t.Fatal(err.Error())
	}

	if result.Foo != "bar" {
		t.Errorf("expected: bar; got: %s", result.Foo)
	}
}

func TestRequest_200UnmarshalPut(t *testing.T) {
	server, client := gapiTestTools(t, 200, `{"name":"mike"}`)
	defer server.Close()

	u := User{
		Name: "mike",
	}
	data, err := json.Marshal(u)
	if err != nil {
		t.Fatal(err.Error())
	}

	result := struct {
		Name string `json:"name"`
	}{}
	q := url.Values{}
	q.Add("a", "b")
	err = client.request("PUT", "/foo", q, bytes.NewBuffer(data), &result)
	if err != nil {
		t.Errorf(err.Error())
	}

	if result.Name != "mike" {
		t.Errorf("expected: name; got: %s", result.Name)
	}
}
