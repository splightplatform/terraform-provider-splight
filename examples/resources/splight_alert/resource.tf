resource "splight_alert" "AlertTest" {
  name        = "AlertTest"
  description = "Created with Terraform"
  type        = "rate"
  rate_unit   = "minute"
  rate_value  = 10
  time_window = 600

  thresholds {
    value       = 4.0
    status      = "no_alert"
    status_text = "CustomStatusText"
  }

  severity        = "sev8"
  operator        = "gt"
  aggregation     = "avg"
  target_variable = "A"

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
    # TODO: agregar qfasset y attr
  }

  alert_items {
    ref_id           = "B"
    type             = "QUERY"
    expression_plain = ""
    # TODO: agregar qfasset y attr
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
