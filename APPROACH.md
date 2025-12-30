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

### 5. Repository & CI/CD Organization (Refactored)
- **Centralized Helm Charts**: Reorganized Helm charts into a dedicated `charts/` directory at the project root. This follows Kubernetes best practices, making it easier to manage deployments across multiple environments.
- **Streamlined CI/CD**: Simplified the GitHub Actions pipeline to focus on high-speed delivery:
  - **Build**: Rapid Docker builds with Amazon ECR integration.
  - **Deploy**: Seamless Helm-based deployments to EKS.
  - **Notify**: Reliable feedback via native GitHub status checks.

### 6. Security & IAM
- **IRSA (IAM Roles for Service Accounts)**: Implemented for the AWS Load Balancer Controller to follow the principle of least privilege.
- **Secret Management**: Database credentials and other sensitive info are managed via Kubernetes Secrets.

### 7. Monitoring & Logging
- **AWS CloudWatch**: Used for infrastructure metrics and centralized log streaming.
- **Metrics Server**: Enabled for real-time cluster resource monitoring.
