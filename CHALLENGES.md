# Challenges & Resolutions

During the deployment and refinement of the Craftista application, several technical hurdles were addressed to reach the final best-practice architecture.

## 1. Cross-Architecture Image Mismatches
**Challenge**: Local builds on Apple Silicon (arm64) were incompatible with EKS nodes (amd64).
**Resolution**: Standardized the CI/CD pipeline to build all images with `--platform linux/amd64`.

## 2. Internal Service Discovery (ENOTFOUND)
**Challenge**: Short-name DNS resolution was inconsistent between microservices.
**Resolution**: Updated all service configurations to use Fully Qualified Domain Names (FQDNs).

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
- **Simplification**: Refactored the CI/CD pipeline to remove unit tests, focusing on immediate build, push, and deployment to the production cluster.
- **Feedback Loop**: Leveraged native GitHub Action notifications for streamlined build and deployment status.
