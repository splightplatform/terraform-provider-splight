---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_dashboard Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_dashboard (Resource)



## Example Usage

```terraform
resource "splight_dashboard" "DashboardTest" {
  name = "DashboardTest"
  related_assets = [
    "1234-1234-1234-1234"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) dashboard name

### Optional

- `description` (String) complementary information for the dashboard
- `related_assets` (Set of String) assets linked

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_dashboard.<name> <dashboard_id>
```
