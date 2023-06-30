package folders

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type mockServerCall struct {
	code int
	body string
}

type mockServer struct {
	upcomingCalls []mockServerCall
	executedCalls []mockServerCall
	server        *httptest.Server
}

func GetClient(t *testing.T, code int, body string) ClientService {
	t.Helper()
	return getClientFromCalls(t, []mockServerCall{{code, body}})
}

func getClientFromCalls(t *testing.T, calls []mockServerCall) ClientService {
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(call.code)
		fmt.Fprint(w, call.body)
		mock.executedCalls = append(mock.executedCalls, call)
	}))

	u, err := url.Parse(mock.server.URL)
	if err != nil {
		t.Fatalf("failed to parse mock sever url: %v", err.Error())
	}

	c := New(httptransport.New(u.Host, "", []string{u.Scheme}), strfmt.NewFormats())

	t.Cleanup(func() {
		mock.server.Close()
	})

	return c
}
