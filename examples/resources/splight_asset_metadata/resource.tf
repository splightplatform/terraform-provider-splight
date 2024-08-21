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

resource "splight_asset_metadata" "my_asset_metadata" {
  name  = "My Asset Metadata"
  type  = "Number"
  unit  = "meters"
  value = jsonencode(10)
  asset = splight_asset.my_asset.id
}
