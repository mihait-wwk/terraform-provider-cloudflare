---
page_title: "cloudflare_hyperdrive_config Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_hyperdrive_config (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `account_id` (String) Identifier
- `caching` (Attributes) (see [below for nested schema](#nestedatt--caching))
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `hyperdrive_id` (String) Identifier
- `name` (String)
- `origin` (Attributes) (see [below for nested schema](#nestedatt--origin))

<a id="nestedatt--caching"></a>
### Nested Schema for `caching`

Optional:

- `disabled` (Boolean) When set to true, disables the caching of SQL responses. (Default: false)
- `max_age` (Number) When present, specifies max duration for which items should persist in the cache. (Default: 60)
- `stale_while_revalidate` (Number) When present, indicates the number of seconds cache may serve the response after it becomes stale. (Default: 15)


<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `account_id` (String) Identifier


<a id="nestedatt--origin"></a>
### Nested Schema for `origin`

Optional:

- `access_client_id` (String) The Client ID of the Access token to use when connecting to the origin database
- `port` (Number) The port (default: 5432 for Postgres) of your origin database.

Read-Only:

- `database` (String) The name of your origin database.
- `host` (String) The host (hostname or IP) of your origin database.
- `scheme` (String) Specifies the URL scheme used to connect to your origin database.
- `user` (String) The user of your origin database.

