// Check if the file exists
// TODO: que pasa si ya cree el file, lo borro y al hacer refresh ya no existe el viejo?force new en ese caso?

package models

import (
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type FileParams struct {
	Path          string
	Checksum      string
	Uploaded      bool
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Parent        string        `json:"parent"`
	Tags          []QueryFilter `json:"tags"`
	RelatedAssets []QueryFilter `json:"assets"`
}

type File struct {
	FileParams
	Id string `json:"id"`
}

func (m *File) GetId() string {
	return m.Id
}

func (m *File) GetParams() Params {
	return &m.FileParams
}

func (m *File) ResourcePath() string {
	return "v2/engine/file/files/"
}

func (m *File) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	tags := convertQueryFilters(d.Get("tags").(*schema.Set).List())
	assets := convertQueryFilters(d.Get("related_assets").(*schema.Set).List())
	path := d.Get("path").(string)

	m.FileParams = FileParams{
		Path:          path,
		Checksum:      d.Get("checksum").(string),
		Uploaded:      d.Get("uploaded").(bool),
		Name:          filepath.Base(path),
		Description:   d.Get("description").(string),
		Parent:        d.Get("parent").(string),
		Tags:          tags,
		RelatedAssets: assets,
	}

	return nil
}

func (m *File) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	if d.Get("checksum").(string) != "" {
		if m.Checksum != d.Get("checksum").(string) {
			d.Set("path", nil)
			return nil
		}
	}

	d.Set("checksum", m.Checksum)
	d.Set("uploaded", true)
	d.Set("description", m.Description)
	d.Set("parent", m.Parent)

	var relatedasets []map[string]any
	for _, relatedAsset := range m.RelatedAssets {
		relatedasets = append(relatedasets, map[string]any{
			"id":   relatedAsset.Id,
			"name": relatedAsset.Name,
		})
	}
	d.Set("related_assets", relatedasets)

	var tags []map[string]any
	for _, tag := range m.Tags {
		tags = append(tags, map[string]any{
			"id":   tag.Id,
			"name": tag.Name,
		})
	}
	d.Set("tags", tags)

	return nil
}
