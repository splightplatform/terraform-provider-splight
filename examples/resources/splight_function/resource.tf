resource "splight_function" "FunctionTest" {
  name            = "FunctionTest"
  description     = "Created with Terraform"
  type            = "rate"
  time_window     = 600
  rate_value      = 10
  rate_unit       = "minute"
  target_variable = "A"
  target_asset = {
    id   = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
    name = "An asset"
  }
  target_attribute = {
    id   = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
    name = "An attribute"
  }

  function_items {
    ref_id           = "A"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
          attribute = "c1d0d94b-5feb-4ebb-a527-0b0a34196252"
        }
      }
    ])
  }

  function_items {
    ref_id           = "B"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
          attribute = "c1d0d94b-5feb-4ebb-a527-0b0a34196252"
        }
      }
    ])
  }
}
