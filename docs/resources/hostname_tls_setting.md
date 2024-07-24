---
page_title: "cloudflare_hostname_tls_setting Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_hostname_tls_setting (Resource)



## Example Usage

```terraform
resource "cloudflare_hostname_tls_setting" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  hostname = "sub.example.com"
  setting  = "min_tls_version"
  value    = "1.2"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `setting_id` (String) The TLS Setting name.
- `value` (Number) The tls setting value.
- `zone_id` (String) Identifier

### Optional

- `hostname` (String) The hostname for which the tls settings are set.

### Read-Only

- `created_at` (String) This is the time the tls setting was originally created for this hostname.
- `status` (String)
- `updated_at` (String) This is the time the tls setting was updated.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_hostname_tls_setting.example <zone_id>/<hostname>/<setting_name>
```