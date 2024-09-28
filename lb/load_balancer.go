package lb

import (
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"paldab/loadbalancer/models"
	"paldab/loadbalancer/strategy"
	"paldab/loadbalancer/utils"
	"regexp"
	"sync"
	"time"
)

type ILoadBalancer interface {
	http.Handler
}

type LoadBalancer struct {
	Servers  []*models.Server
	Strategy strategy.ILoadBalancerStrategy

	sync.RWMutex
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	selectedServer := lb.Next(r)

	if selectedServer == nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	t := time.Now()

	parsedUrl, err := url.Parse(selectedServer.GetUrl())

	if err != nil {
		slog.Error(err.Error())
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedUrl)

	proxy.ServeHTTP(w, r)

	slog.Info("request served in", "server", selectedServer.GetUrl(), "time", time.Since(t))
}

func containsPort(input string) bool {
	portPattern := regexp.MustCompile(`:\d+$`)

	return portPattern.MatchString(input)
}

func checkServerHealth(s *models.Server) bool {
	timeoutThreshold := 2 * time.Second
	serverUrl := utils.RemoveProtocolFromUrl(s.GetUrl())

	if !containsPort(serverUrl) {
		httpPort := ":80"
		serverUrl = serverUrl + httpPort
	}

	conn, err := net.DialTimeout("tcp", serverUrl, timeoutThreshold)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

func (lb *LoadBalancer) Next(request *http.Request) *models.Server {
	defer lb.RUnlock()
	lb.RLock()

	return lb.Strategy.Next(request)
}

func (lb *LoadBalancer) healthCheck() {
	healthCheckInterval := 5 * time.Second
	ticker := time.NewTicker(healthCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Running health checks...
			lb.Lock()

			// Determine healthy services
			healthyServers := make([]*models.Server, 0)
			for _, server := range lb.Servers {
				isServerHealthy := checkServerHealth(server)
				server.IsHealthy = isServerHealthy

				if isServerHealthy {
					healthyServers = append(healthyServers, server)
				}
			}

			// Update servers
			if len(healthyServers) > 0 {
				lb.Strategy.UpdateServers(healthyServers)
				slog.Info("List of servers", "servers", models.DereferenceServers(healthyServers), "amount of servers", len(healthyServers))
			} else {
				slog.Error("There are no healthy servers available!")
			}

			lb.Unlock()
		}
	}
}

func NewLoadBalancer(servers []*models.Server, strategy strategy.ILoadBalancerStrategy, enableHealthCheck bool) ILoadBalancer {
	lb := &LoadBalancer{
		Servers:  servers,
		Strategy: strategy,
	}

	lb.Strategy.UpdateServers(servers)

	if enableHealthCheck {
		go lb.healthCheck()
	}

	return lb
}
