variable "subscription_id" {
  sensitive = true
}

variable "resource_group_name" {

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
