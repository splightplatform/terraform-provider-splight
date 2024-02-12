resource "spl_component_routine" "ComponentTestRoutine" {
  name         = "ComponentTestRoutine"
  description  = "Created with Terraform"
  type         = "IncomingRoutine"
  component_id = "1234-1234-1234-1234"

  config {
    name        = "config_param"
    type        = "bool"
    value       = "true"
    multiple    = false
    required    = true
    sensitive   = false
    description = "Created with Terraform123123"
  }

  output {
    name        = "address"
    description = "destination address for data to be pushed"
    type        = "DataAddress"
    value_type  = "Number"
    multiple    = false
    required    = true
    value = jsonencode({
      "asset" : "1234-1234-1234-1234",
      "attribute" : "1234-1234-1234-1234"
    })
  }
}
