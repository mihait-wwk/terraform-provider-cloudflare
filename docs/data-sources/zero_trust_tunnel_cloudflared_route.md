---
page_title: "cloudflare_zero_trust_tunnel_cloudflared_route Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_tunnel_cloudflared_route (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Cloudflare account ID
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `route_id` (String) UUID of the route.
- `tun_type` (String) The type of tunnel.
- `tunnel_name` (String) A user-friendly name for a tunnel.
- `virtual_network_name` (String) A user-friendly name for the virtual network.

### Read-Only

- `comment` (String) Optional remark describing the route.
- `created_at` (String) Timestamp of when the resource was created.
- `deleted_at` (String) Timestamp of when the resource was deleted. If `null`, the resource has not been deleted.
- `id` (String) UUID of the route.
- `network` (String) The private IPv4 or IPv6 range connected by the route, in CIDR notation.
- `tunnel_id` (String) UUID of the tunnel.
- `virtual_network_id` (String) UUID of the virtual network.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String) Cloudflare account ID

Optional:

- `comment` (String) Optional remark describing the route.
- `existed_at` (String) If provided, include only tunnels that were created (and not deleted) before this time.
- `is_deleted` (Boolean) If `true`, only include deleted routes. If `false`, exclude deleted routes. If empty, all routes will be included.
- `network_subset` (String) If set, only list routes that are contained within this IP range.
- `network_superset` (String) If set, only list routes that contain this IP range.
- `route_id` (String) UUID of the route.
- `tun_types` (String) The types of tunnels to filter separated by a comma.
- `tunnel_id` (String) UUID of the tunnel.
- `virtual_network_id` (String) UUID of the virtual network.

