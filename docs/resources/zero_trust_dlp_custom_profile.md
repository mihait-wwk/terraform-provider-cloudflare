---
page_title: "cloudflare_zero_trust_dlp_custom_profile Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_dlp_custom_profile (Resource)




<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)
- `profiles` (Attributes List) (see [below for nested schema](#nestedatt--profiles))

### Optional

- `allowed_match_count` (Number)
- `context_awareness` (Attributes) Scan the context of predefined entries to only return matches surrounded by keywords. (see [below for nested schema](#nestedatt--context_awareness))
- `description` (String) The description of the profile
- `entries` (Attributes List) Custom entries from this profile (see [below for nested schema](#nestedatt--entries))
- `name` (String)
- `ocr_enabled` (Boolean)
- `profile_id` (String)
- `shared_entries` (Attributes List) Other entries, e.g. predefined or integration. (see [below for nested schema](#nestedatt--shared_entries))

<a id="nestedatt--profiles"></a>
### Nested Schema for `profiles`

Required:

- `entries` (Attributes List) (see [below for nested schema](#nestedatt--profiles--entries))
- `name` (String)

Optional:

- `allowed_match_count` (Number) Related DLP policies will trigger when the match count exceeds the number set.
- `context_awareness` (Attributes) Scan the context of predefined entries to only return matches surrounded by keywords. (see [below for nested schema](#nestedatt--profiles--context_awareness))
- `description` (String) The description of the profile
- `ocr_enabled` (Boolean)
- `shared_entries` (Attributes List) Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your Microsoft Information Protection profiles). (see [below for nested schema](#nestedatt--profiles--shared_entries))

<a id="nestedatt--profiles--entries"></a>
### Nested Schema for `profiles.entries`

Required:

- `enabled` (Boolean)
- `name` (String)

Optional:

- `pattern` (Attributes) (see [below for nested schema](#nestedatt--profiles--entries--pattern))
- `words` (List of String)

<a id="nestedatt--profiles--entries--pattern"></a>
### Nested Schema for `profiles.entries.pattern`

Required:

- `regex` (String)

Optional:

- `validation` (String)



<a id="nestedatt--profiles--context_awareness"></a>
### Nested Schema for `profiles.context_awareness`

Required:

- `enabled` (Boolean) If true, scan the context of predefined entries to only return matches surrounded by keywords.
- `skip` (Attributes) Content types to exclude from context analysis and return all matches. (see [below for nested schema](#nestedatt--profiles--context_awareness--skip))

<a id="nestedatt--profiles--context_awareness--skip"></a>
### Nested Schema for `profiles.context_awareness.skip`

Required:

- `files` (Boolean) If the content type is a file, skip context analysis and return all matches.



<a id="nestedatt--profiles--shared_entries"></a>
### Nested Schema for `profiles.shared_entries`

Required:

- `enabled` (Boolean)
- `entry_id` (String)
- `entry_type` (String)



<a id="nestedatt--context_awareness"></a>
### Nested Schema for `context_awareness`

Required:

- `enabled` (Boolean) If true, scan the context of predefined entries to only return matches surrounded by keywords.
- `skip` (Attributes) Content types to exclude from context analysis and return all matches. (see [below for nested schema](#nestedatt--context_awareness--skip))

<a id="nestedatt--context_awareness--skip"></a>
### Nested Schema for `context_awareness.skip`

Required:

- `files` (Boolean) If the content type is a file, skip context analysis and return all matches.



<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `enabled` (Boolean)
- `name` (String)
- `pattern` (Attributes) (see [below for nested schema](#nestedatt--entries--pattern))

Optional:

- `entry_id` (String)

<a id="nestedatt--entries--pattern"></a>
### Nested Schema for `entries.pattern`

Required:

- `regex` (String)

Optional:

- `validation` (String)



<a id="nestedatt--shared_entries"></a>
### Nested Schema for `shared_entries`

Required:

- `enabled` (Boolean)
- `entry_id` (String)
- `entry_type` (String)

