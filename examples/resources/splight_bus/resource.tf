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

resource "splight_grid" "my_grid" {
  name        = "My Grid"
  description = "My Grid Description"
}

# Fetch tags
data "splight_tags" "my_tags" {}

resource "splight_bus" "my_bus" {
  name        = "My Bus"
  description = "My Bus Description"

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

  # Set the relationships
  grid = splight_grid.my_grid.id

  # You may ommit some keys to use the default values from the API
  nominal_voltage_kv {
    value = jsonencode(2.2)
  }
}
