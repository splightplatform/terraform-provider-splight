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
  name            = "My Asset"
  description     = "My Asset Description"
  custom_timezone = "America/Los_Angeles"

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
  asset = splight_asset.my_asset.id
}

resource "splight_alert" "my_alert" {
  name        = "My Alert"
  description = "My Alert Description"
  type        = "rate"
  rate_unit   = "minute"
  rate_value  = 10
  time_window = 3600

  thresholds {
    value       = 1
    status      = "alert"
    status_text = "Some warning!"
  }

  severity        = "sev1"
  operator        = "lt"
  aggregation     = "max"
  target_variable = "A"

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

  alert_items {
    ref_id           = "A"
    type             = "QUERY"
    expression       = ""
    expression_plain = ""

    query_filter_asset {
      id   = splight_asset.my_asset.id
      name = splight_asset.my_asset.name
    }

    query_filter_attribute {
      id   = splight_asset_attribute.my_attribute.id
      name = splight_asset_attribute.my_attribute.name
    }

    query_group_function = "avg"
    query_group_unit     = "day"

    query_plain = jsonencode(
      [
        {
          "$match" : {
            "asset" : splight_asset.my_asset.id,
            "attribute" : splight_asset_attribute.my_attribute.id,
          }
        }
      ]
    )
  }

  related_assets {
    id   = splight_asset.my_asset.id
    name = splight_asset.my_asset.name
  }
}
