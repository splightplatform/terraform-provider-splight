package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}

func NewClient(hostname string, token string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)

	if err != nil {
		return nil, err
	}

	// Add Splight auth token
	req.Header.Add("Authorization", c.authToken)
	req.Header.Add("Content-Type", "application/json")

	statusCodeAccepted := http.StatusOK

	switch method {
	case "DELETE":
		statusCodeAccepted = http.StatusNoContent
	case "POST":
		statusCodeAccepted = http.StatusCreated
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Element not found: %v", resp.StatusCode)
	}
	if resp.StatusCode != statusCodeAccepted {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read from body: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("Got a non valid status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}
