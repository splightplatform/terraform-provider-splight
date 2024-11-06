---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_server Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_server (Resource)



## Example Usage

```terraform
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
    value       = jsonencode(splight_file.my_file.id)
  }

  env_vars {
    name  = "My Env Var"
    value = "My Env Var Value"
  }

  ports {
    name          = "My Port"
    protocol      = "My Protocol"
    internal_port = 8080
    exposed_port  = 8000
  }

  node                  = splight_node.my_node.id
  machine_instance_size = "very_large"
  log_level             = "error"
  restart_policy        = "Always"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) the name of the server to be created
- `version` (String) [NAME-VERSION] the version of the hub server

### Optional

- `config` (Block Set) static config parameters of the routine (see [below for nested schema](#nestedblock--config))
- `description` (String) optional description to add details of the resource
- `env_vars` (Block Set) environment variables for the server (see [below for nested schema](#nestedblock--env_vars))
- `log_level` (String) log level of the server
- `machine_instance_size` (String) instance size
- `node` (String) id of the compute node where the server runs
- `ports` (Block Set) ports of the server (see [below for nested schema](#nestedblock--ports))
- `restart_policy` (String) restart policy of the server
- `tags` (Block Set) tags of the resource (see [below for nested schema](#nestedblock--tags))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--config"></a>
### Nested Schema for `config`

Required:

- `name` (String)
- `type` (String)

Optional:

- `description` (String)
- `multiple` (Boolean)
- `required` (Boolean)
- `sensitive` (Boolean)
- `value` (String)


<a id="nestedblock--env_vars"></a>
### Nested Schema for `env_vars`

Required:

- `name` (String)
- `value` (String)


<a id="nestedblock--ports"></a>
### Nested Schema for `ports`

Required:

- `exposed_port` (Number)
- `internal_port` (Number)
- `name` (String)
- `protocol` (String)


<a id="nestedblock--tags"></a>
### Nested Schema for `tags`

Required:

- `id` (String) tag id
- `name` (String) tag name

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_server.<name> <server_id>
```