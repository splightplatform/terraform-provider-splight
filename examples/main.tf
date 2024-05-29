terraform {
  required_providers {
    splight = {
      source  = "splightplatform/splight"
      version = "~> 0.1.0"
    }
  }
}

locals {
  assets = {
    AssetTF1 : {
      geometries = [
        {
          type        = "Point"
          coordinates = [-62.040, -35.706]
        }
      ]
    },
    AssetTF2 : {
      geometries = [
        {
          type        = "Point"
          coordinates = [-62.040, -36.706]
        }
      ]
    },
    AssetTF3 : {
      geometries = [
        {
          type        = "Point"
          coordinates = [-62.040, -37.706]
        }
      ]
    }
  }
  attributes = {
    AttributeTF1 : { type : "Number" },
    AttributeTF2 : { type : "Number" }
  }
  metadata = {
    metadataTF1 = { type : "String", value : "METAINFO1" }
    metadataTF2 = { type : "String", value : "METAINFO2" }
    metadataTF3 = { type : "String", value : "METAINFO3" }
  }
  asset_attribute_combinations = merge([
    for assetKey, assetValue in local.assets : {
      for attributeKey, attributeValue in local.attributes : "${assetKey}/${attributeKey}" => {
        asset        = assetValue
        assetKey     = assetKey
        attribute    = attributeValue
        attributeKey = attributeKey
      }
    }
  ]...)
  asset_metadata_combinations = merge([
    for assetKey, assetValue in local.assets : {
      for metadataKey, metadataValue in local.metadata : "${assetKey}/${metadataKey}" => {
        asset       = assetValue
        assetKey    = assetKey
        metadata    = metadataValue
        metadataKey = metadataKey
      }
    }
  ]...)
}

# If the provider configuration is not present, the provider will
# use the active Splight CLI workspace.
provider "splight" {
  hostname = var.hostname
  token    = var.api_token
}

# ASSETS
resource "splight_asset" "AssetMainTest" {
  name        = "AssetTF"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}
resource "splight_asset" "AssetTest" {
  for_each    = local.assets
  name        = each.key
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = each.value.geometries
  })
  related_assets = [
    splight_asset.AssetMainTest.id
  ]
}

resource "splight_asset_attribute" "AssetTestAttribute" {
  for_each = local.asset_attribute_combinations
  name     = each.value.attributeKey
  type     = each.value.attribute.type

  asset = splight_asset.AssetTest[each.value.assetKey].id
}

resource "splight_asset_attribute" "AssetTestFunctionAttribute" {
  for_each = local.assets
  name     = "FunctionAttribute"
  type     = "Number"
  asset    = splight_asset.AssetTest[each.key].id
}

resource "splight_asset_metadata" "AssetTestMetadata" {
  for_each = local.asset_metadata_combinations
  name     = each.value.metadataKey
  value    = each.value.metadata.value
  type     = each.value.metadata.type
  asset    = splight_asset.AssetTest[each.value.assetKey].id
}

# COMPONENTS
resource "splight_component" "ComponentTest" {
  name        = "ComponentTest"
  description = "Created with Terraform"
  version     = "Random-3.1.0"

  input {
    name        = "period"
    type        = "int"
    value       = 10
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "min"
    type        = "int"
    value       = 1
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max"
    type        = "int"
    value       = 150
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "max_iterations"
    type        = "int"
    value       = 3
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
  input {
    name        = "should_crash"
    type        = "bool"
    value       = "true"
    multiple    = false
    required    = false
    sensitive   = false
    description = ""
  }
}

resource "splight_component_routine" "ComponentTestRoutine" {
  for_each = local.asset_attribute_combinations

  name         = "ComponentTestRoutine-${each.key}"
  description  = "Created with Terraform"
  type         = "IncomingRoutine"
  component_id = splight_component.ComponentTest.id

  config {
    name        = "config_param"
    type        = "bool"
    value       = "true"
    multiple    = false
    required    = true
    sensitive   = false
    description = "Created with Terraform123123"
  }

  output {
    name        = "address"
    description = "destination address for data to be pushed"
    type        = "DataAddress"
    value_type  = "Number"
    multiple    = false
    required    = true
    value = jsonencode({
      "asset" : splight_asset.AssetTest[each.value.assetKey].id,
      "attribute" : splight_asset_attribute.AssetTestAttribute[each.key].id
    })
  }
}

# FUNCTIONS AND ALERTS
resource "splight_function" "FunctionTest" {
  for_each        = local.assets
  name            = "FunctionTest-${each.key}"
  description     = "Created with Terraform"
  type            = "rate"
  time_window     = 600
  rate_value      = 10
  rate_unit       = "minute"
  target_variable = "A"
  target_asset = {
    id   = splight_asset.AssetTest[each.key].id
    name = splight_asset.AssetTest[each.key].name
  }
  target_attribute = {
    id   = splight_asset_attribute.AssetTestFunctionAttribute[each.key].id
    name = splight_asset_attribute.AssetTestFunctionAttribute[each.key].name
  }
  function_items {
    ref_id           = "A"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
          attribute = "c1d0d94b-5feb-4ebb-a527-0b0a34196252"
        }
      }
    ])
  }
  function_items {
    ref_id           = "B"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = "49551a15-d79b-40dc-9434-1b33d6b2fcb2"
          attribute = "c1d0d94b-5feb-4ebb-a527-0b0a34196252"
        }
      }
    ])
  }
}

