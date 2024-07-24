---
page_title: "cloudflare_zone_cache_reserve Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zone_cache_reserve (Resource)



## Example Usage

```terraform
resource "cloudflare_zone_cache_reserve" "example" {
  zone_id = "0da42c8d2132a9ddaf714f9e7c920711"
  enabled = true
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) ID of the zone setting.
- `zone_id` (String) Identifier

### Optional

- `value` (String) Value of the Cache Reserve zone setting.

### Read-Only

- `modified_on` (String) last time this setting was modified.

## Import

Import is supported using the following syntax:

```shell
$ terraform import cloudflare_zone_cache_reserve.example <zone_id>
```