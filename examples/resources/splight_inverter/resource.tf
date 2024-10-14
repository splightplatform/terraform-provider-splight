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

resource "splight_inverter" "my_inverter" {
  name        = "My Inverter"
  description = "My Inverter Description"

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

  make {
    value = jsonencode("Weiwei")
  }

  model {
    value = jsonencode("Model A")
  }

  serial_number {
    value = jsonencode(1233440042)
  }

  max_active_power {
    value = jsonencode(1000)
  }

  energy_measurement_type {
    value = jsonencode("daily")
  }
}
