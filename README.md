# Craftista Application - Deployment Guide

[![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-blue)](https://github.com/Abhin-Anilkumar/K8s-demo/actions)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-1.28+-326CE5)](https://kubernetes.io/)
[![Helm](https://img.shields.io/badge/Helm-3.0+-0F1689)](https://helm.sh/)

> **For application details and features**, see [README-original.md](README-original.md)
> 
> **Infrastructure Repository**: [EKS-project-for-8byte](https://github.com/Abhin-Anilkumar/EKS-project-for-8byte) - Terraform IaC for AWS EKS cluster

---

## 📋 Quick Links

- **[Application Details](README-original.md)** - About Craftista, features, and microservices
- **[APPROACH.md](APPROACH.md)** - Design rationale and architectural decisions
- **[CHALLENGES.md](CHALLENGES.md)** - Issues encountered and resolutions
- **[DASHBOARD.md](DASHBOARD.md)** - Monitoring and observability strategy

---

## 🚀 Quick Start

### Prerequisites

- **kubectl** configured for EKS cluster
- **Helm** >= 3.0
- **AWS CLI** with ECR access

### 1. Create Database Secret

```bash
# Production namespace
kubectl create secret generic voting-db-credentials \
  --from-literal=username=postgres \
  --from-literal=password=<RDS_PASSWORD> \
  -n app

# Staging namespace
kubectl create secret generic voting-db-credentials \
  --from-literal=username=postgres \
  --from-literal=password=<RDS_PASSWORD> \
  -n stage-app
```

### 2. Deploy Services (Production)

```bash
# Deploy all microservices
helm upgrade --install frontend ./charts/frontend -n app
helm upgrade --install catalogue ./charts/catalogue -n app
helm upgrade --install voting ./charts/voting -n app
helm upgrade --install recommendation ./charts/recommendation -n app
```

### 3. Verify Deployment

```bash
kubectl get pods -n app
kubectl get svc -n app
```

---

## 🔄 CI/CD Pipeline

### Branching Strategy

- **`main`**: Production deployments (manual approval required)
- **`develop`**: Staging deployments (automatic)
- **PRs to `develop`**: Build validation (Docker build only, no unit tests)

### Pipeline Stages

| Stage | Trigger | Actions | Target |
|-------|---------|---------|--------|
| **Build Validation** | PR to `develop` | Docker build all services | - |
| **Build & Push** | Push to `develop`/`main` | Build + Push to ECR | ECR |
| **Deploy Staging** | Push to `develop` | Helm upgrade | `stage-app` namespace |
| **Deploy Production** | Push to `main` | Manual approval + Helm upgrade | `app` namespace |

---

## 📊 Monitoring

### Grafana Dashboards

**Access Grafana**:
```bash
kubectl port-forward -n monitoring svc/grafana 3000:80
```

**Import Dashboards**:
1. Navigate to **Dashboards → Import**
2. Upload JSON files from `grafana-dashboards/`:
   - `infrastructure-overview.json` - Node metrics, disk, network
   - `application-performance.json` - Pod CPU/memory, restarts, availability

**Configure Data Source**:
- URL: `http://prometheus-server.monitoring.svc.cluster.local`
- Access: Server (default)

### CloudWatch Logs

```bash
# View control plane logs
aws logs tail /aws/eks/prod-eks/cluster --follow
```

---

## 🔒 Security & Secrets

### Secret Management

**Database Credentials**: Stored as Kubernetes Secrets
- Namespace-scoped (`app`, `stage-app`)
- RBAC: Only voting pods can access
- Encrypted at rest via EKS encryption config

**Best Practices**:
- Never commit secrets to Git
- Use AWS Secrets Manager for production (recommended)
- Rotate credentials regularly

---

## 📚 Documentation

- **[Complete Walkthrough](https://github.com/Abhin-Anilkumar/EKS-project-for-8byte/blob/main/terraform/complete-walkthrough.md)**: **Step-by-step technical guide** - Detailed explanation of infrastructure, application, CI/CD, and deployment
- **[README-original.md](README-original.md)** - Application details and features
- **[APPROACH.md](APPROACH.md)** - Design rationale
- **[CHALLENGES.md](CHALLENGES.md)** - Issues and resolutions
- **[DASHBOARD.md](DASHBOARD.md)** - Monitoring strategy
- **[Infrastructure README](https://github.com/Abhin-Anilkumar/EKS-project-for-8byte/blob/main/README.md)** - EKS setup guide

---

## 👤 Author

**Abhin Anilkumar**
- GitHub: [@Abhin-Anilkumar](https://github.com/Abhin-Anilkumar)
- Application: [K8s-demo](https://github.com/Abhin-Anilkumar/K8s-demo)
- Infrastructure: [EKS-project-for-8byte](https://github.com/Abhin-Anilkumar/EKS-project-for-8byte)
