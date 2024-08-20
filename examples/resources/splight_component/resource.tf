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

resource "splight_component" "MQTTConnector" {
  name        = "My Component"
  description = "My Component Description"
  version     = "MQTT-6.5.5"

  # You can set all the tags in the platform
  dynamic "tags" {
    for_each = { for idx, tag in data.splight_tags.my_tags.tags : idx => tag }

    content {
      id   = tags.value.id
      name = tags.value.name
    }
  }

  # You can use the existing tags in the platform
  # by name
  dynamic "tags" {
    for_each = length(data.splight_tags.my_tags.tags) > 0 ? [0] : []
    content {
      name = "ExistingTagName"
      id   = one([for t in data.splight_tags.my_tags.tags : t.id if t.name == "ExistingTagName"])
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
