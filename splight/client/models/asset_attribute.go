package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetAttributeParams struct {
	Asset string `json:"asset"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Unit  string `json:"unit,omitempty"`
}

type AssetAttribute struct {
	AssetAttributeParams
	Id string `json:"id"`
}

func (m *AssetAttribute) GetId() string {
	return m.Id
}

func (m *AssetAttribute) GetParams() Params {
	return &m.AssetAttributeParams
}

func (m *AssetAttribute) ResourcePath() string {
	return "v2/engine/asset/attributes/"
}

func (m *AssetAttribute) FromSchema(d *schema.ResourceData) error {
	m.AssetAttributeParams = AssetAttributeParams{
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Asset: d.Get("asset").(string),
		Unit:  d.Get("unit").(string),
	}
	return nil
}

func (m *AssetAttribute) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("type", m.Type)
	d.Set("asset", m.Asset)
	d.Set("unit", m.Unit)

	return nil
}
