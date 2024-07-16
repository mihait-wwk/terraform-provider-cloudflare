---
page_title: "cloudflare_access_mutual_tls_certificate Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_access_mutual_tls_certificate (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `associated_hostnames` (String) The hostnames of the applications that will use this certificate.
- `certificate_id` (String) UUID
- `created_at` (String)
- `expires_on` (String)
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `fingerprint` (String) The MD5 fingerprint of the certificate.
- `id` (String) The ID of the application that will use this certificate.
- `name` (String) The name of the certificate.
- `updated_at` (String)
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Optional:

- `account_id` (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
- `zone_id` (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.

