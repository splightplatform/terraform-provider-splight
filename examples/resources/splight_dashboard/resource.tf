terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_dashboard" "DashboardTest" {
  name = "DashboardTest"
  // TODO: fill these
  assets = []
  tags   = []
}
