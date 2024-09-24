---
page_title: "cloudflare_zero_trust_device_posture_integration Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_device_posture_integration (Resource)



## Example Usage

```terraform
resource "cloudflare_zero_trust_device_posture_integration" "example" {
  account_id = "f037e56e89293a057740de681ac9abbe"
  name       = "Device posture integration"
  type       = "workspace_one"
  interval   = "24h"
  config = {
    api_url       = "https://example.com/api"
    auth_url      = "https://example.com/connect/token"
    client_id     = "client-id"
    client_secret = "client-secret"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `config` (Attributes) The configuration object containing third-party integration information. (see [below for nested schema](#nestedatt--config))
- `interval` (String) The interval between each posture check with the third-party API. Use `m` for minutes (e.g. `5m`) and `h` for hours (e.g. `12h`).
- `name` (String) The name of the device posture integration.
- `type` (String) The type of device posture integration.

### Read-Only

- `id` (String) API UUID.

<a id="nestedatt--config"></a>
### Nested Schema for `config`

Optional:

- `access_client_id` (String) If present, this id will be passed in the `CF-Access-Client-ID` header when hitting the `api_url`
- `access_client_secret` (String) If present, this secret will be passed in the `CF-Access-Client-Secret` header when hitting the `api_url`
- `api_url` (String) The Workspace One API URL provided in the Workspace One Admin Dashboard.
- `auth_url` (String) The Workspace One Authorization URL depending on your region.
- `client_id` (String) The Workspace One client ID provided in the Workspace One Admin Dashboard.
- `client_key` (String) The Uptycs client secret.
- `client_secret` (String) The Workspace One client secret provided in the Workspace One Admin Dashboard.
- `customer_id` (String) The Crowdstrike customer ID.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zero_trust_device_posture_integration.example <account_id>/<device_posture_integration_id>
```