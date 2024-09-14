package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TagParams struct {
	Name string `json:"name"`
}

type Tag struct {
	TagParams
	ID string `json:"id"`
}

func (m *Tag) GetID() string {
	return m.ID
}

func (m *Tag) GetParams() Params {
	return &m.TagParams
}

func (m *Tag) ResourcePath() string {
	return "v2/account/tags/"
}

func (m *Tag) FromSchema(d *schema.ResourceData) error {
	m.TagParams = TagParams{
		Name: d.Get("name").(string),
	}
	m.ID = d.Id()

	return nil
}

func (m *Tag) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.ID)
	d.Set("name", m.Name)

	return nil
}
