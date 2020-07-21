#################
# backend
#################

terraform {
  backend "s3" {}
  required_version = ">= 0.12"
}


#################
# pre-requisites
#################

module "resourcegroup" {
  source = "bytejunkie/resource-group/azurerm"

  name_strings   = var.name_strings
  name_separator = var.name_separator

  tags = var.tags
}



module "virtual_network" {
  source = "../"
  # insert the required variables here
  name_strings        = var.name_strings
  name_separator      = var.name_separator
  vnet_name_strings   = var.vnet_name_strings
  resource_group_name = module.resourcegroup.resource_group_name
  vnet_address_space  = var.vnet_address_space

  # insert the optional variables here
  vnet_dns_servers = var.vnet_dns_servers
  tags             = var.tags

  depends_on = [
    module.resourcegroup
  ]

}
