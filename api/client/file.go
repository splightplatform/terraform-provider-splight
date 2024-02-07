package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type FileParams struct {
	Description   string         `json:"description"`
	Parent        string         `json:"parent"`
	RelatedAssets []RelatedAsset `json:"assets"`
}

type File struct {
	FileParams
	Checksum string `json:"checksum"`
	ID       string `json:"id"`
}

func (c *Client) ListFiles() (*map[string]File, error) {
	body, err := c.httpRequest("v2/engine/file/files/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]File{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateFile(item *FileParams, filepath string) (*File, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new multipart writer
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)

	// Create a part for the file content
	part, err := writer.CreateFormFile("file", filepath)
	if err != nil {
		return nil, err
	}

	// Copy the file content to the part
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Make the HTTP request
	req, err := http.NewRequest("POST", c.requestPath("v2/engine/file/files/"), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", c.authToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Unexpected status code: " + resp.Status)
	}

	// Decode the response body
	fileResponse := &File{}
	err = json.NewDecoder(resp.Body).Decode(fileResponse)
	if err != nil {
		return nil, err
	}

	return fileResponse, nil
}

func (c *Client) UpdateFile(id string, item *FileParams) (*File, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/file/files/%s/", id), "PATCH", buf)
	if err != nil {
		return nil, err
	}
	file := &File{}
	err = json.NewDecoder(body).Decode(file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (c *Client) RetrieveFile(id string) (*File, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/file/files/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	file := &File{}
	err = json.NewDecoder(body).Decode(file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (c *Client) DeleteFile(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/file/files/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
