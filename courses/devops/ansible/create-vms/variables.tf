variable "resource_group" {
  default = "ansible"
}

variable "location" {
  default = "East US"
}

variable "admin_username" {
  sensitive = true
}

variable "admin_password" {
  sensitive = true
}

