terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "0.1.11"
    }
  }
}

provider "splight" {
  hostname = var.hostname
  token    = var.api_token
}

resource "splight_asset" "MyAssetResource" {
  name        = "MyAsset"
  description = "Created with Terraform"
  geometry = jsonencode({
    type = "GeometryCollection"
    geometries = [
      {
        type        = "Point"
        coordinates = [-62.040, -35.706]
      }
    ]
  })
}
