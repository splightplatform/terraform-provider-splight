package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type FileFolderParams struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type FileFolder struct {
	FileFolderParams
	ID string `json:"id"`
}

func (c *Client) ListFileFolders() (*map[string]FileFolder, error) {
	body, err := c.httpRequest("v2/engine/file/folders/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]FileFolder{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) CreateFileFolder(item *FileFolderParams) (*FileFolder, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("v2/engine/file/folders/", "POST", buf)
	if err != nil {
		return nil, err
	}

	folder := &FileFolder{}
	err = json.NewDecoder(body).Decode(folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (c *Client) UpdateFileFolder(id string, item *FileFolderParams) (*FileFolder, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/file/folders/%s/", id), "PUT", buf)
	if err != nil {
		return nil, err
	}
	folder := &FileFolder{}
	err = json.NewDecoder(body).Decode(folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (c *Client) RetrieveFileFolder(id string) (*FileFolder, error) {
	body, err := c.httpRequest(fmt.Sprintf("v2/engine/file/folders/%s/", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	folder := &FileFolder{}
	err = json.NewDecoder(body).Decode(folder)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (c *Client) DeleteFileFolder(id string) error {
	_, err := c.httpRequest(fmt.Sprintf("v2/engine/file/folders/%s/", id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
