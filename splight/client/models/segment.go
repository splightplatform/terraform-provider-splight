package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SegmentParams struct {
	AssetParams
	Temperature        *AssetAttribute `json:"temperature,omitempty"`
	WindSpeed          *AssetAttribute `json:"wind_speed,omitempty"`
	WindDirection      *AssetAttribute `json:"wind_direction,omitempty"`
	Altitude           *AssetMetadata  `json:"altitude,omitempty,omitempty"`
	Azimuth            *AssetMetadata  `json:"azimuth,omitempty,omitempty"`
	CumulativeDistance *AssetMetadata  `json:"cumulative_distance,omitempty"`
}

type Segment struct {
	SegmentParams
	Id string `json:"id"`
}

func (m *Segment) GetId() string {
	return m.Id
}

func (m *Segment) GetParams() Params {
	return &m.SegmentParams
}

func (m *Segment) ResourcePath() string {
	return "v2/engine/asset/segments/"
}

func (m *Segment) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	m.SegmentParams = SegmentParams{
		AssetParams: AssetParams{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Geometry:    json.RawMessage(d.Get("geometry").(string)),
			Tags:        tags,
			Kind:        kind,
		},
	}

	altitude := convertAssetMetadata(d.Get("altitude").(*schema.Set).List())
	if altitude != nil {
		altitude.Type = "Number"
		altitude.Name = "altitude"
	}
	m.SegmentParams.Altitude = altitude

	azimuth := convertAssetMetadata(d.Get("azimuth").(*schema.Set).List())
	if azimuth != nil {
		azimuth.Type = "Number"
		azimuth.Name = "azimuth"
	}
	m.SegmentParams.Azimuth = azimuth

	cumulativeDistance := convertAssetMetadata(d.Get("cumulative_distance").(*schema.Set).List())
	if cumulativeDistance != nil {
		cumulativeDistance.Type = "Number"
		cumulativeDistance.Name = "cumulative_distance"
	}
	m.SegmentParams.CumulativeDistance = cumulativeDistance

	return nil
}

func (m *Segment) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)
	d.Set("geometry", string(m.AssetParams.Geometry))

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

	d.Set("temperature", []map[string]any{m.Temperature.ToMap()})
	d.Set("wind_speed", []map[string]any{m.WindSpeed.ToMap()})
	d.Set("wind_direction", []map[string]any{m.WindDirection.ToMap()})
	d.Set("altitude", []map[string]any{m.Altitude.ToMap()})
	d.Set("azimuth", []map[string]any{m.Azimuth.ToMap()})
	d.Set("cumulative_distance", []map[string]any{m.CumulativeDistance.ToMap()})

	return nil
}
