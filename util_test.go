package gapi

import (
	"testing"
)

func TestBuildPathAndQuery(t *testing.T) {
	params := map[string]string{
		"fizz": "buzz",
	}
	created := buildPathAndQuery("/foo/bar", params)
	expected := "/foo/bar?fizz=buzz"

	if created != expected {
		t.Error("expected buildPathAndQuery to return the correct URL path with the correct query params")
	}
}
