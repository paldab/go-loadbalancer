package models

import (
	"paldab/loadbalancer/utils"
)

type Server struct {
	Name           string `json:"name,omitempty"`
	Url            string `json:"url,omitempty"`
	HealthEndpoint string `json:"healthEndpoint,omitempty"`
	IsHealthy      bool   `json:"isHealthy"`
}

func (s Server) Equals(other *Server) bool {
	return s.Url == other.Url
}

func (s Server) GetUrl() string {
	if utils.HasHttpPrefix(s.Url) || utils.HasHttpsPrefix(s.Url) {
		return s.Url
	}

	return "http://" + s.Url
}

func (s Server) GetHealthUrl() string {
	defaultHealthUrl := "/health"
	if s.HealthEndpoint == "" {
		return s.Url + defaultHealthUrl
	}

	return s.Url + s.HealthEndpoint
}

func DereferenceServers(servers []*Server) []Server {
	derefArray := make([]Server, 0)

	for _, server := range servers {
		if server != nil && server.Url != "" {
			derefArray = append(derefArray, *server)
		}
	}

	return derefArray
}
