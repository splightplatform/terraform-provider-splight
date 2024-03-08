package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FileParams struct {
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Parent        string         `json:"parent"`
	RelatedAssets []RelatedAsset `json:"assets"`
}

type FileURL struct {
	URL string `json:"url"`
}

type FileDetails struct {
	Checksum     string `json:"checksum"`
	LastModified string `json:"last_modified"`
	Size         string `json:"size"`
}

type File struct {
	FileParams
	ID string `json:"id"`
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
	// Create the File instance
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/file/files/", "POST", buf)
	if err != nil {
		return nil, err
	}
	fileResponse := &File{}
	err = json.NewDecoder(body).Decode(fileResponse)
	if err != nil {
		return nil, err
	}

	// Request File upload_url
	body, err = c.httpRequest(fmt.Sprintf("v2/engine/file/files/%s/upload_url/", fileResponse.ID), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	fileURL := &FileURL{}
	err = json.NewDecoder(body).Decode(fileURL)
	if err != nil {
		return nil, err
	}

	// Upload File content
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Getting size
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileStat.Size()
	// Do the request
	req, err := http.NewRequest("PUT", fileURL.URL, file)
	if err != nil {
		return nil, err
	}
	req.ContentLength = fileSize
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	statusCode := resp.StatusCode
	if statusCode != 200 {
		return nil, fmt.Errorf("wrong status code uploading file %d", statusCode)
	}
	defer resp.Body.Close()

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

func (c *Client) RetrieveFileDetails(id string) (*FileDetails, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/file/files/%s/details", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, nil
	}
	fileDetails := &FileDetails{}
	err = json.NewDecoder(body).Decode(fileDetails)
	if err != nil {
		return nil, err
	}
	return fileDetails, nil
}
