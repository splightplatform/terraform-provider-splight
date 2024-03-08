terraform {
  required_providers {
    spl = {
      source  = "splightplatform/splight"
      version = "0.1.4"
    }
  }
}

provider "spl" {
  hostname = var.spl_hostname
  token    = var.spl_api_token
}

# ASSETS
resource "spl_asset" "AssetMainTest" {
  name        = "AssetTF"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}
