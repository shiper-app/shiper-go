package shiper

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(t *testing.T, status int, contentType, body string) *Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(srv.Close)
	t.Setenv(envBaseURL, srv.URL)

	c, err := NewClient(WithToken("t"))
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestCall(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		c := newTestClient(t, http.StatusOK, "application/json", `[{"id":"e1","branch":"main","enabled":true,"production":true,"projectId":"p1"}]`)
		resp, err := c.ListEnvironments(ctx, "p1")
		if err != nil {
			t.Fatal(err)
		}
		if got := (*resp.JSON200)[0].Id; got != "e1" {
			t.Fatalf("id = %q", got)
		}
	})

	t.Run("non-2xx", func(t *testing.T) {
		c := newTestClient(t, http.StatusNotFound, "application/json", `{"message":"nope"}`)
		_, err := c.GetProject(ctx, "p1")
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound || apiErr.Operation != "GetProject" {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("missing JSON body", func(t *testing.T) {
		c := newTestClient(t, http.StatusOK, "text/plain", "ok")
		var apiErr *APIError
		if _, err := c.GetMe(ctx); !errors.As(err, &apiErr) {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("required ID", func(t *testing.T) {
		c := newTestClient(t, http.StatusOK, "application/json", `{}`)
		if _, err := c.DeleteDomain(ctx, "p1", "e1", " "); err == nil || err.Error() != "domain is required" {
			t.Fatalf("err = %v", err)
		}
	})

	t.Run("nil client", func(t *testing.T) {
		var c *Client
		if _, err := c.GetMe(ctx); !errors.Is(err, errClientNotInitialized) {
			t.Fatalf("err = %v", err)
		}
	})
}
