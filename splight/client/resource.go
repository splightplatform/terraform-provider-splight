package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
)

// TODO: Support files
func Save[T models.SplightModel](c *Client, m T) error {
	url := m.ResourcePath()

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(m.GetParams()); err != nil {
		return err
	}

	method := http.MethodPost
	if m.GetId() != "" {
		method = http.MethodPatch
		url = fmt.Sprintf("%s%s/", url, m.GetId())
	}

	body, err := c.HttpRequest(url, method, buf)
	if err != nil {
		return err
	}

	return json.NewDecoder(body).Decode(m)
}

func Retrieve[T models.SplightModel](c *Client, m T, id string) error {
	url := fmt.Sprintf("%s%s/", m.ResourcePath(), id)

	body, err := c.HttpRequest(url, http.MethodGet, bytes.Buffer{})
	if err != nil {
		return err
	}

	return json.NewDecoder(body).Decode(m)
}

func List[T models.DataSource](c *Client, m T) error {
	url := m.ResourcePath()

	body, err := c.HttpRequest(url, http.MethodGet, bytes.Buffer{})
	if err != nil {
		return err
	}

	return json.NewDecoder(body).Decode(m)
}

func Delete[T models.SplightModel](c *Client, m T, id string) error {
	url := fmt.Sprintf("%s%s/", m.ResourcePath(), id)

	_, err := c.HttpRequest(url, http.MethodDelete, bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
