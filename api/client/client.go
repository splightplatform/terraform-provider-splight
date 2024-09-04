package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/splightplatform/terraform-provider-splight/api/settings"
)

// Client holds the configuration needed to communicate with a server
type Client struct {
	hostname   string       // Server hostname or IP address
	authToken  string       // Authorization token for HTTP requests
	orgID      string       // Splight organization ID
	httpClient *http.Client // Underlying HTTP client for making requests
	userAgent  string       // User-Agent header value for HTTP requests
	context    context.Context
}

// UserAgent defines the structure for constructing the User-Agent header
type UserAgent struct {
	ProductName    string            // Name of the product making the request
	ProductVersion string            // Version of the product making the request
	ExtraInfo      map[string]string // Additional information to include in the User-Agent header
}

// NewClient creates and configures a new Client instance
// hostname: Server hostname or IP address
// token: Authorization token for HTTP requests
// opts: UserAgent configuration for setting the User-Agent header
// Returns: A new Client instance or an error if configuration fails
func NewClient(context context.Context, opts UserAgent) (*Client, error) {
	config, err := settings.LoadSplightConfig(nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		hostname:   config.Workspaces[config.Current].Hostname,
		authToken:  config.Workspaces[config.Current].SecretKey,
		httpClient: &http.Client{Timeout: 60 * time.Second},
		context:    context,
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

// httpRequest performs an HTTP request with retry logic
// path: API endpoint path
// method: HTTP method (GET, POST, etc.)
// body: Request body to be sent (if applicable)
// Returns: Response body reader and nil on success, or an error if the request fails
func (c *Client) HttpRequest(path, method string, body bytes.Buffer) (io.ReadCloser, error) {
	var respBody io.ReadCloser
	var err error

	// Define max retry attempts and initial backoff duration
	maxAttempts := 3
	backoff := 2 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		respBody, err = c.doRequest(path, method, body, attempts)
		if err == nil {
			return respBody, nil
		}

		// Check if the HTTP response status code is an httpError
		if httpErr, ok := err.(*httpError); ok {
			// Retry only if the status code is 503 (Service Unavailable)
			if httpErr.statusCode != http.StatusServiceUnavailable {
				return nil, httpErr
			}
		}

		// Log the retry attempt details
		tflog.Trace(c.context, "retrying HTTP request on 503", map[string]interface{}{
			"path":      path,
			"method":    method,
			"body":      body.String(),
			"userAgent": c.userAgent,
			"attempt":   attempts,
			"error":     err,
		})

		// Exponential backoff with jitter
		time.Sleep(backoff)
		backoff *= 2
	}

	return nil, fmt.Errorf("failed after %d attempts: %w", maxAttempts, err)
}

// doRequest creates and sends an HTTP request
// path: API endpoint path
// method: HTTP method (GET, POST, etc.)
// body: Request body to be sent (if applicable)
// Returns: Response body reader and nil on success, or an error if the request fails
func (c *Client) doRequest(path, method string, body bytes.Buffer, attempt int) (io.ReadCloser, error) {
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

	// Log the request details
	tflog.Trace(c.context, "sending HTTP request", map[string]interface{}{
		"path":      path,
		"method":    method,
		"body":      body.String(),
		"userAgent": c.userAgent,
		"attempt":   attempt,
	})

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

	// Log the response details
	tflog.Trace(c.context, "received HTTP response", map[string]interface{}{
		"path":       path,
		"method":     method,
		"statusCode": resp.StatusCode,
		"body":       string(respBody),
		"attempt":    attempt,
	})

	if resp.StatusCode != statusCodeAccepted {
		return nil, &httpError{
			statusCode: resp.StatusCode,
			body:       string(respBody),
			message:    fmt.Sprintf("got a non-valid status code: %v - %s", resp.StatusCode, string(respBody)),
		}
	}

	return io.NopCloser(bytes.NewBuffer(respBody)), nil
}

// requestPath constructs the full request URL by appending the path to the hostname
// path: API endpoint path
// Returns: The full request URL
func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}

// httpError represents an HTTP error with details
type httpError struct {
	statusCode int    // HTTP status code of the error response
	body       string // Body of the error response
	message    string // Error message
}

// Error returns a string representation of the httpError
func (e *httpError) Error() string {
	return e.message
}
