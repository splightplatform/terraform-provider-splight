package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type SecretParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Secret struct {
	SecretParams
	Id string `json:"id"`
}

func (m *Secret) GetId() string {
	return m.Id
}

func (m *Secret) GetParams() Params {
	return &m.SecretParams
}

func (m *Secret) ResourcePath() string {
	return "v2/engine/secret/secrets/"
}

func (m *Secret) FromSchema(d *schema.ResourceData) error {
	m.Id = d.Id()

	rawValue := d.Get("raw_value").(string)
	m.SecretParams = SecretParams{
		Name:  d.Get("name").(string),
		Value: rawValue,
	}

	return nil
}

func (m *Secret) ToSchema(d *schema.ResourceData) error {
	d.SetId(m.Id)

	d.Set("name", m.Name)
	d.Set("value", m.Value)

	return nil
}
