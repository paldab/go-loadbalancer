# Load Balancer with CRD-Based Server Tracking

This project implements a custom Kubernetes load balancer that dynamically retrieves a list of servers from a Custom Resource Definition (CRD) called `Server`. The load balancer uses this information to distribute traffic using a round-robin strategy, perform health checks, and update the status of each server in the Kubernetes cluster.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Custom Resource Definition (CRD)](#custom-resource-definition-crd)
- [Health Checking](#health-checking)
- [Dynamic Port Detection](#dynamic-port-detection)
- [Usage](#usage)

## Overview

This project demonstrates how to create a custom Kubernetes load balancer that interacts with a CRD-based resource (`Server`). The load balancer retrieves the list of available servers, performs health checks, and routes traffic using a round-robin load-balancing algorithm. The project is designed to integrate seamlessly with Kubernetes, making it easy to track and manage servers across namespaces.

## Features

- **Custom Resource Definition (CRD) for Servers**: Defines a custom resource in Kubernetes to track servers (`Server`), including their URL and health status.
- **Round Robin Load Balancing**: The load balancer distributes traffic evenly across all healthy servers.
- **Health Check System**: The system performs health checks on each server's `healthUrl` and updates the server's health status in the CRD.
- **Dynamic Port Detection**: The load balancer dynamically detects ports from Kubernetes DNS when the server URL does not specify one.
- **Kubernetes Native**: The load balancer is designed to work natively with Kubernetes, leveraging Kubernetes resources and APIs.

## Architecture

1. **Load Balancer**: 
   - Developed in Go, the load balancer interacts with the Kubernetes API to list and track `Server` resources.
   - It performs health checks and updates the status (`isHealthy`) of each server in the Kubernetes CRD.
   - Traffic is distributed using a round-robin load balancing strategy.

2. **Custom Resource Definition (CRD)**:
   - The `Server` CRD defines the structure of a server object with fields like `url`, `healthUrl`, and `isHealthy`.
   - Kubernetes administrators can create and manage `Server` resources, which are monitored by the load balancer.

## Custom Resource Definition (CRD)

The `Server` CRD allows you to define server instances in Kubernetes. Each instance includes the following key fields:

- **url**: The URL of the server.
- **healthUrl**: The endpoint to perform health checks on.
- **isHealthy**: A boolean that is updated based on the result of the health check.

### Example `Server` Object

```yaml
apiVersion: loadbalancer.paldab.io/v1alpha1
kind: Server
metadata:
  name: example-server
spec:
  url: http://example-server.default.svc.cluster.local
  healthUrl: http://example-server.default.svc.cluster.local/health
  isHealthy: false
