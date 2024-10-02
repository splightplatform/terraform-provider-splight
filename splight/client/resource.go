package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
)

type FileURL struct {
	URL string `json:"url"`
}

type FileDetails struct {
	Checksum     string `json:"checksum"`
	LastModified string `json:"last_modified"`
	Size         string `json:"size"`
}

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

	if fileModel, ok := any(m).(models.File); ok {
		// Request file upload url
		body, httpErr = c.HttpRequest(fmt.Sprintf("v2/engine/file/files/%s/upload_url/", fileModel.GetId()), "GET", bytes.Buffer{})
		if httpErr != nil {
			return httpErr
		}

		fileURL := &FileURL{}
		err = json.NewDecoder(body).Decode(fileURL)
		if err != nil {
			return err
		}

		// Upload file content
		file, err := os.Open(fileModel.Path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Getting size
		fileStat, err := file.Stat()
		if err != nil {
			return err
		}
		fileSize := fileStat.Size()

		// Do the request
		req, err := http.NewRequest("PUT", fileURL.URL, file)
		if err != nil {
			return err
		}
		req.ContentLength = fileSize
		// TODO: use the HttpRequest method
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		statusCode := resp.StatusCode
		if statusCode != 200 {
			return fmt.Errorf("wrong status code uploading file %d", statusCode)
		}
		defer resp.Body.Close()
		// TODO: set the checksum and delete the db object is this process fails

	}

	return nil
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
