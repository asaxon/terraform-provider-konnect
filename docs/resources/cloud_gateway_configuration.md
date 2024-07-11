---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "konnect_cloud_gateway_configuration Resource - terraform-provider-konnect"
subcategory: ""
description: |-
  CloudGatewayConfiguration Resource
---

# konnect_cloud_gateway_configuration (Resource)

CloudGatewayConfiguration Resource

## Example Usage

```terraform
resource "konnect_cloud_gateway_configuration" "my_cloudgatewayconfiguration" {
  api_access        = "private"
  configuration_id  = "edaf40f9-9fb0-4ffe-bb74-4e763a6bd471"
  control_plane_geo = "us"
  control_plane_id  = "0949471e-b759-45ba-87ab-ee63fb781388"
  dataplane_groups = [
    {
      autoscale = {
        configuration_data_plane_group_autoscale_autopilot = {
          base_rps = 1
          kind     = "autopilot"
          max_rps  = 1000
        }
      }
      cloud_gateway_network_id = "36ae63d3-efd1-4bec-b246-62aa5d3f5695"
      created_at               = "2022-11-04T20:10:06.927Z"
      egress_ip_addresses = [
        "...",
      ]
      id = "cbb8872a-1f83-4806-bf69-fdf0b4783c7e"
      private_ip_addresses = [
        "...",
      ]
      provider   = "aws"
      region     = "us-east-2"
      state      = "created"
      updated_at = "2022-11-04T20:10:06.927Z"
    },
  ]
  version = "3.2"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `control_plane_geo` (String) Set of control-plane geos supported for deploying cloud-gateways configurations. must be one of ["us", "eu", "au"]
- `control_plane_id` (String)
- `dataplane_groups` (Attributes List) List of data-plane groups that describe where to deploy instances, along with how many instances. (see [below for nested schema](#nestedatt--dataplane_groups))
- `version` (String) Supported gateway version.

### Optional

- `api_access` (String) Type of API access data-plane groups will support for a configuration. must be one of ["private", "public", "private+public"]; Default: "private+public"

### Read-Only

- `created_at` (String) An RFC-3339 timestamp representation of configuration creation date.
- `dataplane_group_config` (Attributes List) Object that describes where data-planes will be deployed to, along with how many instances. (see [below for nested schema](#nestedatt--dataplane_group_config))
- `entity_version` (Number) Positive, monotonically increasing version integer, to serialize configuration changes.
- `id` (String) The ID of this resource.
- `updated_at` (String) An RFC-3339 timestamp representation of configuration update date.

<a id="nestedatt--dataplane_groups"></a>
### Nested Schema for `dataplane_groups`

Optional:

- `autoscale` (Attributes) Not Null (see [below for nested schema](#nestedatt--dataplane_groups--autoscale))
- `cloud_gateway_network_id` (String) Not Null
- `provider` (String) Name of cloud provider. Not Null; must be one of ["aws"]
- `region` (String) Region ID for cloud provider region. Not Null

Read-Only:

- `created_at` (String) An RFC-3339 timestamp representation of data-plane group creation date.
- `egress_ip_addresses` (List of String) List of egress IP addresses for the network that this data-plane group runs on.
- `id` (String) ID of the data-plane group that represents a deployment target for a set of data-planes.
- `private_ip_addresses` (List of String) List of private IP addresses of the internal load balancer that proxies traffic to this data-plane group.
- `state` (String) State of the data-plane group. must be one of ["created", "initializing", "ready", "terminating", "terminated"]
- `updated_at` (String) An RFC-3339 timestamp representation of data-plane group update date.

<a id="nestedatt--dataplane_groups--autoscale"></a>
### Nested Schema for `dataplane_groups.autoscale`

Optional:

- `configuration_data_plane_group_autoscale_autopilot` (Attributes) Object that describes the autopilot autoscaling strategy. (see [below for nested schema](#nestedatt--dataplane_groups--autoscale--configuration_data_plane_group_autoscale_autopilot))
- `configuration_data_plane_group_autoscale_static` (Attributes) Object that describes the static autoscaling strategy. (see [below for nested schema](#nestedatt--dataplane_groups--autoscale--configuration_data_plane_group_autoscale_static))

<a id="nestedatt--dataplane_groups--autoscale--configuration_data_plane_group_autoscale_autopilot"></a>
### Nested Schema for `dataplane_groups.autoscale.configuration_data_plane_group_autoscale_autopilot`

Optional:

- `base_rps` (Number) Base number of requests per second that the deployment target should support. Not Null
- `kind` (String) Not Null; must be one of ["autopilot"]
- `max_rps` (Number) Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.


<a id="nestedatt--dataplane_groups--autoscale--configuration_data_plane_group_autoscale_static"></a>
### Nested Schema for `dataplane_groups.autoscale.configuration_data_plane_group_autoscale_static`

Optional:

- `instance_type` (String) Instance type name to indicate capacity. Not Null; must be one of ["small", "medium", "large"]
- `kind` (String) Not Null; must be one of ["static"]
- `requested_instances` (Number) Number of data-planes the deployment target will contain. Not Null




<a id="nestedatt--dataplane_group_config"></a>
### Nested Schema for `dataplane_group_config`

Read-Only:

- `autoscale` (Attributes) (see [below for nested schema](#nestedatt--dataplane_group_config--autoscale))
- `cloud_gateway_network_id` (String)
- `provider` (String) Name of cloud provider. must be one of ["aws"]
- `region` (String) Region ID for cloud provider region.

<a id="nestedatt--dataplane_group_config--autoscale"></a>
### Nested Schema for `dataplane_group_config.autoscale`

Read-Only:

- `configuration_data_plane_group_autoscale_autopilot` (Attributes) Object that describes the autopilot autoscaling strategy. (see [below for nested schema](#nestedatt--dataplane_group_config--autoscale--configuration_data_plane_group_autoscale_autopilot))
- `configuration_data_plane_group_autoscale_static` (Attributes) Object that describes the static autoscaling strategy. (see [below for nested schema](#nestedatt--dataplane_group_config--autoscale--configuration_data_plane_group_autoscale_static))

<a id="nestedatt--dataplane_group_config--autoscale--configuration_data_plane_group_autoscale_autopilot"></a>
### Nested Schema for `dataplane_group_config.autoscale.configuration_data_plane_group_autoscale_autopilot`

Read-Only:

- `base_rps` (Number) Base number of requests per second that the deployment target should support.
- `kind` (String) must be one of ["autopilot"]
- `max_rps` (Number) Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.


<a id="nestedatt--dataplane_group_config--autoscale--configuration_data_plane_group_autoscale_static"></a>
### Nested Schema for `dataplane_group_config.autoscale.configuration_data_plane_group_autoscale_static`

Read-Only:

- `instance_type` (String) Instance type name to indicate capacity. must be one of ["small", "medium", "large"]
- `kind` (String) must be one of ["static"]
- `requested_instances` (Number) Number of data-planes the deployment target will contain.

## Import

Import is supported using the following syntax:

```shell
terraform import konnect_cloud_gateway_configuration.my_konnect_cloud_gateway_configuration "edaf40f9-9fb0-4ffe-bb74-4e763a6bd471"
```