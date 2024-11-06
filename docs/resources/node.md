---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_node Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_node (Resource)



## Example Usage

```terraform
terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_node" "my_node" {
  name = "My Node"
  type = "splight_hosted"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) name of the resource
- `type` (String) either splight_hosted or self_hosted

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_node.<name> <node_id>
```