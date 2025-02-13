package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TagParams struct {
	Name string `json:"name"`
}

type Tag struct {
	TagParams
	Id string `json:"id"`
}

func (m *Tag) GetId() string {
	return m.Id
}

func (m *Tag) GetParams() Params {
	return &m.TagParams
}

func (m *Tag) ResourcePath() string {
	return "v3/engine/tags/"
}

func (m *Tag) FromSchema(d *schema.ResourceData) error {
	m.TagParams = TagParams{
		Name: d.Get("name").(string),
	}
	m.Id = d.Id()

	return nil
}

func (m *Tag) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)
	d.Set("name", m.Name)

	return nil
}
