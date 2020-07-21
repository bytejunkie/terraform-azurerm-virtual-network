#  AzureRM_VirtualNetwork

## Background
This module deploys the Virtual Network Resource

## usage

Its quite easy to use the module, just by supplying the required parameters. There is (as with most Azure Modules) a pre-requisite for a resource group, so this either needs to be a data resource or another modular deployment, which is how the example code below was written.

```
module "virtual_network" {
  source = "../"
  # insert the required variables here
  name_strings   = var.name_strings
  name_separator = var.name_separator
  resource_group_name = module.resourcegroup.resource_group_name
  vnet_name_strings  = var.vnet_name_strings
  vnet_address_space = var.vnet_address_space

# insert the optional variables here
  vnet_dns_servers   = var.vnet_dns_servers
  tags = var.tags

  depends_on = [
    module.resourcegroup
  ]
}
```

## Required Parameters

``` hcl
variable "name_strings" {
  description = "This should be a list of strings which in conjunction with the seperator make up the resource group name"
  default     = null
}

variable "name_separator" {
  description = "Used with name_strings to make up the resource group name"
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
```

## Optional Parameters

```hcl
variable "vnet_dns_servers" {
  type        = list
  description = "[optional]list of the dns servers to use with the virtual network"
  default     = null
}
```
