# Monitoring & Observability Strategy

This document outlines the implemented monitoring and logging stack for the Craftista application.

## 1. Infrastructure Health Monitoring (Prometheus + Grafana)

### Prometheus Stack Components
- **Prometheus Server**: Collects and stores time-series metrics
- **Node Exporter**: Exposes hardware and OS metrics from each EKS node
- **Kube State Metrics**: Exposes Kubernetes cluster state metrics
- **Alertmanager**: Handles alerts from Prometheus

### Monitored Metrics
- **Infrastructure**: 
  - EKS Node CPU, Memory, Disk, Network I/O
  - Pod resource utilization and restart counts
  - Kubernetes API server performance
- **Applications**: 
  - HTTP request rates and error rates (via service monitors)
  - Pod health and availability
  - Container resource consumption

### Grafana Dashboards
- **Access**: `kubectl port-forward -n monitoring svc/grafana 3000:80`
- **Data Source**: Prometheus at `http://prometheus-server.monitoring.svc.cluster.local`
- **Pre-configured Dashboards**: 
  - **Infrastructure Overview** (`grafana-dashboards/infrastructure-overview.json`): Node CPU/Memory/Disk, Network I/O
  - **Application Performance** (`grafana-dashboards/application-performance.json`): Pod metrics, restarts, availability
- **Import Instructions**: Dashboards → Import → Upload JSON file → Select Prometheus data source

## 2. Database Monitoring (CloudWatch)

### RDS Metrics (via CloudWatch)
- **Database CPU & Memory**: Tracks resource utilization
- **Free Storage Space**: Monitors disk capacity
- **Database Connections**: Active connection count from `voting` service
- **Read/Write IOPS**: Database throughput metrics

## 3. Centralized Logging (CloudWatch Logs)

### Log Groups
- `/aws/eks/prod-eks/cluster`: EKS control plane logs (API server, scheduler, controller manager)

### Log Sources
- EKS control plane audit logs
- Kubernetes system component logs

## 4. Storage Infrastructure (EBS CSI Driver)

### Implementation
- **Driver**: AWS EBS CSI Driver (EKS addon)
- **IAM Role**: IRSA-enabled for `ebs-csi-controller-sa`
- **StorageClass**: `gp2` (default, WaitForFirstConsumer binding mode)
- **Persistent Volumes**: Auto-provisioned for Prometheus and Alertmanager data retention

## 5. Application Access

### Production URL
- **Domain**: [evoqu.in](http://evoqu.in)
- **Exposure**: Kubernetes LoadBalancer service (frontend)
- **DNS Provider**: Hostinger
- **Note**: ALB with Ingress not used due to AWS account restrictions on load balancer creation

### Cluster Scaling
- **Current Nodes**: 5x t3.medium instances
- **Scaling Config**: Min 3, Max 10, Desired 5
- **Auto-Scaling**: Enabled via Cluster Autoscaler

---

## 6. Proof of Monitoring Implementation

### Grafana Access

![Grafana Login](screenshot/grafana/Screenshot%202025-12-30%20at%203.30.09%20PM.png)
*Grafana login interface accessible via port-forward on localhost:3000*

### Infrastructure Monitoring Dashboards

![Infrastructure Overview Dashboard](screenshot/grafana/Screenshot%202025-12-30%20at%203.35.19%20PM.png)
*Infrastructure dashboard showing real-time node CPU, memory, disk usage, and network I/O metrics across all 5 EKS nodes*

![Node Metrics Detail](screenshot/grafana/Screenshot%202025-12-30%20at%203.35.55%20PM.png)
*Detailed node-level metrics with individual graphs for each t3.medium instance*

### Application Performance Dashboards

![Application Performance Dashboard](screenshot/grafana/Screenshot%202025-12-30%20at%203.36.15%20PM.png)
*Application dashboard displaying pod CPU/memory usage, restart counts, and service availability metrics*

![Pod Metrics Detail](screenshot/grafana/Screenshot%202025-12-30%20at%203.36.48%20PM.png)
*Detailed pod-level metrics showing resource consumption for all microservices (frontend, catalogue, voting, recommendation)*

