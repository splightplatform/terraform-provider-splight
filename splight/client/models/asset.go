package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AssetParams struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Geometry       json.RawMessage `json:"geometry"`
	CustomTimezone string          `json:"timezone,omitempty"`
	Tags           []QueryFilter   `json:"tags"`
	Kind           *QueryFilter    `json:"kind"`
}

type Asset struct {
	AssetParams
	Id string `json:"id"`
}

func (m *Asset) GetId() string {
	return m.Id
}

func (m *Asset) GetParams() Params {
	return &m.AssetParams
}

func (m *Asset) ResourcePath() string {
	return "v3/engine/asset/assets/"
}

// validateTimezoneOrGeometry ensures that either one of `timezone` or `geometry` is set,
// or both are set with the same value.
func validateTimezoneOrGeometry(timezone, geometryStr string) error {
	if timezone != "" && geometryStr != "" {
		// If both are set, check if they are equal
		if timezone != geometryStr {
			return fmt.Errorf("if both 'timezone' and 'geometry' are set, they must have the same value")
		}
	}
	return nil
}

func (m *Asset) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Get values of timezone and geometry
	timezone := d.Get("timezone").(string)
	geometryStr := d.Get("geometry").(string)

	// Validate timezone or geometry (or both equal)
	if err := validateTimezoneOrGeometry(timezone, geometryStr); err != nil {
		return err
	}

	// Validate geometry JSON if it's set
	if geometryStr != "" {
		if err := validateJSONString(geometryStr); err != nil {
			return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
		}
	}

	m.AssetParams = AssetParams{
		Name:           d.Get("name").(string),
		Description:    d.Get("description").(string),
		Geometry:       json.RawMessage(geometryStr),
		CustomTimezone: timezone,
		Tags:           tags,
		Kind:           kind,
	}

	return nil
}

func (m *Asset) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("description", m.Description)
	d.Set("geometry", string(m.Geometry))
	d.Set("timezone", m.CustomTimezone)

	var tags []map[string]any
	for _, tag := range m.AssetParams.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	if m.Kind != nil {
		d.Set("kind", []map[string]any{
			{
				"id":   m.Kind.Id,
				"name": m.Kind.Name,
			},
		})
	}

	return nil
}
