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

resource "splight_asset_attribute" "my_attribute" {
  name  = "My Attribute"
  type  = "Number"
  asset = splight_asset.my_asset.id
}

resource "splight_asset" "my_target_asset" {
  name        = "My Target Asset"
  description = "My Target Asset Description"
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

resource "splight_asset_attribute" "my_target_attribute" {
  name  = "My Target Attribute"
  type  = "Number"
  asset = splight_asset.my_target_asset.id
}

resource "splight_function" "FunctionTest" {
  name            = "My Function"
  description     = "My Function Description"
  type            = "rate"
  rate_unit       = "minute"
  rate_value      = 10
  time_window     = 3600 * 12
  target_variable = "B"

  target_asset {
    id   = splight_asset.my_target_asset.id
    name = splight_asset.my_target_asset.name
  }

  target_attribute {
    id   = splight_asset_attribute.my_target_attribute.id
    name = splight_asset_attribute.my_target_attribute.name
  }

  function_items {
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

    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.my_asset.id
          attribute = splight_asset_attribute.my_attribute.id
        }
      }
    ])

  }

  function_items {
    ref_id     = "B"
    type       = "EXPRESSION"
    expression = "A * 2"
    expression_plain = jsonencode({
      "$function" : {
        "body" : "function () { return A * 2 }",
        "args" : [],
        "lang" : "js"
      }
    })

    query_filter_asset {}

    query_filter_attribute {}

    query_plain = ""

  }
}
