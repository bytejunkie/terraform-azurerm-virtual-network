resource "azurerm_virtual_network" "vnet" {
  name                = join(var.name_separator, var.vnet_name_strings)
  location            = var.location
  resource_group_name = var.resource_group_name
  address_space       = var.vnet_address_space

  dns_servers = var.vnet_dns_servers != "" ? var.vnet_dns_servers : null
  tags        = var.tags != "" ? var.tags : null

}
