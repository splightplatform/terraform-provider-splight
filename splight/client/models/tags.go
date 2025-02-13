package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Tags struct {
	Tags []QueryFilter `json:"results"`
}

func (m *Tags) ResourcePath() string {
	return "v3/engine/tags/"
}

func (m *Tags) ToSchema(d *schema.ResourceData) error {
	var tagsMap []map[string]string

	for _, queryFilter := range m.Tags {
		tagMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		tagsMap = append(tagsMap, tagMap)
	}

	d.Set("tags", tagsMap)
	d.SetId("tags")

	return nil
}
