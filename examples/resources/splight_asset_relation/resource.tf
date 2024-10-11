terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Fetch kinds
data "splight_asset_kinds" "my_kinds" {}

# Create a tag
resource "splight_tag" "my_tag" {
  name = "My Tag"
}

# Fetch tags
data "splight_tags" "my_tags" {}

# Create a Segment
resource "splight_segment" "my_segment" {
  name        = "My Segment"
  description = "My Segment Description"

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

  altitude {
    value = jsonencode(1.1)
  }

  azimuth {
    value = jsonencode(1.1)
  }

  cumulative_distance {
    value = jsonencode(1.1)
  }
}

# Create a Line
resource "splight_line" "my_line" {
  name        = "My Line"
  description = "My Line Description"

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

  diameter {
    value = jsonencode(1.1)
  }

  absorptivity {
    value = jsonencode(1.1)
  }

  atmosphere {
    value = jsonencode("clean")
  }

  capacitance {
    value = jsonencode(1.1)
  }

  conductance {
    value = jsonencode(1.1)
  }

  emissivity {
    value = jsonencode(1.1)
  }

  length {
    value = jsonencode(1.1)
  }

  maximum_allowed_current {
    value = jsonencode(1.1)
  }

  maximum_allowed_power {
    value = jsonencode(1.1)
  }

  maximum_allowed_temperature {
    value = jsonencode(1.1)
  }

  maximum_allowed_temperature_lte {
    value = jsonencode(1.1)
  }

  maximum_allowed_temperature_ste {
    value = jsonencode(1.1)
  }

  number_of_conductors {
    value = jsonencode(1.1)
  }

  reactance {
    value = jsonencode(1.1)
  }

  reference_resistance {
    value = jsonencode(1.1)
  }

  resistance {
    value = jsonencode(1.1)
  }

  safety_margin_for_power {
    value = jsonencode(1.1)
  }

  susceptance {
    value = jsonencode(1.1)
  }

  temperature_coeff_resistance {
    value = jsonencode(1.1)
  }
}

resource "splight_asset_relation" "my_relation" {
  name        = "My Relation"
  description = "My Relation Description"

  related_asset_kind {
    name = "Segment"
    id   = one([for k in data.splight_asset_kinds.my_kinds.kinds : k.id if k.name == "Segment"])
  }

  related_asset {
    id   = splight_segment.my_segment.id
    name = splight_segment.my_segment.name
  }

  asset {
    id   = splight_line.my_line.id
    name = splight_line.my_line.name
  }
}
