package models

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SegmentParams struct {
	AssetParams
	Temperature        AssetAttribute `json:"temperature"`
	WindSpeed          AssetAttribute `json:"wind_speed"`
	WindDirection      AssetAttribute `json:"wind_direction"`
	Altitude           AssetMetadata  `json:"altitude"`
	Azimuth            AssetMetadata  `json:"azimuth"`
	CumulativeDistance AssetMetadata  `json:"cumulative_distance"`
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
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       json.RawMessage(d.Get("geometry").(string)),
			CustomTimezone: d.Get("timezone").(string),
			Tags:           tags,
			Kind:           kind,
		},
	}

	// TODO: remove ALL of these sets when API fixes its contract
	temperature := convertAssetAttribute(d.Get("temperature").(*schema.Set).List())
	if temperature == nil {
		temperature = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "temperature",
			},
		}
	}
	m.SegmentParams.Temperature = *temperature

	windSpeed := convertAssetAttribute(d.Get("wind_speed").(*schema.Set).List())
	if windSpeed == nil {
		windSpeed = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "wind_speed",
			},
		}
	}
	m.SegmentParams.WindSpeed = *windSpeed

	windDirection := convertAssetAttribute(d.Get("wind_direction").(*schema.Set).List())
	if windDirection == nil {
		windDirection = &AssetAttribute{
			AssetAttributeParams: AssetAttributeParams{
				Type: "Number",
				Name: "wind_direction",
			},
		}
	}
	m.SegmentParams.WindDirection = *windDirection

	altitude := convertAssetMetadata(d.Get("altitude").(*schema.Set).List())
	if altitude.Type == "" {
		altitude.Type = "Number"
	}
	if altitude.Name == "" {
		altitude.Name = "altitude"
	}
	m.SegmentParams.Altitude = *altitude

	azimuth := convertAssetMetadata(d.Get("azimuth").(*schema.Set).List())
	if azimuth.Type == "" {
		azimuth.Type = "Number"
	}
	if azimuth.Name == "" {
		azimuth.Name = "azimuth"
	}
	m.SegmentParams.Azimuth = *azimuth

	cumulativeDistance := convertAssetMetadata(d.Get("cumulative_distance").(*schema.Set).List())
	if cumulativeDistance.Type == "" {
		cumulativeDistance.Type = "Number"
	}
	if cumulativeDistance.Name == "" {
		cumulativeDistance.Name = "cumulative_distance"
	}
	m.SegmentParams.CumulativeDistance = *cumulativeDistance

	return nil
}

func (m *Segment) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)
	d.Set("geometry", string(m.AssetParams.Geometry))
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

	d.Set("temperature", []map[string]any{m.Temperature.ToMap()})
	d.Set("wind_speed", []map[string]any{m.WindSpeed.ToMap()})
	d.Set("wind_direction", []map[string]any{m.WindDirection.ToMap()})
	d.Set("altitude", []map[string]any{m.Altitude.ToMap()})
	d.Set("azimuth", []map[string]any{m.Azimuth.ToMap()})
	d.Set("cumulative_distance", []map[string]any{m.CumulativeDistance.ToMap()})

	return nil
}
