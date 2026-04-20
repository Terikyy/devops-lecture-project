# Call the cluster's action to retrieve the raw Admin Kubeconfig
data "azapi_resource_action" "aks_kubeconfig" {
  type                   = "Microsoft.ContainerService/managedClusters@2023-01-01"
  resource_id            = azapi_resource.aks.id
  action                 = "listClusterAdminCredential"
  response_export_values = ["kubeconfigs"]
  depends_on             = [azapi_resource.aks]
}

# Decode the output locally to supply the nested Helm/Kubernetes providers
locals {
  kubeconfig_raw = jsondecode(data.azapi_resource_action.aks_kubeconfig.output).kubeconfigs[0].value
  kubeconfig     = yamldecode(base64decode(local.kubeconfig_raw))
  cluster        = local.kubeconfig["clusters"][0]["cluster"]
  user           = local.kubeconfig["users"][0]["user"]
}

# Dynamically initialize Helm using AKS credentials
provider "helm" {
  kubernetes {
    host                   = local.cluster["server"]
    client_certificate     = base64decode(local.user["client-certificate-data"])
    client_key             = base64decode(local.user["client-key-data"])
    cluster_ca_certificate = base64decode(local.cluster["certificate-authority-data"])
  }
}

# Dynamically initialize Kubernetes using AKS credentials
provider "kubernetes" {
  host                   = local.cluster["server"]
  client_certificate     = base64decode(local.user["client-certificate-data"])
  client_key             = base64decode(local.user["client-key-data"])
  cluster_ca_certificate = base64decode(local.cluster["certificate-authority-data"])
}

# Define the namespace for ArgoCD
resource "kubernetes_namespace" "argocd" {
  metadata {
    name = "argocd"
  }
}

# Workaround: Ensure Helm Repositories are initialized locally for other PCs
# This avoids the "no cached repo found" error when Helm pulls Bitnami Redis dependencies
resource "null_resource" "init_helm_repos" {
  triggers = {
    repository = "https://argoproj.github.io/argo-helm"
  }

  provisioner "local-exec" {
    command = "helm repo add bitnami https://charts.bitnami.com/bitnami && helm repo add argo https://argoproj.github.io/argo-helm && helm repo update"
  }
}

# Install ArgoCD via Helm
resource "helm_release" "argocd" {
  name       = "argocd"
  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"
  version    = "7.0.0" # Fixed version for stability
  namespace  = kubernetes_namespace.argocd.metadata[0].name
  wait       = true # Ensures resources are up before next steps

  # Expose the server via an Azure LoadBalancer
  set {
    name  = "server.service.type"
    value = "LoadBalancer"
  }

  depends_on = [
    azapi_resource.aks,
    null_resource.init_helm_repos
  ]
}

# Retrieve the LoadBalancer Service info and Admin Secret 
data "kubernetes_service" "argocd_server" {
  metadata {
    name      = "argocd-server"
    namespace = kubernetes_namespace.argocd.metadata[0].name
  }
  depends_on = [helm_release.argocd]
}

# Bootstrap your base root Application (argocd/root.yml)
resource "local_file" "kube_config" {
  content  = base64decode(local.kubeconfig_raw)
  filename = "${path.module}/.kubeconfig"
}

resource "null_resource" "wait_for_crd" {
  depends_on = [helm_release.argocd, local_file.kube_config]

  provisioner "local-exec" {
    # Check explicitly for CRD readiness
    command = "kubectl --kubeconfig=${local_file.kube_config.filename} wait --for condition=established --timeout=120s crd/applications.argoproj.io"
  }
}

resource "null_resource" "deploy_argocd_root" {
  triggers = {
    # Update resource if the manifest changes
    manifest_sha1 = sha1(file("${path.module}/../argocd/root.yml"))
  }

  depends_on = [null_resource.wait_for_crd]

  provisioner "local-exec" {
    # Use kubectl to avoid terraform plan API validation issues
    command = "kubectl --kubeconfig=${local_file.kube_config.filename} apply -f ${path.module}/../argocd/root.yml"
  }
}
