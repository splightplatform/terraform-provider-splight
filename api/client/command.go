package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CommandParams struct {
	Name    string        `json:"name"`
	Asset   *QueryFilter  `json:"asset"`
	Actions []QueryFilter `json:"actions"`
}

type Command struct {
	CommandParams
	ID string `json:"id"`
}

func (c *Client) ListCommands() (*map[string]Command, error) {
	body, err := c.httpRequest("v2/engine/command/commands/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Command{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateCommand(item *CommandParams) (*Command, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/command/commands/", "POST", buf)
	if err != nil {
		return nil, err
	}

	data := &Command{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) UpdateCommand(id string, item *CommandParams) (*Command, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/command/commands/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	data := &Command{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) RetrieveCommand(id string) (*Command, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/commands/commands/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	data := &Command{}
	err = json.NewDecoder(body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) DeleteCommand(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/commands/commands/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
