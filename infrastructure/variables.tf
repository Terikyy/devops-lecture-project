variable "location" {
  description = "The Azure region where resources will be created"
  type        = string
  default     = "westeurope"
}

variable "resource_group_name" {
  description = "The name of the Azure Resource Group"
  type        = string
}

variable "aks_cluster_name" {
  description = "The name of the AKS cluster"
  type        = string
}

variable "node_count" {
  description = "The number of nodes in the default node pool"
  type        = number
  default     = 2
}

variable "vm_size" {
  description = "VM size for AKS nodes"
  type        = string
}

variable "kubernetes_version" {
  description = "Kubernetes version (major.minor oder major.minor.patch)"
  type        = string
  default     = "1.34"
}
