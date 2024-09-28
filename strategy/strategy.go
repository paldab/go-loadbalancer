package strategy

import (
	"net/http"
	"paldab/loadbalancer/models"
)

type ILoadBalancerStrategy interface {
	Next(request *http.Request) *models.Server
	UpdateServers(servers []*models.Server)
}
