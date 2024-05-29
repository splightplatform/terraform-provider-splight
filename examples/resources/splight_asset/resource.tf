resource "splight_asset" "AssetMainTest" {
  name        = "AssetTF"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}
