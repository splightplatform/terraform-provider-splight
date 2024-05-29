terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "~> 0.1.0"
    }
  }
}

# If the provider configuration is not present, the provider will
# use the active Splight CLI workspace.
provider "splight" {
  hostname = var.hostname
  token    = var.api_token
}
