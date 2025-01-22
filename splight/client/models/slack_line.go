package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SlackLineParams struct {
	AssetParams
	SwitchStatusStart *AssetAttribute `json:"switch_status_start"`
	SwitchStatusEnd   *AssetAttribute `json:"switch_status_end"`
}

type SlackLine struct {
	SlackLineParams
	Id string `json:"id"`
}

func (m *SlackLine) GetId() string {
	return m.Id
}

func (m *SlackLine) GetParams() Params {
	return &m.SlackLineParams
}

func (m *SlackLine) ResourcePath() string {
	return "v2/engine/asset/slack-lines/"
}

func (m *SlackLine) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Validate geometry JSON
	geometryStr := d.Get("geometry").(string)
	if err := validateJSONString(geometryStr); err != nil {
		return fmt.Errorf("geometry must be a JSON encoded GeoJSON")
	}

	m.SlackLineParams = SlackLineParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(geometryStr),
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	return nil
}

func (m *SlackLine) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)
	d.Set("geometry", string(m.AssetParams.Geometry))
	d.Set("timezone", m.AssetParams.CustomTimezone)
	d.Set("switch_status_start", []map[string]any{m.SwitchStatusStart.ToMap()})
	d.Set("switch_status_end", []map[string]any{m.SwitchStatusEnd.ToMap()})

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
