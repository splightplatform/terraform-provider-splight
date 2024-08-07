---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_file Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_file (Resource)



## Example Usage

```terraform
resource "splight_file" "FileInnerTest" {
  path        = "./variables.tf"
  description = "Sample file for testing inner file"
  parent      = "1234-1234-1234-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `path` (String) the path for the file resource in your system

### Optional

- `description` (String) complementary information to describe the file
- `parent` (String) the id reference for a folder to be placed in
- `related_assets` (Set of String) assets to be linked

### Read-Only

- `checksum` (String)
- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_file.<name> <file_id>
```
