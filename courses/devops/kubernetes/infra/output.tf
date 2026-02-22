output "next_steps" {
  value = <<EOT

    Cluster deployed successfully!
    To access the cluster, run the following command:
    az aks get-credentials --resource-group ${var.resource_group_name} --name ${var.aks_cluster_name}

    Access to public ip created:
    http://${azurerm_public_ip.ingress_ip.ip_address}

  EOT
}

output "resource_group_name" {
  value = var.resource_group_name
}

output "aks_cluster_name" {
  value = var.aks_cluster_name
}

