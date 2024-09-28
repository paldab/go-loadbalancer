package strategy

import (
	"math/rand"
	"net/http"
	"paldab/loadbalancer/models"
	"sync"
	"time"
)

type RandomStrategy struct {
	Servers []*models.Server
	sync.RWMutex
}

func (s *RandomStrategy) Next(r *http.Request) *models.Server {
	defer s.Unlock()
	s.Lock()

	rand.NewSource(time.Now().UnixNano())

	randomIdx := rand.Intn(len(s.Servers))

	return s.Servers[randomIdx]
}

func (s *RandomStrategy) UpdateServers(servers []*models.Server) {
	defer s.Unlock()
	s.Lock()

	s.Servers = servers
}

func NewRandomStrategy(servers []*models.Server) ILoadBalancerStrategy {
	return &RandomStrategy{
		Servers: servers,
	}
}
