# Grafana Dashboards

This directory contains pre-configured Grafana dashboards for monitoring the Craftista EKS cluster.

## Dashboards

### 1. Infrastructure Overview (`infrastructure-overview.json`)
Monitors the health and resource utilization of the EKS cluster infrastructure.

**Panels:**
- **Cluster CPU Usage**: CPU consumption across all nodes
- **Cluster Memory Usage**: Memory utilization per node
- **Node Disk Usage**: Disk space usage percentage
- **Network I/O**: Network traffic (RX/TX) per interface
- **Pod Count by Namespace**: Total pods per namespace
- **Node Status**: Number of ready nodes

### 2. Application Performance (`application-performance.json`)
Tracks the operational metrics of microservices in `app` and `stage-app` namespaces.

**Panels:**
- **Pod CPU Usage by Service**: CPU consumption per microservice pod
- **Pod Memory Usage by Service**: Memory usage per pod
- **Pod Restart Count**: Restart frequency (indicates crashes)
- **Network Traffic by Service**: Network I/O per microservice
- **Service Availability**: Percentage of running vs desired replicas
- **Total Pods Running**: Count of active pods
- **Pod Status by Phase**: Distribution of pod states (Running, Pending, Failed)

## How to Import

1. Access Grafana: `kubectl port-forward -n monitoring svc/grafana 3000:80`
2. Navigate to **Dashboards** → **Import**
3. Upload the JSON file or paste its contents
4. Select **Prometheus** as the data source
5. Click **Import**

## Data Source Configuration

Ensure Prometheus is configured as a data source in Grafana:
- **URL**: `http://prometheus-server.monitoring.svc.cluster.local`
- **Access**: Server (default)
