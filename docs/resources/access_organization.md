---
page_title: "cloudflare_access_organization Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_organization (Resource)



## Example Usage

```terraform
resource "cloudflare_access_organization" "example" {
  account_id                         = "f037e56e89293a057740de681ac9abbe"
  name                               = "example.cloudflareaccess.com"
  auth_domain                        = "example.cloudflareaccess.com"
  is_ui_read_only                    = false
  user_seat_expiration_inactive_time = "720h"
  auto_redirect_to_identity          = false

  login_design = {
    background_color = "#ffffff"
    text_color       = "#000000"
    logo_path        = "https://example.com/logo.png"
    header_text      = "My header text"
    footer_text      = "My footer text"
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `auth_domain` (String) The unique subdomain assigned to your Zero Trust organization.
- `name` (String) The name of your Zero Trust organization.

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `allow_authenticate_via_warp` (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
- `auto_redirect_to_identity` (Boolean) When set to `true`, users skip the identity provider selection step during login.
- `custom_pages` (Attributes) (see [below for nested schema](#nestedatt--custom_pages))
- `is_ui_read_only` (Boolean) Lock all settings as Read-Only in the Dashboard, regardless of user permission. Updates may only be made via the API or Terraform for this account when enabled.
- `login_design` (Attributes) (see [below for nested schema](#nestedatt--login_design))
- `session_duration` (String) The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
- `ui_read_only_toggle_reason` (String) A description of the reason why the UI read only field is being toggled.
- `user_seat_expiration_inactive_time` (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count. Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
- `warp_auth_session_duration` (String) The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

### Read-Only

- `created_at` (String)
- `id` (String) The name of your Zero Trust organization.
- `updated_at` (String)

<a id="nestedatt--custom_pages"></a>
### Nested Schema for `custom_pages`

Optional:

- `forbidden` (String) The uid of the custom page to use when a user is denied access after failing a non-identity rule.
- `identity_denied` (String) The uid of the custom page to use when a user is denied access.


<a id="nestedatt--login_design"></a>
### Nested Schema for `login_design`

Optional:

- `background_color` (String) The background color on your login page.
- `footer_text` (String) The text at the bottom of your login page.
- `header_text` (String) The text at the top of your login page.
- `logo_path` (String) The URL of the logo on your login page.
- `text_color` (String) The text color on your login page.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_access_organization.example <account_id>
```