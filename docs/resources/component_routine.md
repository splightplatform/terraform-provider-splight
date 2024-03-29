---
page_title: "spl_component_routine Resource - spl"
subcategory: ""
description: |-
  
---

# spl_component_routine (Resource)



## Example Usage

```terraform
resource "spl_component_routine" "ComponentTestRoutine" {
  name         = "ComponentTestRoutine"
  description  = "Created with Terraform"
  type         = "IncomingRoutine"
  component_id = "1234-1234-1234-1234"

  config {
    name        = "config_param"
    type        = "bool"
    value       = "true"
    multiple    = false
    required    = true
    sensitive   = false
    description = "Created with Terraform123123"
  }

  output {
    name        = "address"
    description = "destination address for data to be pushed"
    type        = "DataAddress"
    value_type  = "Number"
    multiple    = false
    required    = true
    value = jsonencode({
      "asset" : "1234-1234-1234-1234",
      "attribute" : "1234-1234-1234-1234"
    })
  }
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `component_id` (String) reference to component to be attached
- `name` (String) name of the routine
- `type` (String) [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)

### Optional

- `config` (Block List) static config parameters of the routine (see [below for nested schema](#nestedblock--config))
- `description` (String) optional complementary information about the routine
- `input` (Block List) asset attribute where to read data. Only valid for OutgoingRoutine (see [below for nested schema](#nestedblock--input))
- `output` (Block List) asset attribute where to ingest data. Only valid for IncomingRoutine (see [below for nested schema](#nestedblock--output))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:

- `description` (String)
- `multiple` (Boolean)
- `name` (String)
- `required` (Boolean)
- `sensitive` (Boolean)
- `type` (String)
- `value` (String)


<a id="nestedblock--input"></a>
### Nested Schema for `input`

Required:

- `description` (String)
- `multiple` (Boolean)
- `name` (String)
- `required` (Boolean)
- `type` (String)
- `value` (String)
- `value_type` (String)


<a id="nestedblock--output"></a>
### Nested Schema for `output`

Required:

- `description` (String)
- `multiple` (Boolean)
- `name` (String)
- `required` (Boolean)
- `type` (String)
- `value` (String)
- `value_type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] spl_component_routine.<name> <component_routine_id>
```