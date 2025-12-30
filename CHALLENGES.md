# Challenges & Resolutions

During the deployment and refinement of the Craftista application, several technical hurdles were addressed to reach the final best-practice architecture.

## 1. Cross-Architecture Image Mismatches
**Challenge**: Local builds on Apple Silicon (arm64) were incompatible with EKS nodes (amd64).
**Resolution**: Standardized the CI/CD pipeline to build all images with `--platform linux/amd64`.

## 2. Internal Service Discovery (ENOTFOUND)
**Challenge**: Short-name DNS resolution was inconsistent between microservices.
**Resolution**: Updated all service configurations to use Fully Qualified Domain Names (FQDNs).
- **Feedback Loop**: Leveraged native GitHub Action notifications for streamlined build and deployment status.

## 6. Staging Deployment Failures
**Challenge**: The `voting` service failed to start in the `stage-app` namespace due to missing credentials.
**Resolution**: Identified that `voting-db-credentials` existed only in `app`. Replicated the secret to `stage-app` to resolve the `CreateContainerConfigError`.

## 3. Network Isolation
**Challenge**: Split node groups caused connectivity hangs due to isolated Security Groups.
**Resolution**: Consolidated all pods into a single managed node group sharing a common security context.

## 4. Persistence Gaps
**Challenge**: Data loss in the `voting` service during pod restarts.
**Resolution**: Integrated the service with RDS PostgreSQL, securing credentials via Kubernetes Secrets.

## 5. Repository Structure & CI/CD Complexity (Refactored)
**Challenge**: Scattered Helm charts and complex unit testing stages slowed down the deployment cycle.
**Resolution**:
- **Consolidation**: Moved all Helm charts to a centralized `charts/` directory for better maintainability.
- **Workflow Evolution**: Shifted to a professional **Main/Develop** branching strategy.
  - `pr` -> `develop`: Build Validation (No unit tests).
  - `push` -> `develop`: Automatic deployment to Staging (`stage-app`).
  - `push` -> `main`: Manual deployment to Production (`app`).

## 7. Monitoring Stack Persistent Storage
**Challenge**: Prometheus and Grafana pods stuck in Pending state due to unbound PersistentVolumeClaims.
**Resolution**: 
- Installed AWS EBS CSI Driver as an EKS addon with IRSA-enabled IAM role
- Set `gp2` StorageClass as default for automatic volume provisioning
- PVCs successfully bound and monitoring stack became operational

## 8. EKS Cluster Scaling
**Challenge**: Initial cluster had only 2 nodes, insufficient for monitoring stack and application workloads.
**Resolution**: 
- Scaled node group from 2 to 5 nodes (min=3, max=10, desired=5)
- Used AWS CLI to bypass Terraform validation error when min_size > current desired_size
- All monitoring and application pods successfully scheduled across 5 nodes

## 9. Application Exposure & ALB Restrictions
**Challenge**: AWS account has restrictions preventing Application Load Balancer creation.
**Resolution**: 
- Used Kubernetes LoadBalancer service type for frontend exposure
- Configured DNS record in Hostinger pointing `evoqu.in` to LoadBalancer external IP
- Application successfully accessible at [evoqu.in](http://evoqu.in)
- **Trade-off**: Direct LoadBalancer exposure instead of ALB with Ingress (less cost-effective but functional)
