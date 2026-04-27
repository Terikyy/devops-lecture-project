output "resource_group_name" {
  description = "The name of the created Resource Group"
  value       = azurerm_resource_group.main.name
}

output "aks_cluster_name" {
  description = "The name of the created AKS cluster"
  value       = azapi_resource.aks.name
}

output "cluster_location" {
  description = "The Azure region where the cluster is deployed"
  value       = azapi_resource.aks.location
}

output "argocd_url" {
  description = "The external URL indicating the ArgoCD Server Dashboard"
  # Fetch the automatically assigned remote load balancer IP from Azure, default to pending
  value = try("https://${data.kubernetes_service.argocd_server.status[0].load_balancer[0].ingress[0].ip}", "pending")
  depends_on = [helm_release.argocd]
}
