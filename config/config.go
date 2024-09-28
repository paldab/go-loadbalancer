package config

import (
	"errors"
	"fmt"
	"os"
	"paldab/loadbalancer/models"
	"paldab/loadbalancer/strategy"
	"slices"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	defaultHealthCheck = true
)

func GetEnableHealthEnv() bool {
	envStr := os.Getenv("ENABLE_HEALTHCHECK")

	enableHealthCheck, err := strconv.ParseBool(envStr)

	if err != nil {
		return defaultHealthCheck
	}

	return enableHealthCheck
}

func GetServers() []*models.Server {
	staticServers := getStaticServersFromEnv()
	staticServersYaml := getStaticServersFromYaml()
	k8Servers := getServersFromCRD()

	return slices.Concat(staticServers, staticServersYaml, k8Servers)
}

func GetLoadBalancingStrategy(servers []*models.Server) strategy.ILoadBalancerStrategy {
	envStr := os.Getenv("STRATEGY")

	return strategy.GetStrategy(envStr, servers)
}

// Focused on local development or outside kubernetes environment
func getStaticServersFromEnv() []*models.Server {
	envStr := os.Getenv("STATIC_SERVERS")

	if envStr == "" {
		return []*models.Server{}
	}

	data := strings.Split(envStr, ",")
	servers := make([]*models.Server, 0)
	for _, url := range data {
		servers = append(servers, &models.Server{Url: url})
	}

	return servers
}

type YamlConfig struct {
	Servers []models.Server
}

// Focused on local development or outside kubernetes environment
func getStaticServersFromYaml() []*models.Server {
	yamlFilePath := "./servers.yaml"
	data, err := os.ReadFile(yamlFilePath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []*models.Server{}
		}

		fmt.Printf("Error reading config file: %v\n", err)
		return nil
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		fmt.Printf("Error unmarshalling YAML: %v\n", err)
		return nil
	}

	var servers []*models.Server
	for _, server := range yamlConfig.Servers {
		servers = append(servers, &server)
	}

	return servers
}
