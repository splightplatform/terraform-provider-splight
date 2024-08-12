---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "splight_dashboard_alertevents_chart Resource - terraform-provider-splight"
subcategory: ""
description: |-
  
---

# splight_dashboard_alertevents_chart (Resource)



## Example Usage

```terraform
terraform {
  required_providers {
    splight = {
      source = "splightplatform/splight"
    }
  }
}

resource "splight_asset" "AssetTest" {
  name        = "AssetTest"
  description = "Created with Terraform"
  geometry = jsonencode({
    type       = "GeometryCollection"
    geometries = []
  })
}

resource "splight_asset_attribute" "AttributeTest1" {
  name  = "Attribute1"
  type  = "Number"
  unit  = "meters"
  asset = splight_asset.AssetTest.id
}

resource "splight_asset_attribute" "AttributeTest2" {
  name  = "Attribute2"
  type  = "Number"
  unit  = "seconds"
  asset = splight_asset.AssetTest.id
}

resource "splight_dashboard" "DashboardTest" {
  name           = "DashboardTest"
  related_assets = []
}

resource "splight_dashboard_tab" "DashboardTabTest" {
  name      = "TabTest"
  order     = 0
  dashboard = splight_dashboard.DashboardTest.id
}

resource "splight_dashboard_alertevents_chart" "DashboardChartTest" {
  name               = "ChartTest"
  tab                = splight_dashboard_tab.DashboardTabTest.id
  timestamp_gte      = "now - 7d"
  timestamp_lte      = "now"
  description        = "Chart description"
  min_height         = 1
  min_width          = 4
  display_time_range = true
  labels_display     = true
  labels_aggregation = "last"
  labels_placement   = "bottom"
  show_beyond_data   = true
  height             = 10
  width              = 20
  collection         = "default"

  filter_name       = "some name"
  filter_old_status = ["warning"]
  filter_new_status = ["no_alert", "warning"]

  chart_items {
    ref_id           = "A"
    type             = "QUERY"
    color            = "red"
    expression_plain = ""
    query_filter_asset {
      id   = splight_asset.AssetTest.id
      name = splight_asset.AssetTest.name
    }
    query_filter_attribute {
      id   = splight_asset_attribute.AttributeTest1.id
      name = splight_asset_attribute.AttributeTest1.name
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest.id
          attribute = splight_asset_attribute.AttributeTest1.id
        }
      },
      {
        "$addFields" : {
          "timestamp" : {
            "$dateTrunc" : {
              "date" : "$timestamp",
              "unit" : "day",
              "binSize" : 1
            }
          }
        }
      },
      {
        "$group" : {
          "_id" : "$timestamp",
          "value" : {
            "$last" : "$value"
          },
          "timestamp" : {
            "$last" : "$timestamp"
          }
        }
      }
    ])
  }

  chart_items {
    ref_id           = "B"
    color            = "blue"
    type             = "QUERY"
    expression_plain = ""
    query_filter_asset {
      id   = splight_asset.AssetTest.id
      name = splight_asset.AssetTest.name
    }
    query_filter_attribute {
      id   = splight_asset_attribute.AttributeTest2.id
      name = splight_asset_attribute.AttributeTest2.name
    }
    query_plain = jsonencode([
      {
        "$match" = {
          asset     = splight_asset.AssetTest.id
          attribute = splight_asset_attribute.AttributeTest2.id
        }
      },
      {
        "$addFields" : {
          "timestamp" : {
            "$dateTrunc" : {
              "date" : "$timestamp",
              "unit" : "hour",
              "binSize" : 1
            }
          }
        }
      },
      {
        "$group" : {
          "_id" : "$timestamp",
          "value" : {
            "$last" : "$value"
          },
          "timestamp" : {
            "$last" : "$timestamp"
          }
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `chart_items` (Block Set, Min: 1) chart traces to be included (see [below for nested schema](#nestedblock--chart_items))
- `name` (String) name of the chart
- `tab` (String) id for the tab where to place the chart
- `timestamp_gte` (String) date in isoformat or shortcut string where to end reading
- `timestamp_lte` (String) date in isoformat or shortcut string where to start reading

### Optional

- `collection` (String)
- `description` (String) chart description
- `display_time_range` (Boolean) whether to display the time range or not
- `filter_name` (String) filter name
- `filter_new_status` (List of String) filter new status
- `filter_old_status` (List of String) filter old status
- `height` (Number) chart height in px
- `labels_aggregation` (String) [last|avg|...] aggregation
- `labels_display` (Boolean) whether to display the labels or not
- `labels_placement` (String) [right|bottom] placement
- `min_height` (Number) minimum chart height
- `min_width` (Number) minimum chart width
- `position_x` (Number) chart x position
- `position_y` (Number) chart y position
- `refresh_interval` (String) refresh interval
- `relative_window_time` (String) relative window time
- `show_beyond_data` (Boolean) whether to show data which is beyond timestamp_lte or not
- `thresholds` (Block Set) optional static lines to be added to the chart as references (see [below for nested schema](#nestedblock--thresholds))
- `timezone` (String) chart timezone
- `value_mappings` (Block Set) optional mappings to transform data with rules (see [below for nested schema](#nestedblock--value_mappings))
- `width` (Number) chart width in cols (max 20)

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--chart_items"></a>
### Nested Schema for `chart_items`

Required:

- `color` (String)
- `expression_plain` (String)
- `query_filter_asset` (Block Set, Min: 1, Max: 1) Asset/Attribute filter (see [below for nested schema](#nestedblock--chart_items--query_filter_asset))
- `query_filter_attribute` (Block Set, Min: 1, Max: 1) Asset/Attribute filter (see [below for nested schema](#nestedblock--chart_items--query_filter_attribute))
- `query_plain` (String)
- `ref_id` (String)
- `type` (String)

Optional:

- `hidden` (Boolean)
- `label` (String)
- `query_group_function` (String)
- `query_group_unit` (String)
- `query_limit` (Number)
- `query_sort_direction` (Number)

<a id="nestedblock--chart_items--query_filter_asset"></a>
### Nested Schema for `chart_items.query_filter_asset`

Optional:

- `id` (String) ID of the resource
- `name` (String) name of the resource


<a id="nestedblock--chart_items--query_filter_attribute"></a>
### Nested Schema for `chart_items.query_filter_attribute`

Optional:

- `id` (String) ID of the resource
- `name` (String) name of the resource



<a id="nestedblock--thresholds"></a>
### Nested Schema for `thresholds`

Required:

- `color` (String)
- `display_text` (String)
- `value` (Number)


<a id="nestedblock--value_mappings"></a>
### Nested Schema for `value_mappings`

Required:

- `display_text` (String)
- `match_value` (String)
- `order` (Number)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import [options] splight_dashboard_alertevents_chart.<name> <dashboard_chart_id>
```