output "finish" {
  value = <<EOF

    Congrats! Yours virtual machines are finished.

  EOF
}

output "vm1_public_ip" {
  value = azurerm_linux_virtual_machine.vm1.public_ip_address
}

output "vm2_public_ip" {
  value = azurerm_linux_virtual_machine.vm2.public_ip_address
}

