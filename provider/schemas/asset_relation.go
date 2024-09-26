package schemas

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func SchemaAssetRelation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "relation id",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "relation name",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "relation description",
		},
		"related_asset_kind": {
			Type:        schema.TypeSet,
			Required:    true,
			MaxItems:    1,
			Description: "kind of the target relation asset",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "kind name",
					},
				},
			},
		},
		"asset": {
			Type:        schema.TypeSet,
			Required:    true,
			MaxItems:    1,
			ForceNew:    true,
			Description: "asset where the relation origins",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "asset id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "asset name",
					},
				},
			},
		},
		"related_asset": {
			Type:        schema.TypeSet,
			Required:    true,
			MaxItems:    1,
			Description: "target asset of the relation",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "asset id",
					},
					"name": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "asset name",
					},
				},
			},
		},
	}
}
