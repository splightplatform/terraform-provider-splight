terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

# Create an asset and an attribute to write the output
resource "splight_asset" "my_asset" {
  name        = "My Asset"
  description = "My Asset Description"
  geometry = jsonencode({
    type = "GeometryCollection"
    geometries = [
      {
        type        = "Point"
        coordinates = [0, 0]
      }
    ]
  })
}

resource "splight_asset_attribute" "my_attribute" {
  name  = "My Attribute"
  type  = "Number"
  unit  = "kWh"
  asset = splight_asset.my_asset.id
}

# Create a tag
resource "splight_tag" "my_tag" {
  name = "My Tag"
}

# Fetch tags
data "splight_tags" "my_tags" {}

resource "splight_component" "my_component" {
  name        = "My Component"
  description = "My Component Description"
  version     = "MQTT-6.5.5"

  # Use an existing tag if it exists in the platform by name
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
    name        = "broker_host"
    type        = "str"
    description = "Address of the broker"
    value       = jsonencode("12.197.221.32")
  }

  input {
    name        = "broker_port"
    type        = "int"
    description = "Port of the broker"
    value       = jsonencode(8883)
  }

  input {
    name        = "broker_username"
    type        = "str"
    required    = false
    description = "The username used to authenticate the MQTT client"
    value       = jsonencode("my-user")
  }

  input {
    name        = "broker_password"
    type        = "str"
    required    = false
    description = "The password used to authenticate the MQTT client"
    value       = jsonencode("my-password")
  }

  input {
    name        = "topics"
    type        = "str"
    multiple    = true
    description = "Subscribe to specific topics for better performance; wildcards allowed."
    value       = jsonencode(["#"])
  }

  input {
    name        = "proxy"
    type        = "str"
    required    = false
    description = "The proxy address used to connect to the client"
  }

  input {
    name        = "use_ssl"
    type        = "str"
    description = "Whether to establish a TLS connection or not"
    value       = jsonencode("no")
  }

  input {
    name        = "buffer_timeout"
    type        = "int"
    description = "Time in seconds before buffered values are sent"
    value       = jsonencode(45)
  }

  input {
    name        = "buffer_size"
    type        = "int"
    description = "Maximum number of values to accumulate before sending"
    value       = jsonencode(1000)
  }

  input {
    name        = "write_routines_interval"
    description = "Check for values every N seconds within [now - N now]. Only the latest one is considered."
    type        = "int"
    value       = jsonencode(300)
  }
}

# Create a routine for the component
resource "splight_component_routine" "my_routine" {
  name         = "My Routine"
  description  = "My Routine Description"
  type         = "MQTTReadNumber"
  component_id = splight_component.my_component.id

  config {
    name        = "path"
    description = "JSON path used to parse the message. Only the first match is saved"
    multiple    = false
    required    = true
    sensitive   = false
    type        = "String"
    value       = jsonencode("[*].my.value.key")
  }

  config {
    name        = "topic"
    description = "The routine only processes messages from topic. Wildcards are not allowed"
    multiple    = false
    required    = true
    sensitive   = false
    type        = "String"
    value       = jsonencode("my/topic/path")
  }

  output {
    name        = "value"
    description = "The received values are saved to this AssetAttribute"
    multiple    = false
    required    = true
    sensitive   = false
    type        = "DataAddress"
    value_type  = "Number"

    value {
      asset     = splight_asset.my_asset.id
      attribute = splight_asset_attribute.my_attribute.id
    }
  }
}
