terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_node" "my_node" {
  name = "My Node"
  type = "splight_hosted"
}
