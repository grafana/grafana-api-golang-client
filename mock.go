package gapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

type mockServerCall struct {
	code      int
	body      string
	reqURI    string
	reqMethod string
	reqBody   string
}

type mockServer struct {
	upcomingCalls []mockServerCall
	executedCalls []mockServerCall
	server        *httptest.Server
}

func (m *mockServer) Close() {
	m.server.Close()
}

func gapiTestTools(t *testing.T, code int, body string) *Client {
	return gapiTestToolsFromCalls(t, []mockServerCall{{code: code, body: body}})
}

func gapiTestToolsFromCalls(t *testing.T, calls []mockServerCall) *Client {
	t.Helper()

	mock := &mockServer{
		upcomingCalls: calls,
	}

	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		call := mock.upcomingCalls[0]
		if len(calls) > 1 {
			mock.upcomingCalls = mock.upcomingCalls[1:]
		} else {
			mock.upcomingCalls = nil
		}
		w.WriteHeader(call.code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, call.body)
		mock.executedCalls = append(mock.executedCalls, call)
		call.testRequestData(t, r)
	}))

	tr := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(mock.server.URL)
		},
	}

	httpClient := &http.Client{Transport: tr}

	client, err := New("http://my-grafana.com", Config{APIKey: "my-key", Client: httpClient})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		mock.Close()
	})

	return client
}

func (c mockServerCall) testRequestData(t *testing.T, req *http.Request) {
	t.Helper()

	if c.reqURI != "" && c.reqURI != req.URL.RequestURI() {
		t.Errorf("got wrong request URI, expected: %s, got: %s", c.reqURI, req.URL.RequestURI())
	}

	if c.reqMethod != "" && c.reqMethod != req.Method {
		t.Errorf("got wrong request method, expected: %s, got: %s", c.reqMethod, req.Method)
	}

	if c.reqBody == "" {
		return
	} else if req.Body == nil {
		t.Errorf("got wrong request body, expected: %s, got: <nil>", c.reqBody)
		return
	}

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("got unexpected error reading request body: %v", err)
	}

	if strings.TrimSpace(c.reqBody) != strings.TrimSpace(string(reqBody)) {
		t.Errorf("got wrong request body, expected: %s, got: %s", c.reqBody, string(reqBody))
	}
}
