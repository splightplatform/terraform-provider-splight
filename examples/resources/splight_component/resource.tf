resource "splight_component" "ComponentTest" {
  name        = "ComponentTest"
  description = "Created with Terraform"
  version     = "Random-3.1.0"

  input {
    name        = "period"
    type        = "int"
    value       = 10
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "min"
    type        = "int"
    value       = 1
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max"
    type        = "int"
    value       = 150
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max_iterations"
    type        = "int"
    value       = 3
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "should_crash"
    type        = "bool"
    value       = "true"
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
}
