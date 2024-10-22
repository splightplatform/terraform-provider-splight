---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_generator Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_generator (Resource)



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

resource "splight_generator" "my_generator" {
  name        = "My Generator"
  description = "My Generator Description"

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

  co2_coefficient {
    value = jsonencode(1.1)
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `co2_coefficient` (Block Set, Min: 1, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--co2_coefficient))
- `name` (String) name of the resource

### Optional

- `description` (String) description of the resource
- `geometry` (String) geo position and shape of the resource
- `tags` (Block Set) tags of the resource (see [below for nested schema](#nestedblock--tags))

### Read-Only

- `active_power` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--active_power))
- `daily_emission_avoided` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--daily_emission_avoided))
- `daily_energy` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--daily_energy))
- `id` (String) The ID of this resource.
- `kind` (Set of Object) kind of the resource (see [below for nested schema](#nestedatt--kind))
- `reactive_power` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--reactive_power))

<a id="nestedblock--co2_coefficient"></a>
### Nested Schema for `co2_coefficient`

Required:

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


<a id="nestedatt--active_power"></a>
### Nested Schema for `active_power`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--daily_emission_avoided"></a>
### Nested Schema for `daily_emission_avoided`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--daily_energy"></a>
### Nested Schema for `daily_energy`

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


<a id="nestedatt--reactive_power"></a>
### Nested Schema for `reactive_power`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_generator.<name> <generator_id>
```