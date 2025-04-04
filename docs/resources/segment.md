---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_segment Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_segment (Resource)



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

resource "splight_segment" "my_segment" {
  name        = "My Segment"
  description = "My Segment Description"

  # This is overridden by the GeoJSON location
  # and will show perma diff if both are set
  timezone = "America/Los_Angeles"

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

  # You may ommit some keys to use the default values from the API
  altitude {
    value = jsonencode(400)
  }

  azimuth {
    value = jsonencode(1)
  }

  cumulative_distance {
    value = jsonencode(1)
  }

  reference_sag {
    value = jsonencode(1)
  }

  reference_temperature {
    value = jsonencode(1)
  }

  span_length {
    value = jsonencode(1)
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) name of the resource

### Optional

- `altitude` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--altitude))
- `azimuth` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--azimuth))
- `cumulative_distance` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--cumulative_distance))
- `description` (String) description of the resource
- `geometry` (String) geo position and shape of the resource
- `reference_sag` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--reference_sag))
- `reference_temperature` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--reference_temperature))
- `span_length` (Block Set, Max: 1) attribute of the resource (see [below for nested schema](#nestedblock--span_length))
- `tags` (Block Set) tags of the resource (see [below for nested schema](#nestedblock--tags))
- `timezone` (String) timezone that overrides location-based timezone of the resource

### Read-Only

- `id` (String) The ID of this resource.
- `kind` (Set of Object) kind of the resource (see [below for nested schema](#nestedatt--kind))
- `temperature` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--temperature))
- `wind_direction` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--wind_direction))
- `wind_speed` (Set of Object) attribute of the resource (see [below for nested schema](#nestedatt--wind_speed))

<a id="nestedblock--altitude"></a>
### Nested Schema for `altitude`

Required:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--azimuth"></a>
### Nested Schema for `azimuth`

Required:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--cumulative_distance"></a>
### Nested Schema for `cumulative_distance`

Required:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--reference_sag"></a>
### Nested Schema for `reference_sag`

Required:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--reference_temperature"></a>
### Nested Schema for `reference_temperature`

Required:

- `value` (String) metadata value

Read-Only:

- `asset` (String) reference to the asset to be linked to
- `id` (String) id of the resource
- `name` (String) name of the resource
- `type` (String) [String|Boolean|Number] type of the data to be ingested in this attribute
- `unit` (String) unit of measure


<a id="nestedblock--span_length"></a>
### Nested Schema for `span_length`

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


<a id="nestedatt--kind"></a>
### Nested Schema for `kind`

Read-Only:

- `id` (String)
- `name` (String)


<a id="nestedatt--temperature"></a>
### Nested Schema for `temperature`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--wind_direction"></a>
### Nested Schema for `wind_direction`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)


<a id="nestedatt--wind_speed"></a>
### Nested Schema for `wind_speed`

Read-Only:

- `asset` (String)
- `id` (String)
- `name` (String)
- `type` (String)
- `unit` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_segment.<name> <segment_id>
```
