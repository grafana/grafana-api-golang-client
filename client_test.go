package gapi

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
)

func TestNew_basicAuth(t *testing.T) {
	c, err := New("user:pass", "http://my-grafana.com")
	if err != nil {
		t.Errorf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://user:pass@my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}
}

func TestNew_tokenAuth(t *testing.T) {
	c, err := New("123", "http://my-grafana.com")
	if err != nil {
		t.Errorf("expected error to be nil; got: %s", err.Error())
	}

	expected := "http://my-grafana.com"
	if c.baseURL.String() != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.baseURL.String())
	}

	expected = "Bearer 123"
	if c.key != expected {
		t.Errorf("expected error: %s; got: %s", expected, c.key)
	}
}

func TestNew_invalidURL(t *testing.T) {
	_, err := New("123", "://my-grafana.com")

	expected := "parse \"://my-grafana.com\": missing protocol scheme"
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_200(t *testing.T) {
	server, client := gapiTestTools(200, `{"foo":"bar"}`)
	defer server.Close()

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRequest_201(t *testing.T) {
	server, client := gapiTestTools(201, `{"foo":"bar"}`)
	defer server.Close()

	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRequest_400(t *testing.T) {
	server, client := gapiTestTools(400, `{"foo":"bar"}`)
	defer server.Close()

	expected := `status: 400, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_500(t *testing.T) {
	server, client := gapiTestTools(500, `{"foo":"bar"}`)
	defer server.Close()

	expected := `status: 500, body: {"foo":"bar"}`
	err := client.request("GET", "/foo", url.Values{}, nil, nil)
	if err.Error() != expected {
		t.Errorf("expected error: %v; got: %s", expected, err.Error())
	}
}

func TestRequest_200Unmarshal(t *testing.T) {
	server, client := gapiTestTools(200, `{"foo":"bar"}`)
	defer server.Close()

	result := struct {
		Foo string `json:"foo"`
	}{}
	err := client.request("GET", "/foo", url.Values{}, nil, &result)
	if err != nil {
		t.Errorf(err.Error())
	}

	if result.Foo != "bar" {
		t.Errorf("expected: bar; got: %s", result.Foo)
	}
}

func TestRequest_200UnmarshalPut(t *testing.T) {
	server, client := gapiTestTools(200, `{"name":"mike"}`)
	defer server.Close()

	u := User{
		Name: "mike",
	}
	data, err := json.Marshal(u)
	if err != nil {
		t.Errorf(err.Error())
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
