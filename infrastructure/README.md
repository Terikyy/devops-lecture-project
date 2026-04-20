# Infrastructure as Code (OpenTofu)

This directory contains the Infrastructure as Code (IaC) configuration used to provision and manage cloud resources on Microsoft Azure. The setup is implemented using OpenTofu (Terraform-compatible).

## Provisioned Resources

The configuration defines the following components:

- **Azure Resource Group**  
  Serves as a logical container for all deployed resources.

- **Azure Kubernetes Service (AKS)**  
  Managed Kubernetes cluster configured with a system-assigned managed identity.  
  Provides the execution environment for containerized workloads.

- **ArgoCD (GitOps)**  
  Installed via Helm within the AKS cluster.
  Uses local `helm` to initialize dependencies seamlessly. Uses `kubectl` to actively wait for Custom Resource Definitions (CRDs) to establish before applying objects.
  Bootstraps the base root Application (`argocd/root.yml`) to enable GitOps deployments securely and without plan validation errors.

## Prerequisites

Ensure the following requirements are met before deployment:

- OpenTofu (>= 1.5.0)
- Azure CLI (`az`)
- `kubectl` CLI (used dynamically by Terraform to validate CRDs)
- `helm` CLI (used dynamically by Terraform for repository initialization)
- Active Azure subscription
- Authenticated Azure session (`az login`)

## Usage

### 1. Authenticate with Azure

```bash
az login
```

### 2. Initialize the Environment

Initializes the working directory and downloads required providers.

```bash
tofu init
```

### 3. Validate Configuration

Checks the configuration for syntax and internal consistency.

```bash
tofu validate
```

### 4. Review Execution Plan

Displays the planned changes before applying them.

```bash
tofu plan
```

### 5. Apply Configuration

Provisions the defined infrastructure in Azure.

```bash
tofu apply
```

### 6. Access the Cluster

Retrieve credentials to interact with the Kubernetes cluster:

```bash
az aks get-credentials \
  --resource-group <resource_group_name> \
  --name <aks_cluster_name>
```

### 7. Destroy Resources

Removes all provisioned resources:

```bash
tofu destroy
```