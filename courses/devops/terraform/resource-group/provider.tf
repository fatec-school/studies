# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {
  }
  subscription_id = var.subscription_id
}

terraform {
  # Specify the required providers
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=4.1.0"
    }
  }

  # Specify the Terraform version
  required_version = ">=1.0"
}

