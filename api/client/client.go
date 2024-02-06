package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname string, token string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	req.Header.Add("Content-Type", "application/json")
	statusCodeAccepted := http.StatusOK
	switch method {
	case "GET":
		statusCodeAccepted = http.StatusOK
	case "DELETE":
		statusCodeAccepted = http.StatusNoContent
	case "POST":
		statusCodeAccepted = http.StatusCreated
	case "PUT":
		statusCodeAccepted = http.StatusOK
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != statusCodeAccepted && resp.StatusCode != http.StatusNotFound {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read from body: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("Got a non valid status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}
