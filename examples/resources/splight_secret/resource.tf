terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_secret" "my_secret" {
  name      = "My Secret"
  raw_value = "My Secret Value"
}
