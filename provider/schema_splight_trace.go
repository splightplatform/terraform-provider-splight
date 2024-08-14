package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Trace() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		Description: "traces to be used to compute the results",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:        schema.TypeString,
					Computed:    true,
					Description: "ID of the function item",
				},
				"ref_id": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "identifier of the variable (i.e 'A')",
				},
				"type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "either QUERY or EXPRESSION",
					ValidateFunc: validation.StringInSlice([]string{
						"QUERY",
						"EXPRESSION",
					}, false),
				},
				"expression": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "how the expression is shown (i.e 'A * 2')",
				},
				"expression_plain": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "actual mongo query containing the expression",
				},
				"query_filter_asset":     QueryFilter(),
				"query_filter_attribute": QueryFilter(),
				"query_group_filter": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "function used to aggregate data",
					ValidateFunc: validation.StringInSlice([]string{
						"max",
						"min",
						"avg",
						"sum",
						"last",
					}, false),
				},
				"query_group_unit": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "time window to apply the aggregation",
					ValidateFunc: validation.StringInSlice([]string{
						"second",
						"minute",
						"hour",
						"day",
						"month",
					}, false),
				},
				"query_plain": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "actual mongo query",
				},
			},
		},
	}
}
