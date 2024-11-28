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

  tap_pos {
    value = jsonencode(0)
  }

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
