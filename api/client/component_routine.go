package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ComponentRoutineConfigParam struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Multiple    bool             `json:"multiple"`
	Required    bool             `json:"required"`
	Sensitive   bool             `json:"sensitive"`
	Type        string           `json:"type"`
	Value       *json.RawMessage `json:"value"`
}

type ComponentRoutineDataAddress struct {
	Asset     string `json:"asset"`
	Attribute string `json:"attribute"`
}

type ComponentRoutineDataAddresses []ComponentRoutineDataAddress

func (c *ComponentRoutineDataAddresses) UnmarshalJSON(data []byte) error {
	// Attempt to unmarshal data into a single ComponentRoutineDataAddress
	var single ComponentRoutineDataAddress
	if err := json.Unmarshal(data, &single); err == nil {
		*c = ComponentRoutineDataAddresses{single}
		return nil
	}

	// Attempt to unmarshal data into a slice of ComponentRoutineDataAddress
	var slice []ComponentRoutineDataAddress
	if err := json.Unmarshal(data, &slice); err == nil {
		*c = slice
		return nil
	}

	return fmt.Errorf("failed to unmarshal ComponentRoutineDataAddresses")
}

type ComponentRoutineIOParam struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ValueType   string                        `json:"value_type"`
	Multiple    bool                          `json:"multiple"`
	Required    bool                          `json:"required"`
	Value       ComponentRoutineDataAddresses `json:"value"`
}

type ComponentRoutineParams struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Type        string                        `json:"type"`
	ComponentId string                        `json:"component_id"`
	Input       []ComponentRoutineIOParam     `json:"input"`
	Output      []ComponentRoutineIOParam     `json:"output"`
	Config      []ComponentRoutineConfigParam `json:"config"`
}

type ComponentRoutine struct {
	ComponentRoutineParams
	ID string `json:"id"`
}

func (c *Client) ListComponentRoutines() (*map[string]ComponentRoutine, error) {
	body, err := c.httpRequest("v2/engine/component/routines/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]ComponentRoutine{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateComponentRoutine(item *ComponentRoutineParams) (*ComponentRoutine, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/component/routines/", "POST", buf)
	if err != nil {
		return nil, err
	}

	component := &ComponentRoutine{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) UpdateComponentRoutine(id string, item *ComponentRoutineParams) (*ComponentRoutine, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/component/routines/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	component := &ComponentRoutine{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) RetrieveComponentRoutine(id string) (*ComponentRoutine, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/component/routines/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	component := &ComponentRoutine{}
	err = json.NewDecoder(body).Decode(component)
	if err != nil {
		return nil, err
	}
	return component, nil
}

func (c *Client) DeleteComponentRoutine(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/component/routines/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
