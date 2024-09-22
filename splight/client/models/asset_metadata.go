package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetMetadataParams struct {
	Asset string          `json:"asset"`
	Name  string          `json:"name"`
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
	Unit  string          `json:"unit,omitempty"`
}

type AssetMetadata struct {
	AssetMetadataParams
	Id string `json:"id"`
}

func (m *AssetMetadata) GetId() string {
	return m.Id
}

func (m *AssetMetadata) GetParams() Params {
	return &m.AssetMetadataParams
}

func (m *AssetMetadata) ToMap() map[string]any {
	result := map[string]any{
		"id":    m.Id,
		"asset": m.Asset,
		"name":  m.Name,
		"type":  m.Type,
		"value": string(m.Value),
	}

	// TODO: validate
	// Only include "unit" if it's not empty (omitempty behavior)
	if m.Unit != "" {
		result["unit"] = m.Unit
	}

	return result
}

func (m *AssetMetadata) ResourcePath() string {
	return "v2/engine/asset/metadata/"
}

func convertAssetMetadata(data []any) *AssetMetadata {
	if len(data) == 0 {
		return nil
	}
	metadataMap := data[0].(map[string]any)
	return &AssetMetadata{
		AssetMetadataParams: AssetMetadataParams{
			Asset: metadataMap["asset"].(string),
			Name:  metadataMap["name"].(string),
			Type:  metadataMap["type"].(string),
			Value: json.RawMessage(metadataMap["value"].(string)),
			Unit:  metadataMap["unit"].(string),
		},
		Id: metadataMap["id"].(string),
	}
}

func (m *AssetMetadata) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	m.AssetMetadataParams = AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: json.RawMessage(d.Get("value").(string)),
		Unit:  d.Get("unit").(string),
	}

	return nil
}

func (m *AssetMetadata) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("asset", m.Asset)
	d.Set("name", m.Name)
	d.Set("type", m.Type)
	d.Set("value", m.Value)
	d.Set("unit", m.Unit)

	return nil
}
