# Monitoring Dashboards

This document outlines the monitoring strategy and the two primary dashboards designed for the Craftista EKS environment.

## 1. Infrastructure Health Dashboard (AWS CloudWatch)
This dashboard provides a high-level view of the underlying resources powering the application.

### Key Widgets:
- **Cluster Resource Utilization (EKS)**:
  - **CPU Utilization**: Tracks average CPU usage across the `default` node group.
  - **Memory Utilization**: Monitors memory pressure to inform vertical scaling decisions.
- **Node Status**: Number of ready nodes vs. desired state.
- **Network I/O**: Tracks total bytes transmitted/received to identify potential bottlenecks.

## 2. Application & Database Performance Dashboard
This dashboard focuses on the user experience and data persistence layer.

### Key Widgets:
- **Application Performance (K8s / ALB)**:
  - **HTTP Request Rate (ALB)**: Total Number of requests per second flowing through the Ingress.
  - **Error Rate (4xx/5xx)**: Visualizes anomalous spikes in HTTP errors.
  - **Target Response Time**: Average latency for backend microservice responses (Catalogue, Voting, etc.).
- **Database Metrics (RDS)**:
  - **Database CPU & Free Storage**: Critical metrics for DB health.
  - **Database Connections**: Monitors the connection pool from the Voting service.
  - **Read/Write IOPS**: Tracks database throughput.

## 1. Metrics Collection (Prometheus)
- **Tool**: kube-prometheus-stack (Helm)
- **Namespace**: `monitoring`
- **Targets**:
  - **Infrastructure**: EKS Nodes, Kubelet, API Server (CPU, Memory, Disk).
  - **Applications**: Pod resource usage, restart counts.

## 2. Visualization (Grafana)
- **Access**: Port-forward service via `kubectl port-forward svc/prometheus-grafana 3000:80 -n monitoring`.
- **Recommended Dashboards**:
  1. **Cluster Overview**: Global view of node health and capacity.
  2. **Pod Performance**: CPU/Memory utilization per microservice.

## 3. Logging Strategy
- **Centralized Logging**: CloudWatch Logs (via Fluent Bit) or ELK Stack.
- **Log Sources**:
  - Application STDOUT/STDERR
  - System/Node logs
  - ALB Access logs
- **Log Groups**: 
  - `/aws/eks/prod-eks/cluster`: Control plane logs.
  - `/aws/containerinsights/prod-eks/application`: Application-specific logs.
