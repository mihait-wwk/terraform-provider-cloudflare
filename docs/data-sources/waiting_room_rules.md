---
page_title: "cloudflare_waiting_room_rules Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_waiting_room_rules (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (String) The action to take when the expression matches.
- `description` (String) The description of the rule.
- `enabled` (Boolean) When set to true, the rule is enabled.
- `expression` (String) Criteria defining when there is a match for the current rule.
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `id` (String) The ID of the rule.
- `last_updated` (String)
- `version` (String) The version of the rule.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `waiting_room_id` (String)
- `zone_id` (String) Identifier

