package gapi

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockServerCall struct {
	code   int
	header http.Header
	body   string

	expectedReqHeader http.Header
	expectedReqBody   []byte
}

type mockServer struct {
	upcomingCalls []mockServerCall
	executedCalls []mockServerCall
	server        *httptest.Server
}

func gapiTestTools(t *testing.T, code int, body string) *Client {
	t.Helper()
	return gapiTestToolsFromCalls(t, []mockServerCall{{code: code, body: body}})
}

func gapiTestToolsFromCalls(t *testing.T, calls []mockServerCall) *Client {
	t.Helper()

	mock := &mockServer{
		upcomingCalls: calls,
	}

	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(mock.upcomingCalls) == 0 {
			t.Fatalf("unexpected call to %s %s", r.Method, r.URL)
		}
		call := mock.upcomingCalls[0]
		if len(calls) > 1 {
			mock.upcomingCalls = mock.upcomingCalls[1:]
		} else {
			mock.upcomingCalls = nil
		}

		if call.expectedReqHeader != nil {
			// The client attaches other headers that we may not be interested in. Only
			// check header names that are specified via expectedReqHeader.
			for k := range call.expectedReqHeader {
				require.Equal(t, call.expectedReqHeader.Values(k), r.Header.Values(k))
			}
		}
		if call.expectedReqBody != nil {
			reqBody, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			require.Equal(t, call.expectedReqBody, reqBody)
		}

		w.Header().Set("Content-Type", "application/json")
		for k, vals := range call.header {
			for _, v := range vals {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(call.code)

		fmt.Fprint(w, call.body)
		mock.executedCalls = append(mock.executedCalls, call)
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
		mock.server.Close()
	})

	return client
}
