terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Create a tag
resource "splight_tag" "my_tag" {
  name = "My Tag"
}

# Fetch tags
data "splight_tags" "my_tags" {}

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

resource "splight_dashboard" "my_dashboard" {
  name        = "My Dashboard"
  description = "My Dashboard Description"

  # Set related assets
  related_assets {
    id   = splight_asset.my_asset.id
    name = splight_asset.my_asset.name
  }

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
}

resource "splight_dashboard_tab" "my_dashboard_tab" {
  name      = "My Dashboard Tab"
  order     = 0
  dashboard = splight_dashboard.my_dashboard.id
}
