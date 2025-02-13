package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetRelationParams struct {
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	RelatedAssetKind QueryFilter  `json:"related_asset_kind"`
	Asset            QueryFilter  `json:"asset"`
	RelatedAsset     *QueryFilter `json:"related_asset"`
}

type AssetRelation struct {
	AssetRelationParams
	Id string `json:"id"`
}

func (m *AssetRelation) GetId() string {
	return m.Id
}

func (m *AssetRelation) GetParams() Params {
	return &m.AssetRelationParams
}

func (m *AssetRelation) ResourcePath() string {
	return "v3/engine/asset/relations/"
}

func (m *AssetRelation) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	relatedAssetKind := d.Get("related_asset_kind").(*schema.Set).List()
	asset := d.Get("asset").(*schema.Set).List()

	relatedAsset := d.Get("related_asset").(*schema.Set).List()

	var parsedRelatedAsset *QueryFilter = nil
	if len(relatedAsset) == 0 {
		parsedRelatedAsset = convertSingleQueryFilter(relatedAsset)
	}

	m.AssetRelationParams = AssetRelationParams{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		RelatedAssetKind: *convertSingleQueryFilter(relatedAssetKind),
		Asset:            *convertSingleQueryFilter(asset),
		RelatedAsset:     parsedRelatedAsset,
	}

	return nil
}

func (m *AssetRelation) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("id", m.Id)
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("related_asset_kind", []map[string]any{
		{
			"id":   m.RelatedAssetKind.Id,
			"name": m.RelatedAssetKind.Name,
		},
	})
	d.Set("asset", []map[string]any{
		{
			"id":   m.Asset.Id,
			"name": m.Asset.Name,
		},
	})
	if m.RelatedAsset != nil {
		d.Set("related_asset", []map[string]any{
			{
				"id":   m.RelatedAsset.Id,
				"name": m.RelatedAsset.Name,
			},
		})
	}

	return nil
}
