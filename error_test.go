package gapi

import (
	"testing"
)

const (
	someErrorJSON          = `{"message":"some unknown error occurred"}`
	dataSourceNotFoundJSON = `{"message":"datasource not found"}`
)

func TestGAPIError(t *testing.T) {
	server, client := gapiTestTools(t, 500, someErrorJSON)
	defer server.Close()

	_, err := client.DataSource(100)
	gapiErr := GetGApiError(err)
	if gapiErr == nil {
		t.Fatal("Error should be a GApiError")
	}
	if gapiErr.StatusCode() != 500 {
		t.Error("Error status code should be 500")
	}

	if gapiErr.Message() != someErrorJSON {
		t.Errorf("Error message should be %s", someErrorJSON)
	}
}

func TestIsNotFoundError(t *testing.T) {
	server, client := gapiTestTools(t, 404, dataSourceNotFoundJSON)
	defer server.Close()

	_, err := client.DataSource(100)
	if !IsNotFound(err) {
		t.Fatal("Error should be a GApiError")
	}
}
