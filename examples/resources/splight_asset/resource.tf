resource "splight_asset" "AssetMainTest" {
  name        = "AssetTF"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })

  kind {
    id   = "1234-1234-1234-1234"
    name = "Line"
  }
}
