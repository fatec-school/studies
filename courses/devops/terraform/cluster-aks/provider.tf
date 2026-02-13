terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.1.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "3.0.1"
    }

    helm = {
      source  = "hashicorp/helm"
      version = "3.1.1"
    }
  }
}

provider "azurerm" {
  features {}
}

provider "kubernetes" {
  client_certificate     = base64decode()
  client_key             = base64decode()
  cluster_ca_certificate = base64decode()
}

provider "helm" {
  kubernetes = {
    client_certificate     = base64decode()
    client_key             = base64decode()
    cluster_ca_certificate = base64decode()
  }
}
