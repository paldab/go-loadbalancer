package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"paldab/loadbalancer/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

func isKubernetesEnvironment() bool {
	host := os.Getenv("KUBERNETES_SERVICE_HOST")
	port := os.Getenv("KUBERNETES_SERVICE_PORT")

	return host != "" || port != ""
}

type K8ServerObject struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Spec struct {
	URL       string `json:"url"`
	HealthURL string `json:"healthUrl"`
	IsHealthy bool   `json:"isHealthy"`
}

func getServersFromCRD() []*models.Server {
	// Could not fetch servers from Kubernetes cluster. You are not in a Kubernetes environment!
	if !isKubernetesEnvironment() {
		return []*models.Server{}
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error creating in-cluster config: %v\n", err)
		os.Exit(1)
	}

	// Dynamic client
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("Could not create dynamic client: %s\n", err.Error())
	}

	serverCrd := schema.GroupVersionResource{
		Group:    "loadbalancer.paldab.io",
		Version:  "v1alpha1",
		Resource: "servers",
	}

	resources, err := dynamicClient.Resource(serverCrd).Namespace("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("Could not fetch Server resource: %s\n", err.Error())
	}

	var servers []*models.Server

	for _, item := range resources.Items {
		server, err := getServerFromKubernetesResource(item)

		if err != nil {
			fmt.Printf("Error unmarshaling to Server object: %s\n", err.Error())
			continue
		}

		servers = append(servers, server)
	}

	return servers
}

func getServerFromKubernetesResource(item unstructured.Unstructured) (*models.Server, error) {
	// Convert to Json
	jsonData, err := json.Marshal(item.Object)
	if err != nil {
		return nil, err
	}

	var serverObject *K8ServerObject
	if err := json.Unmarshal(jsonData, &serverObject); err != nil {
		return nil, err
	}

	server := &models.Server{
		Name:           serverObject.Metadata.Name,
		Url:            serverObject.Spec.URL,
		HealthEndpoint: serverObject.Spec.HealthURL,
	}

	return server, nil
}
