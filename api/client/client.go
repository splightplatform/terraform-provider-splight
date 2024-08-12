package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
	userAgent  string
}

// NewClient initializes the client with the given hostname and token, and configures the User-Agent header
func NewClient(hostname, token string) (*Client, error) {
	client := &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	// Retrieve the email to configure the User-Agent
	email, err := client.RetrieveEmail()
	if err != nil {
		return nil, err
	}

	// TODO: set version from linker
	client.userAgent = fmt.Sprintf("terraform-provider-splight-version-%s", email)
	return client, nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (io.ReadCloser, error) {
	var respBody io.ReadCloser
	var err error

	// Define max retry attempts and initial backoff duration
	maxAttempts := 3
	backoff := 2 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		respBody, err = c.doRequest(path, method, body)
		if err == nil {
			return respBody, nil
		}

		// If error is not due to a 503 status code, don't retry
		if err, ok := err.(*httpError); ok && err.statusCode != http.StatusServiceUnavailable {
			return nil, err
		}

		log.Printf("Attempt %d: %v", attempts, err)

		// Exponential backoff with jitter
		time.Sleep(backoff)
		backoff *= 2
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", maxAttempts, err)
}

func (c *Client) doRequest(path, method string, body bytes.Buffer) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.userAgent)
	statusCodeAccepted := http.StatusOK

	switch method {
	case http.MethodDelete:
		statusCodeAccepted = http.StatusNoContent
	case http.MethodPost:
		statusCodeAccepted = http.StatusCreated
	}

	log.Printf("Sending %s request to %s", method, req.URL)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("element not found: %v", resp.StatusCode)
	}

	if resp.StatusCode != statusCodeAccepted {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read from body: %v", resp.StatusCode)
		}
		return nil, &httpError{
			statusCode: resp.StatusCode,
			body:       string(respBody),
			message:    fmt.Sprintf("got a non-valid status code: %v - %s", resp.StatusCode, string(respBody)),
		}
	}

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return io.NopCloser(bytes.NewBuffer(respBody)), nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}

func (c *Client) RetrieveOrgId() (string, error) {
	body, err := c.httpRequest("v2/account/user/organizations/", "GET", bytes.Buffer{})
	if err != nil {
		return "", err
	}
	orgs := map[string]interface{}{}
	err = json.NewDecoder(body).Decode(&orgs)
	if err != nil {
		return "", err
	}
	if len(orgs) == 0 {
		return "", fmt.Errorf("No organizations found")
	}
	orgId := orgs["id"].(string)
	return orgId, nil
}

func (c *Client) RetrieveEmail() (string, error) {
	body, err := c.httpRequest("v2/account/user/profile/", "GET", bytes.Buffer{})
	if err != nil {
		return "", err
	}
	profile := map[string]interface{}{}
	err = json.NewDecoder(body).Decode(&profile)
	if err != nil {
		return "", err
	}
	if email, ok := profile["email"].(string); ok {
		return email, nil
	}
	return "", fmt.Errorf("User email not found")
}

// httpError represents an HTTP error with a status code and message
type httpError struct {
	statusCode int
	body       string
	message    string
}

func (e *httpError) Error() string {
	return e.message
}
