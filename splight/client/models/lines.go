package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Lines struct {
	Lines []QueryFilter `json:"results"`
}

func (m *Lines) ResourcePath() string {
	return "v2/engine/asset/lines/"
}

func (m *Lines) ToSchema(d *schema.ResourceData) error {
	var linesMap []map[string]string

	for _, queryFilter := range m.Lines {
		lineMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		linesMap = append(linesMap, lineMap)
	}

	d.Set("lines", linesMap)
	d.SetId("lines")

	return nil
}
