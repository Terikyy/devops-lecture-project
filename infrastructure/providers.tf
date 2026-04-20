terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.65.0"
    }

    azapi = {
      source  = "azure/azapi"
      version = "~> 1.10"
    }

    helm = {
      source  = "hashicorp/helm"
      version = "~> 2.12.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 2.24.0"
    }
  }
}

provider "azurerm" {
  features {}
  skip_provider_registration = true
}

provider "azapi" {}