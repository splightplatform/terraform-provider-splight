terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "~> 0.1.0"
    }
  }
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

resource "splight_dashboard_actionlist_chart" "DashboardChartTest" {
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

  action_list_type  = "table"
  filter_name       = "some name"
  filter_asset_name = "some asset name"

  chart_items {
    ref_id           = "A"
    type             = "QUERY"
    color            = "red"
    expression_plain = ""
    query_filter_asset {
      id   = "1234-1234-1234-1234"
      name = "query filter asset name"
    }
    query_filter_attribute {
      id   = "1234-1234-1234-1234"
      name = "query filter attribute name"
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "1234-1234-1234-1234"
          attribute = "1234-1234-1234-1234"
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
      id   = "1234-1234-1234-1234"
      name = "query filter asset name"
    }
    query_filter_attribute {
      id   = "1234-1234-1234-1234"
      name = "query filter attribute name"
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "1234-1234-1234-1234"
          attribute = "1234-1234-1234-1234"
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