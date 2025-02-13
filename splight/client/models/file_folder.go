package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type FileFolderParams struct {
	Name   string `json:"name"`
	Parent string `json:"parent,omitempty"`
}

type FileFolder struct {
	FileFolderParams
	Id string `json:"id"`
}

func (m *FileFolder) GetId() string {
	return m.Id
}

func (m *FileFolder) GetParams() Params {
	return &m.FileFolderParams
}

func (m *FileFolder) ResourcePath() string {
	return "v3/engine/file/folders/"
}

func (m *FileFolder) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	m.FileFolderParams = FileFolderParams{
		Name:   d.Get("name").(string),
		Parent: d.Get("parent").(string),
	}

	return nil
}

func (m *FileFolder) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("parent", m.Parent)

	return nil
}
