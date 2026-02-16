output "public_ip_address" {
  value       = azurerm_public_ip.public_ip.ip_address
  description = "VM public IP"
}

output "nsg_name" {
  value       = azurerm_network_security_group.nsg.name
  description = "Network Security Group Name"
}

output "resource_group_name" {
  value = azurerm_resource_group.rg.name
}
