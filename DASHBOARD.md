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
- `/aws/containerinsights/prod-eks/application`: Application container logs

### Log Sources
- Application STDOUT/STDERR from all microservices
- System/Node logs
- EKS control plane audit logs

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
