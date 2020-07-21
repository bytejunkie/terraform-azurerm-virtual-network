variable "name_strings" {
  description = "This should be a list of strings which in conjunction with the seperator make up the resource group name"
  default     = null
}

variable "name_separator" {
  description = "Used with name_strings to make up the resource group name"
  default     = null
}

variable "resource_group_name" {
  description = "the resource group name which is used to deploy the virtual network"
  default     = null
}

variable "location" {
  description = "Used to define the location in which to deploy the resource"
  default     = "West Europe"
}

variable "vnet_name_strings" {
  description = "Used with name_strings to make up the virtual network name"
  default     = null
}

variable "vnet_address_space" {
  type        = list
  description = "the ipv4 address space to use with the virtual network"
  default     = null
}

variable "vnet_dns_servers" {
  type        = list
  description = "[optional]list of the dns servers to use with the virtual network"
  default     = null
}

variable "tags" {
  type    = map
  default = null
}

variable "module_depends_on" {
  type        = any
  description = "A list of external resources the module depends_on"
  default     = []
}
