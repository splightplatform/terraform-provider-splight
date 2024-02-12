resource "spl_alert" "AlertTest" {
  name            = "AlertTest"
  description     = "Created with Terraform"
  type            = "rate"
  time_window     = 600
  rate_value      = 10
  rate_unit       = "minute"
  target_variable = "A"
  operator        = "gt"
  aggregation     = "avg"
  severity        = "sev8"
  thresholds {
    value       = 4.0
    status      = "no_alert"
    status_text = "CustomStatusText"
  }
  alert_items {
    ref_id           = "A"
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
  alert_items {
    ref_id           = "B"
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
  related_assets = [
    "1234-1234-1234-1234"
  ]
}
