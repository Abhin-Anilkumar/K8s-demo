# Project Approach

This document outlines the architectural decisions and technical rationale for the Craftista EKS deployment.

## Architectural Overview

The goal was to deploy a highly available, scalable, and secure microservices application on AWS EKS. The architecture leverages managed services wherever possible to minimize operational overhead.

### 1. Network Infrastructure (VPC)
- **Multi-AZ Deployment**: Subnets are distributed across multiple Availability Zones to ensure high availability.
- **Segmentation**: Public subnets host the ALB, while private subnets host worker nodes and RDS instances.
- **ALB Discovery**: Subnets are correctly tagged for AWS Load Balancer Controller discovery.

### 2. EKS Cluster Design
- **Unified Node Group**: Consolidated workloads into a single node group to simplify networking and security group management.
- **Architecture Compatibility**: Images are built for `linux/amd64` to match the EKS worker node architecture.

### 3. Service Discovery & Connectivity
- **Internal FQDNs**: Microservices use Full Qualified Domain Names (e.g., `catalogue.app.svc.cluster.local`) for reliable internal communication.

### 4. Persistence
- **RDS PostgreSQL**: The `voting` service is integrated with a managed RDS PostgreSQL instance for data persistence.

### 5. Repository & CI/CD Organization (Main/Develop)
- **Centralized Helm Charts**: Reorganized Helm charts into a dedicated `charts/` directory at the project root for consistent deployments.
- **Microservices Branching**:
  - **Develop**: Targeted for `stage-app` namespace.
  - **Main**: Targeted for `app` namespace with manual approval.
- **Streamlined CI/CD**: 
  - **Validation**: "Build-only" checks on PRs to ensure code integrity (no unit tests).
  - **Staging**: Automatic deployment to `stage-app`.
  - **Production**: Gated deployment to `app`.

### 6. Security & IAM
- **IRSA (IAM Roles for Service Accounts)**: Implemented for the AWS Load Balancer Controller to follow the principle of least privilege.
- **Secret Management**: Database credentials and other sensitive info are managed via Kubernetes Secrets.

### 7. Security Best Practices
- **Network Isolation**: Worker nodes and databases reside in private subnets; only the ALB is public.
- **Least Privilege**: IAM Roles for Service Accounts (IRSA) restrict pod permissions.
- **Secret Management**: Sensitive data (DB credentials) is managed via Kubernetes Secrets, not hardcoded.
- **Container Security**: Images are built distroless/minimal where possible to reduce attack surface.

### 8. Cost Optimization
- **Right-Sizing**: Selected `t3.medium` instances to balance performance and cost for this workload.
- **Auto-Scaling**: Configured Horizontal Pod Autoscalers (HPA) to scale down during low traffic.
- **Spot Instances**: (Optional) Ready to leverage Spot instances for stateless frontend/catalogue services.

### 9. Reliability & Backups
- **Database Backups**: RDS is configured with automated daily snapshots (7-day retention).
- **State Management**: Terraform state is locked via DynamoDB to prevent corruption.
- **High Availability**: Multi-AZ deployment ensures resilience against zone failures.

### 10. Monitoring & Observability
- **Prometheus + Grafana**: Deployed in `monitoring` namespace for comprehensive metrics collection
  - **Infrastructure Metrics**: Node CPU, memory, disk, network I/O
  - **Application Metrics**: Pod resource usage, restart counts, availability
  - **Dashboards**: Pre-configured Grafana dashboards for infrastructure and application monitoring
- **CloudWatch Logs**: EKS control plane logs (API server, scheduler, controller manager)
  - Log Group: `/aws/eks/prod-eks/cluster`
- **EBS CSI Driver**: Installed as EKS addon to enable persistent storage for Prometheus/Grafana

### 11. Cluster Scaling
- **Node Group Scaling**: Configured with min=3, max=10, desired=5 nodes
- **Instance Type**: t3.medium (2 vCPU, 4GB RAM) for balanced performance
- **Auto-Scaling**: Cluster Autoscaler enabled for dynamic scaling based on workload

### 12. Application Exposure
- **Production URL**: [evoqu.in](http://evoqu.in)
- **Exposure Method**: Kubernetes LoadBalancer service + DNS (Hostinger)
- **ALB Limitation**: AWS account restrictions prevent ALB creation, requiring direct LoadBalancer service exposure
- **DNS Configuration**: Domain `evoqu.in` points to LoadBalancer external IP
