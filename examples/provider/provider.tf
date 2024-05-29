terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "~> 0.1.0"
    }
  }
}

# If the provider configuration is not present, the provider will
# use the ones from the active Splight CLI workspace.
provider "splight" {
  hostname = "https://api.splight-ai.com"
  token    = "Splight <access_id> <secret_key>"
}
