package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quay/clair/v4/health"
)

func TestReadinessHandler(t *testing.T) {
	server := httptest.NewServer(health.ReadinessHandler())
	client := server.Client()

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	// default handler state should return StatusServiceUnavailable
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("failed to do request: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected %d got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}

	// signal to handler that process is ready. should return StatusOK
	health.Ready()
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to do request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, resp.StatusCode)
	}

	// signal to handler that process is unready. should return StatusServiceUnavailable
	health.UnReady()
	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("failed to do request: %v", err)
	}
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected %d got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}
}
