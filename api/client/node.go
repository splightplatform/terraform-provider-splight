package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type NodeParams struct {
	InstanceType   string `json:"instance_type"`
	Name           string `json:"name"`
	OrganizationId string `json:"organization_id"`
	Region         string `json:"region"`
}

type Node struct {
	NodeParams
	ID string `json:"id"`
}

func (c *Client) ListNodes() (*map[string]Node, error) {
	body, err := c.httpRequest("v2/backoffice/compute/nodes/splighthosted/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Node{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateNode(item *NodeParams) (*Node, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/backoffice/compute/nodes/splighthosted/", "POST", buf)
	if err != nil {
		return nil, err
	}

	node := &Node{}
	err = json.NewDecoder(body).Decode(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) RetrieveNode(id string) (*Node, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/backoffice/compute/nodes/splighthosted/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	node := &Node{}
	err = json.NewDecoder(body).Decode(node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) DeleteNode(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/backoffice/compute/nodes/splighthosted/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
