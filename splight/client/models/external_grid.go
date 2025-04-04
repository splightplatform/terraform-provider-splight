package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExternalGridParams struct {
	AssetParams
	Bus  *AssetRelationship `json:"bus,omitempty"`
	Grid *AssetRelationship `json:"grid,omitempty"`
}

type ExternalGrid struct {
	ExternalGridParams
	Id string `json:"id"`
}

func (m *ExternalGrid) GetId() string {
	return m.Id
}

func (m *ExternalGrid) GetParams() Params {
	return &m.ExternalGridParams
}

func (m *ExternalGrid) ResourcePath() string {
	return "v3/engine/asset/external-grids/"
}

func (m *ExternalGrid) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Get values of timezone and geometry
	timezone := d.Get("timezone").(string)
	geometryStr := d.Get("geometry").(string)
	busId := d.Get("bus").(string)
	gridId := d.Get("grid").(string)

	var busRel *AssetRelationship = nil
	if busId != "" {
		busRel = &AssetRelationship{
			RelatedAssetId: ResourceId{
				Id: busId,
			},
		}
	}

	var gridRel *AssetRelationship = nil
	if gridId != "" {
		gridRel = &AssetRelationship{
			RelatedAssetId: ResourceId{
				Id: gridId,
			},
		}
	}

	// Validate geometry JSON if it's set
	if geometryStr != "" {
		if err := validateJSONString(geometryStr); err != nil {
			return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
		}
	}

	// Check if geometryStr is empty and handle accordingly
	var geometry *json.RawMessage
	if geometryStr != "" {
		// Convert string to json.RawMessage
		raw := json.RawMessage(geometryStr)
		geometry = &raw
	}

	m.ExternalGridParams = ExternalGridParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       geometry,
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
		Bus:  busRel,
		Grid: gridRel,
	}

	return nil
}

func (m *ExternalGrid) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)

	if m.Bus != nil {
		d.Set("bus", m.Bus.RelatedAssetId.Id)
	} else {
		d.Set("bus", "")
	}

	if m.Grid != nil {
		d.Set("grid", m.Grid.RelatedAssetId.Id)
	} else {
		d.Set("grid", "")
	}

	var geometryStr string
	if m.Geometry != nil {
		geometryStr = string(*m.Geometry)
	} else {
		geometryStr = ""
	}
	d.Set("geometry", geometryStr)

	d.Set("timezone", m.AssetParams.CustomTimezone)

	var tags []map[string]any
	for _, tag := range m.AssetParams.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	d.Set("kind", []map[string]any{
		{
			"id":   m.AssetParams.Kind.Id,
			"name": m.AssetParams.Kind.Name,
		},
	})

	return nil
}
