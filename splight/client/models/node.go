package models

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NodeParams struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Node struct {
	NodeParams
	Id string `json:"id"`
}

func (m *Node) GetId() string {
	return m.Id
}

func (m *Node) GetParams() Params {
	return &m.NodeParams
}

func (m *Node) ResourcePath() string {
	return "v3/engine/compute/nodes/all/"
}

func (m *Node) FromSchema(d *schema.ResourceData) error {
	m.NodeParams = NodeParams{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}
	m.Id = d.Id()

	return nil
}

func (m *Node) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)
	d.Set("name", m.Name)
	d.Set("type", m.Type)

	return nil
}
