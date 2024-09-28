package strategy

import "paldab/loadbalancer/models"

func GetStrategy(strategy string, servers []*models.Server) ILoadBalancerStrategy {
	switch strategy {
	case "ROUND_ROBIN":
		return NewRoundRobinStrategy(servers)
	case "RANDOM":
		return NewRandomStrategy(servers)
	default:
		return NewRoundRobinStrategy(servers)
	}
}
