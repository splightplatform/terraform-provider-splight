package models

type SecretParams struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Secret struct {
	SecretParams
	ID string `json:"id"`
}
