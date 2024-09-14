package models

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// TODO: can we enforce Params to be a pointer type?
type Params interface{}

type Identifiable interface {
	GetID() string
}

type Pathable interface {
	ResourcePath() string
}

type SchemaReadable interface {
	FromSchema(d *schema.ResourceData) error
}

type SchemaWritable interface {
	ToSchema(d *schema.ResourceData) error
}

type ParamsProvider interface {
	GetParams() Params
}

type SplightObject interface{}

type DataSource interface {
	SplightObject
	Pathable
	SchemaWritable
}

type SplightModel interface {
	SplightObject
	Identifiable
	ParamsProvider
	Pathable
	SchemaReadable
	SchemaWritable
}
