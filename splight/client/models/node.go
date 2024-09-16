package models

type NodeParams struct {
	InstanceType   string `json:"instance_type"`
	Name           string `json:"name"`
	OrganizationId string `json:"organization_id"`
	Region         string `json:"region"`
}

type Node struct {
	NodeParams
	Id string `json:"id"`
}
