package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type AssetMetadataParams struct {
	Asset string `json:"asset"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Unit  string `json:"unit,omitempty"`
}

type AssetMetadata struct {
	AssetMetadataParams
	ID string `json:"id"`
}

func (m *AssetMetadata) GetID() string {
	return m.ID
}

func (m *AssetMetadata) GetParams() Params {
	return &m.AssetMetadataParams
}

func (m *AssetMetadata) ResourcePath() string {
	return "v2/engine/asset/metadata/"
}

func (m *AssetMetadata) FromSchema(d *schema.ResourceData) error {
	m.AssetMetadataParams = AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
		Unit:  d.Get("unit").(string),
	}

	return nil
}

func (m *AssetMetadata) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)

	d.Set("asset", m.Asset)
	d.Set("name", m.Name)
	d.Set("type", m.Type)
	d.Set("value", m.Value)
	d.Set("unit", m.Unit)

	return nil
}
