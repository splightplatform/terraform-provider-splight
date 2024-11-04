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

# Create a file for the Mosquitto configuration
resource "splight_file" "my_file" {
  path        = "./my_file"
  description = "My File"
}

# Create node for the server to run
resource "splight_node" "my_node" {
  name = "My Node"
  type = "splight_hosted"
}

resource "splight_server" "my_server" {
  name        = "My Server"
  description = "My Server Description"
  version     = "MQTTBroker-0.0.4"

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

  config {
    name        = "mosquitto_conf"
    type        = "File"
    description = "Mosquitto configuration file"
    value       = splight_file.my_file.id
  }

  node                  = splight_node.my_node.id
  machine_instance_size = "very_large"
  log_level             = "error"
  restart_policy        = "Always"
}
