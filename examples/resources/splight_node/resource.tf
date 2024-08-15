terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_node" "my_node" {
  name          = "My Node"
  instance_type = "t2.micro"
  region        = "us-east-1"
}
