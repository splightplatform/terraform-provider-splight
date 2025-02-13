package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SlackGeneratorParams struct {
	AssetParams
}

type SlackGenerator struct {
	SlackGeneratorParams
	Id string `json:"id"`
}

func (m *SlackGenerator) GetId() string {
	return m.Id
}

func (m *SlackGenerator) GetParams() Params {
	return &m.SlackGeneratorParams
}

func (m *SlackGenerator) ResourcePath() string {
	return "v3/engine/asset/slack-generators/"
}

func (m *SlackGenerator) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()
	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Validate geometry JSON
	geometryStr := d.Get("geometry").(string)
	if err := validateJSONString(geometryStr); err != nil {
		return fmt.Errorf("geometry field contains %w", err)
	}

	geometry := json.RawMessage(geometryStr)
	m.SlackGeneratorParams = SlackGeneratorParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       &geometry,
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	return nil
}

func (m *SlackGenerator) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)

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
