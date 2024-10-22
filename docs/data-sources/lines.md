---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_lines Data Source - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_lines (Data Source)



## Example Usage

```terraform
data "splight_lines" "lines" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `tags` (List of Object) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `id` (String)
- `name` (String)