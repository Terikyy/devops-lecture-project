locals {
  common_tags = {
    environment = "dev"
    project     = "devops-lecture-project"
  }
}

resource "azurerm_resource_group" "main" {
  name     = var.resource_group_name
  location = var.location

  tags = local.common_tags
}

resource "azapi_resource" "aks" {
  type      = "Microsoft.ContainerService/managedClusters@2023-01-01"
  name      = var.aks_cluster_name
  parent_id = azurerm_resource_group.main.id
  location  = azurerm_resource_group.main.location

  body = jsonencode({
    properties = {
      kubernetesVersion = var.kubernetes_version
      dnsPrefix         = "${var.aks_cluster_name}-dns"

      agentPoolProfiles = [
        {
          name   = "default"
          count  = var.node_count
          vmSize = var.vm_size
          mode   = "System"
          type   = "VirtualMachineScaleSets"
        }
      ]

      servicePrincipalProfile = {
        clientId = "msi"
      }
    }

    identity = {
      type = "SystemAssigned"
    }

    tags = local.common_tags
  })
}