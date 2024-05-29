resource "splight_dashboard_chart" "DashboardChartTest" {
  name          = "ChartTest"
  type          = "timeseries"
  tab           = "1234-1234-1234-1234"
  timestamp_gte = "now - 6h"
  timestamp_lte = "now"

  chart_items {
    ref_id           = "A"
    type             = "QUERY"
    color            = "red"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "1234-1234-1234-1234"
          attribute = "1234-1234-1234-1234"
        }
      }
    ])
  }

  chart_items {
    ref_id           = "B"
    color            = "blue"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "1234-1234-1234-1234"
          attribute = "1234-1234-1234-1234"
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
