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
		"unit":  m.Unit,
	}

	// NOTE: convert floats to ints without loosing decimals.
	// This is done because API casts the values
	var rawValue any
	if err := json.Unmarshal(m.Value, &rawValue); err == nil {
		// Check if the value is a float
		if floatVal, ok := rawValue.(float64); ok {
			// Check if the float can be converted to an integer without loss
			if float64(int(floatVal)) == floatVal {
				result["value"] = int(floatVal)
			} else {
				result["value"] = string(m.Value)
			}
		} else {
			result["value"] = string(m.Value)
		}
	} else {
		result["value"] = string(m.Value)
	}

	return result
}

func (m *AssetMetadata) ResourcePath() string {
	return "v3/engine/asset/metadata/"
}

func convertAssetMetadata(data []any) (*AssetMetadata, error) {
	// Handle empty input
	if len(data) == 0 {
		return &AssetMetadata{AssetMetadataParams: AssetMetadataParams{}}, nil
	}

	// Type assertion for metadata map
	metadataMap := data[0].(map[string]any)

	// Extract value and validate JSON
	valueStr := metadataMap["value"].(string)
	if err := validateJSONString(valueStr); err != nil {
		return nil, fmt.Errorf("metadata value must be JSON encoded")
	}

	// Construct and return AssetMetadata
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
		return fmt.Errorf("metadata value must be JSON encoded")
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
