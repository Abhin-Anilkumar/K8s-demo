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

## Implementation Details

### Metrics Collection
- **Infrastructure**: Native CloudWatch metrics for EC2 (EKS nodes) and RDS.
- **Application**: The `voting` service exposes Prometheus metrics via Spring Boot Actuator, which are scraped for high-granularity performance data.
- **Cluster Metrics**: Provided by the `metrics-server` installed on the cluster, enabling `kubectl top` and HPA.

### Logging
- **Centralized Logging**: All containers stream logs to `stdout`/`stderr`, which are captured by AWS CloudWatch Logs. 
- **Log Groups**: 
  - `/aws/eks/prod-eks/cluster`: Control plane logs.
  - `/aws/containerinsights/prod-eks/application`: Application-specific logs.
