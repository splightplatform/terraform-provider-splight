package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetMetadataParams struct {
	Asset string          `json:"asset"`
	Name  string          `json:"name"`
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
	Unit  string          `json:"unit,omitempty"` // API complains when unit is empty
}

type AssetMetadata struct {
	AssetMetadataParams
	Id string `json:"id,omitempty"` // NOTE: needed for Line, Segment, etc
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
		"unit":  m.Unit,
	}

	return result
}

func (m *AssetMetadata) ResourcePath() string {
	return "v2/engine/asset/metadata/"
}

func convertAssetMetadata(data []any) (*AssetMetadata, error) {
	if len(data) == 0 {
		return nil, nil
	}
	metadataMap := data[0].(map[string]any)

	// Validate value JSON
	valueStr := metadataMap["value"].(string)
	if err := validateJSONString(valueStr); err != nil {
		return nil, fmt.Errorf("metadata value JSON must be json encoded")
	}

	return &AssetMetadata{
		AssetMetadataParams: AssetMetadataParams{
			Asset: metadataMap["asset"].(string),
			Name:  metadataMap["name"].(string),
			Type:  metadataMap["type"].(string),
			Value: json.RawMessage(valueStr),
			Unit:  metadataMap["unit"].(string),
		},
		Id: metadataMap["id"].(string),
	}, nil
}

func (m *AssetMetadata) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	// Validate geometry JSON
	valueStr := d.Get("geometry").(string)
	if err := validateJSONString(valueStr); err != nil {
		return fmt.Errorf("metadata value JSON must be json encoded")
	}

	m.AssetMetadataParams = AssetMetadataParams{
		Asset: d.Get("asset").(string),
		Name:  d.Get("name").(string),
		Type:  d.Get("type").(string),
		Value: json.RawMessage(valueStr),
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
