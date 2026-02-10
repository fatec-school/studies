output "vm_name" {
  description = "Created VM name"
  value       = azurerm_linux_virtual_machine.vm.name
}

output "public_ip_address" {
  description = "VM public IP address"
  value       = azurerm_public_ip.public_ip.ip_address
}

output "network_interface_name" {
  description = "Network Interface Card (NIC) name"
  value       = azurerm_network_interface.nic.name
}

 