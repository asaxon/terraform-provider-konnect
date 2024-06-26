# Changelog

# Changelog

## 0.2.3
> Released on 2024/05/14

### Features
* Add support for the `konnect_cloud_gateway_configuration`, `konnect_cloud_gateway_custom_domain`, `konnect_cloud_gateway_network` and `konnect_cloud_gateway_transit_gateway` resources

### Bug Fixes
* Fix `konnect_portal_product_version` creation bug

## 0.2.2
> Released on 2024/05/04

### Features
* Add support for the `konnect_system_account`, `konnect_system_account_access_token`, `konnect_system_account_role` and `konnect_system_account_team` resources

## 0.2.1
> Released on 2024/04/26

### Features
* Add support for the `konnect_gateway_data_plane_client_certificate` resource

## 0.2.0
> Released on 2024/04/18

### BREAKING CHANGES
* The provider `server_url` no longer contains an API version. Update your provider configuration to use `https://us.api.konghq.com` (or your chosen region)
* `labels` on the `konnect_gateway_control_plane` resource now accept a map of values rather than using `jsonencode()`

### Fixes
* Mesh control planes can be created in regions other than NA
* Gateway control planes now accept `CLUSTER_TYPE_CONTROL_PLANE` in the `cluster_type` field

## 0.1.0
> Released on 2024/04/16

* Initial release