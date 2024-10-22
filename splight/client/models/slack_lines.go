package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type SlackLines struct {
	SlackLines []QueryFilter `json:"results"`
}

func (m *SlackLines) ResourcePath() string {
	return "v2/engine/asset/slack-lines/"
}

func (m *SlackLines) ToSchema(d *schema.ResourceData) error {
	var linesMap []map[string]string

	for _, queryFilter := range m.SlackLines {
		lineMap := map[string]string{
			"id":   queryFilter.Id,
			"name": queryFilter.Name,
		}
		linesMap = append(linesMap, lineMap)
	}

	d.Set("slack_lines", linesMap)
	d.SetId("slack_lines")

	return nil
}
