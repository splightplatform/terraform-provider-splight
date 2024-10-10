package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
)

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

	body, httpErr := c.HttpRequest(url, method, buf)
	if httpErr != nil {
		return httpErr
	}

	err := json.NewDecoder(body).Decode(m)
	if err != nil {
		return err
	}

	if fileModel, ok := any(m).(*models.File); ok {
		if !fileModel.Uploaded {
			// TODO: delete model if this fails
			err := c.UploadFile(fileModel)
			if err != nil {
				return err
			}

			// So we do not try to upload the file again
			fileModel.Uploaded = true

			err = c.UpdateFileChecksum(fileModel)
			if err != nil {
				return fmt.Errorf("error retrieving checksum for file: %w", err)
			}
		}
	}

	return nil
}

func Retrieve[T models.SplightModel](c *Client, m T, id string) error {
	url := fmt.Sprintf("%s%s/", m.ResourcePath(), id)

	body, httpErr := c.HttpRequest(url, http.MethodGet, bytes.Buffer{})
	if httpErr != nil {
		return httpErr
	}

	err := json.NewDecoder(body).Decode(m)
	if err != nil {
		return fmt.Errorf("error decoding model: %w", err)
	}

	if fileModel, ok := any(m).(*models.File); ok {
		httpErr := c.UpdateFileChecksum(fileModel)
		if httpErr != nil {
			return fmt.Errorf("error retrieving checksum for file: %w", err)
		}
	}

	return nil
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