resource "splight_alert" "AlertTest" {
  for_each        = splight_asset.AssetTest
  name            = "AlertTest-${each.key}"
  description     = "Created with Terraform"
  type            = "rate"
  time_window     = 600
  rate_value      = 10
  rate_unit       = "minute"
  target_variable = "A"
  operator        = "gt"
  aggregation     = "avg"
  severity        = "sev8"
  thresholds {
    value       = 4.0
    status      = "no_alert"
    status_text = "CustomStatusText"
  }
  alert_items {
    ref_id           = "A"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest[each.key].id
          attribute = splight_asset_attribute.AssetTestFunctionAttribute[each.key].id
        }
      }
    ])
  }
  alert_items {
    ref_id           = "B"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest[each.key].id
          attribute = splight_asset_attribute.AssetTestFunctionAttribute[each.key].id
        }
      }
    ])
  }
  related_assets = [
    splight_asset.AssetMainTest.id
  ]
}

# DASHBOARDS
resource "splight_dashboard" "DashboardTest" {
  for_each = local.assets
  name     = "DashboardTest-${each.key}"
  related_assets = [
    splight_asset.AssetMainTest.id
  ]
}

resource "splight_dashboard_tab" "DashboardTabTest" {
  for_each  = local.assets
  name      = "TabTest"
  order     = 0
  dashboard = splight_dashboard.DashboardTest[each.key].id
}

resource "splight_dashboard_chart" "DashboardChartTest" {
  for_each      = splight_dashboard_tab.DashboardTabTest
  name          = "ChartTest"
  type          = "timeseries"
  tab           = each.value.id
  timestamp_gte = "now - 6h"
  timestamp_lte = "now"
  chart_items {
    ref_id           = "A"
    type             = "QUERY"
    color            = "red"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest[each.key].id
          attribute = splight_asset_attribute.AssetTestAttribute["${each.key}/AttributeTF1"].id
        }
      }
    ])
  }
  chart_items {
    ref_id           = "B"
    color            = "blue"
    type             = "QUERY"
    expression_plain = ""
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest[each.key].id
          attribute = splight_asset_attribute.AssetTestAttribute["${each.key}/AttributeTF2"].id
        }
      }
    ])
  }
  thresholds {
    color        = "#00edcf"
    display_text = "T1Test"
    value        = 13.1
  }
  value_mappings {
    display_text = "MODIFICADO"
    match_value  = "123.3"
    type         = "exact_match"
    order        = 0
  }
}

# FILES
resource "splight_file" "FileTest" {
  file        = "./main.tf"
  description = "Sample file for testing"
  related_assets = [
    splight_asset.AssetMainTest.id
  ]
}

resource "splight_file_folder" "FileFolderTest" {
  name = "FolderTF"
}

resource "splight_file" "FileInnerTest" {
  file        = "./variables.tf"
  description = "Sample file for testing inner file"
  parent      = splight_file_folder.FileFolderTest.id
}

# SECRETS
resource "splight_secret" "SecretTest" {
  name      = "SecretTest"
  raw_value = var.splight_secret
}

# IMPORT RESOURCES
resource "splight_asset" "AssetImportTest" {
  name        = "AssetImported"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}

resource "splight_secret" "SecretImportTest" {
  name      = "SecretImported"
  raw_value = var.splight_secret
}
