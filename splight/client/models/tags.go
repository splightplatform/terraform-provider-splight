package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Tags struct {
	Tags []Tag `json:"results"`
}

func (m *Tags) ResourcePath() string {
	return "v2/account/tags/"
}

func (m *Tags) ToSchema(d *schema.ResourceData) error {
	var tagsMap []map[string]string

	for _, tag := range m.Tags {
		tagMap := map[string]string{
			"id":   tag.Id,
			"name": tag.Name,
		}
		tagsMap = append(tagsMap, tagMap)
	}

	d.Set("tags", tagsMap)
	d.SetId("tags")

	return nil
}
