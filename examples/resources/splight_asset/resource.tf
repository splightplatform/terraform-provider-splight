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

resource "splight_asset" "my_asset" {
  name        = "My Asset"
  description = "My Asset Description"

  # You can use the existing tags in the platform
  # by name
  dynamic "tags" {
    for_each = length(data.splight_tags.my_tags.tags) > 0 ? [0] : []
    content {
      name = "ExistingTagName"
      id   = one([for t in data.splight_tags.my_tags.tags : t.id if t.name == "ExistingTagName"])
    }
  }

  # Or use the one created
  tags {
    name = splight_tag.my_tag.name
    id   = splight_tag.my_tag.id
  }

  # Choose the kind by name
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
