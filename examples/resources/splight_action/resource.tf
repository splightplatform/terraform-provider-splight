terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_asset" "my_asset" {
  name        = "My Asset"
  description = "My Asset Description"
  timezone    = "America/Los_Angeles"

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

resource "splight_asset_attribute" "my_attribute" {
  name  = "My Attribute"
  type  = "Number"
  unit  = "meters"
  asset = splight_asset.my_asset.id
}

resource "splight_action" "my_action" {
  name = "My Action"

  asset {
    id   = splight_asset.my_asset.id
    name = splight_asset.my_asset.name
  }

  setpoints {
    value = jsonencode(1) # Should match the type of the attribute
    attribute {
      id   = splight_asset_attribute.my_attribute.id
      name = splight_asset_attribute.my_attribute.name
    }
  }
}
