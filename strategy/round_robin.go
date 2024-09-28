package strategy

import (
	"net/http"
	"paldab/loadbalancer/models"
	"paldab/loadbalancer/queue"
	"sync"
)

type RoundRobinStrategy struct {
	Servers *queue.Queue[models.Server]
	sync.RWMutex
}

func (s *RoundRobinStrategy) SelectServer(servers *queue.Queue[models.Server]) *models.Server {
	if servers.Len() == 0 {
		return nil
	}

	selectedServer, ok := servers.Dequeue()
	if !ok {
		return nil
	}

	if selectedServer.IsHealthy {
		defer servers.Enqueue(selectedServer)
		return &selectedServer
	}

	return nil
}

func (s *RoundRobinStrategy) Next(request *http.Request) *models.Server {
	defer s.Unlock()
	s.Lock()

	return s.SelectServer(s.Servers)
}

func (s *RoundRobinStrategy) UpdateServers(servers []*models.Server) {
	defer s.Unlock()
	s.Lock()

	s.Servers.Clear()

	for _, server := range servers {
		if server.IsHealthy {
			s.Servers.Enqueue(*server)
		}
	}
}

func NewRoundRobinStrategy(servers []*models.Server) ILoadBalancerStrategy {
	q := queue.NewQueue[models.Server]()

	for _, server := range servers {
		q.Enqueue(*server)
	}

	return &RoundRobinStrategy{
		Servers: &q,
	}
}
