package models

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type SegmentParams struct {
	AssetParams
	Temperature          *AssetAttribute    `json:"temperature"`
	WindSpeed            *AssetAttribute    `json:"wind_speed"`
	WindDirection        *AssetAttribute    `json:"wind_direction"`
	Altitude             AssetMetadata      `json:"altitude"`
	Azimuth              AssetMetadata      `json:"azimuth"`
	CumulativeDistance   AssetMetadata      `json:"cumulative_distance"`
	ReferenceSag         AssetMetadata      `json:"reference_sag"`
	ReferenceTemperature AssetMetadata      `json:"reference_temperature"`
	SpanLength           AssetMetadata      `json:"span_length"`
	Line                 *AssetRelationship `json:"line"`
}

type ResourceId struct {
	Id string `json:"id"`
}

type AssetRelationship struct {
	RelatedAssetId ResourceId `json:"related_asset"`
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
	return "v3/engine/asset/segments/"
}

func validateJSONString(s string) error {
	if s == "" {
		return nil
	}
	var js json.RawMessage
	if err := json.Unmarshal([]byte(s), &js); err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}
	return nil
}

func (m *Segment) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	kind := convertSingleQueryFilter(d.Get("kind").(*schema.Set).List())
	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())

	// Get values of timezone and geometry
	timezone := d.Get("timezone").(string)
	geometryStr := d.Get("geometry").(string)
	lineId := d.Get("line").(string)

	var lineRel *AssetRelationship = nil
	if lineId != "" {
		lineRel = &AssetRelationship{
			RelatedAssetId: ResourceId{
				Id: lineId,
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

	m.SegmentParams = SegmentParams{
		AssetParams: AssetParams{
			Name:           d.Get("name").(string),
			Description:    d.Get("description").(string),
			Geometry:       geometry,
			CustomTimezone: timezone,
			Tags:           tags,
			Kind:           kind,
		},
		Line: lineRel,
	}

	altitude, err := convertAssetMetadata(d.Get("altitude").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid altitude metadata: %w", err)
	}
	if altitude.Type == "" {
		altitude.Type = "Number"
	}
	if altitude.Name == "" {
		altitude.Name = "altitude"
	}
	m.SegmentParams.Altitude = *altitude

	azimuth, err := convertAssetMetadata(d.Get("azimuth").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid azimuth metadata: %w", err)
	}
	if azimuth.Type == "" {
		azimuth.Type = "Number"
	}
	if azimuth.Name == "" {
		azimuth.Name = "azimuth"
	}
	m.SegmentParams.Azimuth = *azimuth

	cumulativeDistance, err := convertAssetMetadata(d.Get("cumulative_distance").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid cumulative distance metadata: %w", err)
	}
	if cumulativeDistance.Type == "" {
		cumulativeDistance.Type = "Number"
	}
	if cumulativeDistance.Name == "" {
		cumulativeDistance.Name = "cumulative_distance"
	}
	m.SegmentParams.CumulativeDistance = *cumulativeDistance

	referenceSag, err := convertAssetMetadata(d.Get("reference_sag").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid reference sag metadata: %w", err)
	}
	if referenceSag.Type == "" {
		referenceSag.Type = "Number"
	}
	if referenceSag.Name == "" {
		referenceSag.Name = "reference_sag"
	}
	m.SegmentParams.ReferenceSag = *referenceSag

	referenceTemperature, err := convertAssetMetadata(d.Get("reference_temperature").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid reference temperature metadata: %w", err)
	}
	if referenceTemperature.Type == "" {
		referenceTemperature.Type = "Number"
	}
	if referenceTemperature.Name == "" {
		referenceTemperature.Name = "reference_temperature"
	}
	m.SegmentParams.ReferenceTemperature = *referenceTemperature

	spanLength, err := convertAssetMetadata(d.Get("span_length").(*schema.Set).List())
	if err != nil {
		return fmt.Errorf("invalid span length metadata: %w", err)
	}
	if spanLength.Type == "" {
		spanLength.Type = "Number"
	}
	if spanLength.Name == "" {
		spanLength.Name = "span_length"
	}
	m.SegmentParams.SpanLength = *spanLength

	return nil
}

func (m *Segment) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.AssetParams.Name)
	d.Set("description", m.AssetParams.Description)

	// TODO: set the rel asset id

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

	d.Set("temperature", []map[string]any{m.Temperature.ToMap()})
	d.Set("wind_speed", []map[string]any{m.WindSpeed.ToMap()})
	d.Set("wind_direction", []map[string]any{m.WindDirection.ToMap()})
	d.Set("altitude", []map[string]any{m.Altitude.ToMap()})
	d.Set("azimuth", []map[string]any{m.Azimuth.ToMap()})
	d.Set("cumulative_distance", []map[string]any{m.CumulativeDistance.ToMap()})
	d.Set("reference_sag", []map[string]any{m.ReferenceSag.ToMap()})
	d.Set("reference_temperature", []map[string]any{m.ReferenceTemperature.ToMap()})
	d.Set("span_length", []map[string]any{m.SpanLength.ToMap()})

	return nil
}
