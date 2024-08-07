---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_secret Resource - terraform-provider-splight"
subcategory: ""
description: |-
  Provides a Cloudflare Observatory Scheduled Test resource.
---

# splight_secret (Resource)

Provides a Cloudflare Observatory Scheduled Test resource.

## Example Usage

```terraform
resource "splight_secret" "SecretTest" {
  name      = "SecretTest"
  raw_value = "ASUPERSECR3T"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `raw_value` (String, Sensitive)

### Read-Only

- `id` (String) The ID of this resource.
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_secret.<name> <secret_id>
```
