resource "splight_asset_metadata" "AssetTestMetadata" {
  name  = "Key"
  type  = "Number"
  unit  = "meters"
  value = jsonencode(10)
  asset = "1234-1234-1234-1234"
}
