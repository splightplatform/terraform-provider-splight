---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_transformer Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_transformer (Resource)



## Example Usage

```terraform
terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Create a tag
resource "splight_tag" "my_tag" {
  name = "My Tag"
}

# Fetch tags
data "splight_tags" "my_tags" {}

resource "splight_transformer" "my_transformer" {
  name        = "My Transformer"
  description = "My Transformer Description"
  timezone    = "America/Los_Angeles"

  # Use an existing tag in the platform
  dynamic "tags" {
    for_each = { for tag in data.splight_tags.my_tags.tags : tag.name => tag if tag.name == "Existing Tag" }

    content {
      name = tags.value.name
      id   = tags.value.id
    }
  }

  # Or use the one created
  tags {
    name = splight_tag.my_tag.name
    id   = splight_tag.my_tag.id
  }

  geometry = jsonencode({
    type = "GeometryCollection"
    geometries = [
      {
        type        = "Point"
        coordinates = [0, 0]
      }
    ]
  })

  # You may leave some metadata values unset, in order to use the defaults 
  tap_pos {}

  xn_ohm {
    value = jsonencode(0)
  }

  standard_type {
    value = jsonencode("")
  }

  capacitance {
    value = jsonencode(10.7)
  }

  conductance {
    value = jsonencode(0.001)
  }

  maximum_allowed_current {
    value = jsonencode(1.18)
  }

  maximum_allowed_power {
    value = jsonencode(450)
  }

  reactance {
    value = jsonencode(21.8)
  }

  resistance {
    value = jsonencode(0.21)
  }

  safety_margin_for_power {
    value = jsonencode(5)
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `capacitance` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--capacitance))
- `conductance` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--conductance))
- `maximum_allowed_current` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--maximum_allowed_current))
- `maximum_allowed_power` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--maximum_allowed_power))
- `name` (String) name of the resource
- `reactance` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--reactance))
- `resistance` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--resistance))
- `safety_margin_for_power` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--safety_margin_for_power))
- `standard_type` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--standard_type))
- `tap_pos` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--tap_pos))
- `xn_ohm` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--xn_ohm))

### Optional

- `description` (String) description of the resource
- `geometry` (String) geo position and shape of the resource
- `tags` (Block Set) tags of the resource (see [below for nested schema](#nestedblock--tags))
- `timezone` (String) timezone that overrides location-based timezone of the resource

### Read-Only

- `active_power_hv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--active_power_hv))
- `active_power_loss` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--active_power_loss))
- `active_power_lv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--active_power_lv))
- `contingency` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--contingency))
- `current_hv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--current_hv))
- `current_lv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--current_lv))
- `id` (String) The ID of this resource.
- `kind` (Set of Object) kind of the resource (see [below for nested schema](#nestedatt--kind))
- `reactive_power_hv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--reactive_power_hv))
- `reactive_power_loss` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--reactive_power_loss))
- `reactive_power_lv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--reactive_power_lv))
- `switch_status_hv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--switch_status_hv))
- `switch_status_lv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--switch_status_lv))
- `voltage_hv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--voltage_hv))
- `voltage_lv` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--voltage_lv))

<a id="nestedblock--capacitance"></a>
### Nested Schema for `capacitance`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--conductance"></a>
### Nested Schema for `conductance`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--maximum_allowed_current"></a>
### Nested Schema for `maximum_allowed_current`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--maximum_allowed_power"></a>
### Nested Schema for `maximum_allowed_power`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--reactance"></a>
### Nested Schema for `reactance`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--resistance"></a>
### Nested Schema for `resistance`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--safety_margin_for_power"></a>
### Nested Schema for `safety_margin_for_power`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--standard_type"></a>
### Nested Schema for `standard_type`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--tap_pos"></a>
### Nested Schema for `tap_pos`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--xn_ohm"></a>
### Nested Schema for `xn_ohm`

Optional:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Required:

- `id` (String) tag id
- `name` (String) tag name


<a id="nestedatt--active_power_hv"></a>
### Nested Schema for `active_power_hv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--active_power_loss"></a>
### Nested Schema for `active_power_loss`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--active_power_lv"></a>
### Nested Schema for `active_power_lv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--contingency"></a>
### Nested Schema for `contingency`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--current_hv"></a>
### Nested Schema for `current_hv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--current_lv"></a>
### Nested Schema for `current_lv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--kind"></a>
### Nested Schema for `kind`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--reactive_power_hv"></a>
### Nested Schema for `reactive_power_hv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--reactive_power_loss"></a>
### Nested Schema for `reactive_power_loss`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--reactive_power_lv"></a>
### Nested Schema for `reactive_power_lv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--switch_status_hv"></a>
### Nested Schema for `switch_status_hv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--switch_status_lv"></a>
### Nested Schema for `switch_status_lv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--voltage_hv"></a>
### Nested Schema for `voltage_hv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--voltage_lv"></a>
### Nested Schema for `voltage_lv`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_transformer.<name> <transformer_id>
```