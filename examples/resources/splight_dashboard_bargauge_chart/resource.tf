terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "~> 0.1.0"
    }
  }
}

resource "splight_asset" "AssetTest" {
  name        = "AssetTest"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}

resource "splight_asset_attribute" "AttributeTest1" {
  name  = "Attribute1"
  type  = "Number"
  unit  = "meters"
  asset = splight_asset.AssetTest.id
}

resource "splight_asset_attribute" "AttributeTest2" {
  name  = "Attribute2"
  type  = "Number"
  unit  = "seconds"
  asset = splight_asset.AssetTest.id
}

resource "splight_dashboard" "DashboardTest_fz" {
  name           = "DashboardTest"
  related_assets = []
}

resource "splight_dashboard_tab" "DashboardTabTest_fz" {
  name      = "TabTest"
  order     = 0
  dashboard = splight_dashboard.DashboardTest.id
}

resource "splight_dashboard_bargauge_chart" "DashboardChartTest" {
  name               = "ChartTest"
  tab                = splight_dashboard_tab.DashboardTabTest.id
  timestamp_gte      = "now - 7d"
  timestamp_lte      = "now"
  description        = "Chart description"
  min_height         = 1
  min_width          = 4
  display_time_range = true
  labels_display     = true
  labels_aggregation = "last"
  labels_placement   = "bottom"
  show_beyond_data   = true
  height             = 10
  width              = 20
  collection         = "default"

  max_limit          = 100
  number_of_decimals = 2
  orientation        = "vertical"

  chart_items {
    ref_id           = "A"
    type             = "QUERY"
    color            = "red"
    expression_plain = ""
    query_filter_asset {
      id   = splight_asset.AssetTest.id
      name = splight_asset.AssetTest.name
    }
    query_filter_attribute {
      id   = splight_asset.AttributeTest1.id
      name = splight_asset.AttributeTest1.name
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest.id
          attribute = splight_asset.AttributeTest1.id
        }
      },
      {
        "$addFields" : {
          "timestamp" : {
            "$dateTrunc" : {
              "date" : "$timestamp",
              "unit" : "day",
              "binSize" : 1
            }
          }
        }
      },
      {
        "$group" : {
          "_id" : "$timestamp",
          "value" : {
            "$last" : "$value"
          },
          "timestamp" : {
            "$last" : "$timestamp"
          }
        }
      }
    ])
  }

  chart_items {
    ref_id           = "B"
    color            = "blue"
    type             = "QUERY"
    expression_plain = ""
    query_filter_asset {
      id   = splight_asset.AssetTest.id
      name = splight_asset.AssetTest.name
    }
    query_filter_attribute {
      id   = splight_asset.AttributeTest2.id
      name = splight_asset.AttributeTest2.name
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest.id
          attribute = splight_asset.AttributeTest2.id
        }
      },
      {
        "$addFields" : {
          "timestamp" : {
            "$dateTrunc" : {
              "date" : "$timestamp",
              "unit" : "hour",
              "binSize" : 1
            }
          }
        }
      },
      {
        "$group" : {
          "_id" : "$timestamp",
          "value" : {
            "$last" : "$value"
          },
          "timestamp" : {
            "$last" : "$timestamp"
          }
        }
      }
    ])
  }

  thresholds {
    color        = "#00edcf"
    display_text = "T1Test"
    value        = 13.1
  }

  value_mappings {
    display_text = "MODIFICADO"
    match_value  = "123.3"
    type         = "exact_match"
    order        = 0
  }
}