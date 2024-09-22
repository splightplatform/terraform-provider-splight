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

func (m *AssetAttribute) ToMap() map[string]any {
	result := map[string]any{
		"id":    m.Id,
		"asset": m.Asset,
		"name":  m.Name,
		"type":  m.Type,
	}

	// TODO: validate
	// Only include "unit" if it's not empty (omitempty behavior)
	if m.Unit != "" {
		result["unit"] = m.Unit
	}

	return result
}

func (m *AssetAttribute) ResourcePath() string {
	return "v2/engine/asset/attributes/"
}

func convertAssetAttribute(data []any) *AssetAttribute {
	if len(data) == 0 {
		return nil
	}
	attributeMap := data[0].(map[string]any)
	return &AssetAttribute{
		AssetAttributeParams: AssetAttributeParams{
			Asset: attributeMap["asset"].(string),
			Name:  attributeMap["name"].(string),
			Type:  attributeMap["type"].(string),
			Unit:  attributeMap["unit"].(string),
		},
		Id: attributeMap["id"].(string),
	}
}

func (m *AssetAttribute) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

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
