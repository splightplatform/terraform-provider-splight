package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type RelatedAsset struct {
	Id string `json:"id"`
}

type AssetParams struct {
	Geometry      json.RawMessage `json:"geometry"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	RelatedAssets []RelatedAsset  `json:"assets"`
	Tags          []Tag           `json:"tags"`
	Kind          *AssetKind      `json:"kind"`
}

type Asset struct {
	AssetParams
	ID string `json:"id"`
}

func (m *Asset) GetID() string {
	return m.ID
}

func (m *Asset) GetParams() Params {
	return &m.AssetParams
}

func (m *Asset) ResourcePath() string {
	return "v2/engine/asset/assets/"
}

func (m *Asset) FromSchema(d *schema.ResourceData) error {
	relatedAssetsSet := d.Get("related_assets").(*schema.Set).List()
	relatedAssets := make([]RelatedAsset, len(relatedAssetsSet))
	for i, relatedAsset := range relatedAssetsSet {
		relatedAssets[i] = RelatedAsset{
			Id: relatedAsset.(string),
		}
	}

	var kind *AssetKind
	kinds := d.Get("kind").(*schema.Set).List()
	if len(kinds) > 0 {
		kindMap := kinds[0].(map[string]interface{})
		kind = &AssetKind{
			ID: kindMap["id"].(string),
			AssetKindParams: AssetKindParams{
				Name: kindMap["name"].(string),
			},
		}
	}

	tagsInterface := d.Get("tags").(*schema.Set).List()
	tags := make([]Tag, len(tagsInterface))
	for i, item := range tagsInterface {
		tagItem := item.(map[string]interface{})
		tags[i] = Tag{
			ID: tagItem["id"].(string),
			TagParams: TagParams{
				Name: tagItem["name"].(string),
			},
		}
	}

	m.AssetParams = AssetParams{
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Geometry:      json.RawMessage(d.Get("geometry").(string)),
		RelatedAssets: relatedAssets,
		Tags:          tags,
		Kind:          kind,
	}
	m.ID = d.Id()

	return nil
}

func (m *Asset) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)
	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("related_assets", m.RelatedAssets)
	d.Set("geometry", string(m.Geometry))
	d.Set("tags", m.Tags)

	if m.Kind != nil {
		d.Set("kind", []map[string]interface{}{
			{
				"id":   m.Kind.ID,
				"name": m.Kind.Name,
			},
		})
	} else {
		d.Set("kind", nil)
	}

	return nil
}
