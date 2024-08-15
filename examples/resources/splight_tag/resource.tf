terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_tag" "my_tag" {
  name = "My Tag"
}
