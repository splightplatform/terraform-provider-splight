package models

type TagParams struct {
	Name string `json:"name"`
}

type Tag struct {
	SplightDatabaseBaseModel
	TagParams
}

type Tags struct {
	Tags []Tag `json:"results"`
}

func (t Tag) ResourcePath() string {
	return "v2/account/tags/"
}
