terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Create a tag
resource "splight_tag" "my_tag" {
  name = "My Tag"
}

# Fetch tags
data "splight_tags" "my_tags" {}

# Create node for the algorithm to run
resource "splight_node" "my_node" {
  name = "My Node"
  type = "splight_hosted"
}

resource "splight_algorithm" "my_algorithm" {
  name        = "My Algorithm"
  description = "My Algorithm Description"
  version     = "SplightForecaster-2.1.1"

  # Use an existing tag in the platform
  dynamic "tags" {
    for_each = { for tag in data.splight_tags.my_tags.tags : tag.name => tag if tag.name == "Existing Tag" }

    content {
      name = tags.value.name
      id   = tags.value.id
    }
  }

  # Or use the one created
  tags {
    name = splight_tag.my_tag.name
    id   = splight_tag.my_tag.id
  }

  input {
    name        = "model"
    description = "Model to be used for making predictions"
    type        = "str"
    required    = false
    value       = jsonencode("XGB")
  }

  input {
    name        = "future_preds"
    description = "Number of predictions to make"
    type        = "int"
    required    = false
    value       = jsonencode(168)
  }

  input {
    name        = "frequency"
    description = "The data frequency"
    type        = "str"
    required    = false
    value       = jsonencode("H")
  }

  input {
    name        = "interpolation_method"
    description = "How we will interpolate missing data points."
    type        = "str"
    required    = false
    value       = jsonencode("linear")
  }

  input {
    name        = "fetched_data_grouping"
    description = "How we will group the fetched data."
    type        = "str"
    required    = false
    value       = jsonencode("last")
  }

  input {
    name        = "percentile"
    description = "Percentile value if you want to use a non-deterministic model"
    type        = "float"
    required    = false
    value       = jsonencode("")
  }

  node                  = splight_node.my_node.id
  machine_instance_size = "very_large"
  log_level             = "error"
  restart_policy        = "Always"
}
