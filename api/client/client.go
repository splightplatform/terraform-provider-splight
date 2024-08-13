package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

// Client holds all of the information required to connect to a server
type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
	userAgent  string
}

type UserAgent struct {
	ProductName    string
	ProductVersion string
	ExtraInfo      map[string]string
}

// NewClient returns a new client configured to communicate on a server with the
// given hostname and port and to send an Authorization Header with the value of
// token
func NewClient(hostname, token string, opts UserAgent) (*Client, error) {
	client := &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{Timeout: 60 * time.Second},
	}

	// Retrieve the email to configure the User-Agent
	email, err := client.RetrieveEmail()
	if err != nil {
		return nil, err
	}

	// Get system details and default values
	defaultInfo := map[string]string{
		"email": email,
		"OS":    runtime.GOOS,
		"Arch":  runtime.GOARCH,
		"Go":    runtime.Version(),
	}

	// Merge default values with provided options
	for key, value := range opts.ExtraInfo {
		defaultInfo[key] = value
	}

	// Construct the User-Agent string
	userAgent := fmt.Sprintf("%s/%s", opts.ProductName, opts.ProductVersion)
	for key, value := range defaultInfo {
		userAgent += fmt.Sprintf(";%s=%s", key, value)
	}

	client.userAgent = userAgent

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

		// Check if the HTTP response status code it's an httpError
		if httpErr, ok := err.(*httpError); ok {
			// Retry only if the status code is 503 (Service Unavailable)
			if httpErr.statusCode != http.StatusServiceUnavailable {
				return nil, httpErr
			}
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

	// Read the response body into a buffer
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != statusCodeAccepted {
		return nil, &httpError{
			statusCode: resp.StatusCode,
			body:       string(respBody),
			message:    fmt.Sprintf("got a non-valid status code: %v - %s", resp.StatusCode, string(respBody)),
		}
	}

	return io.NopCloser(bytes.NewBuffer(respBody)), nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
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
