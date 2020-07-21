output "virtual_network_name" {
  value = azurerm_virtual_network.vnet.name
}
output "virtual_network_location" {
  value = azurerm_virtual_network.vnet.location
}
output "virtual_network_address_space" {
  value = azurerm_virtual_network.vnet.address_space
}
output "virtual_network_dns_servers" {
  value = azurerm_virtual_network.vnet.dns_servers
}
output "virtual_network_tags" {
  value = azurerm_virtual_network.vnet.tags
}
