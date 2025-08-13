terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Fetch kinds
data "splight_asset_kinds" "my_kinds" {}

# Create a Segment
resource "splight_asset" "my_segment" {
  name = "My Segment"

  # This overrides the timezone computed from the geolocation
  custom_timezone = "America/Los_Angeles"

  kind {
    name = "Segment"
    id   = one([for k in data.splight_asset_kinds.my_kinds.kinds : k.id if k.name == "Segment"])
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

}

# Create a Line
resource "splight_asset" "my_line" {
  name = "My Line"

  # This overrides the timezone computed from the geolocation
  custom_timezone = "America/Los_Angeles"

  kind {
    name = "Line"
    id   = one([for k in data.splight_asset_kinds.my_kinds.kinds : k.id if k.name == "Line"])
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
}

resource "splight_asset_relation" "my_relation" {
  name        = "My Relation"
  description = "My Relation Description"

  related_asset_kind {
    name = "Segment"
    id   = one([for k in data.splight_asset_kinds.my_kinds.kinds : k.id if k.name == "Segment"])
  }

  related_asset {
    id   = splight_asset.my_segment.id
    name = splight_asset.my_segment.name
  }

  asset {
    id   = splight_asset.my_line.id
    name = splight_asset.my_line.name
  }
}
