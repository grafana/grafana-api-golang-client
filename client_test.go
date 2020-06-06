package gapi

import (
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
