package main

import (
	"log/slog"
	"net/http"
	"paldab/loadbalancer/config"
	"paldab/loadbalancer/lb"
	"time"
)

const (
	port = ":8080"
)

func main() {
	mux := http.NewServeMux()

	servers := config.GetServers()
	enableHealthCheck := config.GetEnableHealthEnv()
	strat := config.GetLoadBalancingStrategy(servers)
	strat.UpdateServers(servers)

	lb := lb.NewLoadBalancer(servers, strat, enableHealthCheck)

	mux.Handle("/", lb)

	time.Sleep(4 * time.Second)
	slog.Info("Listening...")
	if err := http.ListenAndServe(port, mux); err != nil {
		slog.Error("Server failed to start", "error", err.Error())
	}
}
