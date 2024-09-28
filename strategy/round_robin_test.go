package strategy

import (
	"net/http"
	"net/http/httptest"
	"paldab/loadbalancer/models"
	"testing"
)

func setupServers() []*models.Server {
	return []*models.Server{
		{Name: "test1", Url: "http://localhost:3001", IsHealthy: true},
		{Name: "test3", Url: "http://localhost:3003", IsHealthy: true},
		{Name: "test2", Url: "http://localhost:3002", IsHealthy: true},
	}
}

func TestRoundRobinNext(t *testing.T) {
	servers := setupServers()
	strat := NewRoundRobinStrategy(servers)
	strat.UpdateServers(servers)

	if strat == nil {
		t.Fatalf("Expected RoundRobinStrategy to be initialized, but got nil")
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// For each server, check if round robin next is equal to server
	for _, server := range servers {
		targetServer := strat.Next(req)

		if targetServer == nil {
			t.Fatalf("Expected selected server to be non-nil, but got nil")
		}

		if targetServer.Name != server.Name {
			t.Errorf("Expected selected server to be '%s'. Got %s", server.Name, targetServer.Name)
		}
	}

	// check if servers get requeued in order
	server := strat.Next(req)
	if server.Name != "test1" {
		t.Errorf("Expected selected server to be 'test1'. Got %s", server.Name)
	}
}
