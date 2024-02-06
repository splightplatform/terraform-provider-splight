terraform {
  required_providers {
    spl = {
      source  = "local/splight/spl"
      version = "0.1.0"
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

provider "spl" {
  hostname = var.spl_hostname
  token    = var.spl_api_token
}

# ASSETS
resource "spl_asset" "AssetTest" {
  for_each    = local.assets
  name        = each.key
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = each.value.geometries
  })
}

resource "spl_asset_attribute" "AssetTestAttribute" {
  for_each = local.asset_attribute_combinations
  name     = each.value.attributeKey
  type     = each.value.attribute.type

  asset = spl_asset.AssetTest[each.value.assetKey].id
}

resource "spl_asset_attribute" "AssetTestFunctionAttribute" {
  for_each = local.assets
  name     = "FunctionAttribute"
  type     = "Number"
  asset    = spl_asset.AssetTest[each.key].id
}

resource "spl_asset_metadata" "AssetTestMetadata" {
  for_each = local.asset_metadata_combinations
  name     = each.value.metadataKey
  value    = each.value.metadata.value
  type     = each.value.metadata.type
  asset    = spl_asset.AssetTest[each.value.assetKey].id
}

# COMPONENTS
resource "spl_component" "ComponentTest" {
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

resource "spl_component_routine" "ComponentTestRoutine" {
  for_each = local.asset_attribute_combinations

  name         = "ComponentTestRoutine-${each.key}"
  description  = "Created with Terraform"
  type         = "IncomingRoutine"
  component_id = spl_component.ComponentTest.id

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
      "asset" : spl_asset.AssetTest[each.value.assetKey].id,
      "attribute" : spl_asset_attribute.AssetTestAttribute[each.key].id
    })
  }
}

# FUNCTIONS AND ALERTS
resource "spl_function" "FunctionTest" {
  for_each        = local.assets
  name            = "FunctionTest-${each.key}"
  description     = "Created with Terraform"
  type            = "rate"
  time_window     = 600
  rate_value      = 10
  rate_unit       = "minute"
  target_variable = "A"
  target_asset = {
    id   = spl_asset.AssetTest[each.key].id
    name = spl_asset.AssetTest[each.key].name
  }
  target_attribute = {
    id   = spl_asset_attribute.AssetTestFunctionAttribute[each.key].id
    name = spl_asset_attribute.AssetTestFunctionAttribute[each.key].name
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

resource "spl_alert" "AlertTest" {
  for_each        = local.assets
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
          asset     = spl_asset.AssetTest[each.key].id
          attribute = spl_asset_attribute.AssetTestFunctionAttribute[each.key].id
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
          asset     = spl_asset.AssetTest[each.key].id
          attribute = spl_asset_attribute.AssetTestFunctionAttribute[each.key].id
        }
      }
    ])
  }

}

# DASHBOARDS
resource "spl_dashboard" "DashboardTest" {
  for_each = local.assets
  name     = "DashboardTest-${each.key}"
}

resource "spl_dashboard_tab" "DashboardTabTest" {
  for_each  = local.assets
  name      = "TabTest"
  order     = 0
  dashboard = spl_dashboard.DashboardTest[each.key].id
}

resource "spl_dashboard_chart" "DashboardChartTest" {
  for_each      = spl_dashboard_tab.DashboardTabTest
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
          asset     = spl_asset.AssetTest[each.key].id
          attribute = spl_asset_attribute.AssetTestAttribute["${each.key}/AttributeTF1"].id
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
          asset     = spl_asset.AssetTest[each.key].id
          attribute = spl_asset_attribute.AssetTestAttribute["${each.key}/AttributeTF2"].id
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
resource "spl_file" "FileTest" {
  file        = ".idea/FileTF"
  description = "Sample file for testing"
}

resource "spl_file_folder" "FileFolderTest" {
  name = "FolderTF"
}

resource "spl_file_folder" "FileFolderInnerTest" {
  name   = "InnerFolderTF"
  parent = spl_file_folder.FileFolderTest.id
}

resource "spl_file" "FileInnerTest" {
  file        = ".idea/FileTFCopy"
  description = "Sample file for testing inner file"
  parent      = spl_file_folder.FileFolderInnerTest.id
}

# SECRETS
resource "spl_secret" "SecretTest" {
  name      = "SecretTest"
  raw_value = var.spl_secret
}

# IMPORT RESOURCES
resource "spl_asset" "AssetImportTest" {
  name        = "AssetImported"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}

resource "spl_secret" "SecretImportTest" {
  name      = "SecretImported"
  raw_value = var.spl_secret
}
