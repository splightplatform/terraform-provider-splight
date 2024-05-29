resource "splight_component" "ComponentTest" {
  name        = "ComponentTest"
  description = "Created with Terraform"
  version     = "Random-3.1.0"

  input {
    name        = "period"
    type        = "int"
    value       = jsonencode(10)
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "min"
    type        = "int"
    value       = jsonencode(1)
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max"
    type        = "int"
    value       = jsonencode(150)
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max_iterations"
    type        = "int"
    value       = jsonencode(3)
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "should_crash"
    type        = "bool"
    value       = jsonencode("true")
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
}
