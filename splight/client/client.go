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
	"github.com/splightplatform/terraform-provider-splight/splight/settings"
)

// Client holds the configuration needed to communicate with a server
type Client struct {
	hostname   string       // Server hostname or IP address
	authToken  string       // Authorization token for HTTP requests
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
func NewClient(context context.Context, opts UserAgent) (*Client, error) {
	splightConfig, err := settings.LoadSplightConfig(nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		hostname:   splightConfig.Hostname,
		authToken:  splightConfig.Token,
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

// HttpRequest performs an HTTP request with retry logic
func (c *Client) HttpRequest(path, method string, body bytes.Buffer) (io.ReadCloser, *HttpError) {
	var respBody io.ReadCloser
	var err error

	maxAttempts := 3
	backoff := 2 * time.Second

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		respBody, err = c.doRequest(path, method, body, attempts)
		if err == nil {
			return respBody, nil
		}

		if httpErr, ok := err.(*HttpError); ok {
			if httpErr.StatusCode != http.StatusServiceUnavailable {
				return nil, httpErr
			}
		}

		tflog.Trace(c.context, "retrying HTTP request on 503", map[string]interface{}{
			"path":      path,
			"method":    method,
			"body":      body.String(),
			"userAgent": c.userAgent,
			"attempt":   attempts,
			"error":     err,
		})

		time.Sleep(backoff)
		backoff *= 2
	}

	return nil, &HttpError{
		StatusCode: http.StatusRequestTimeout,
		Message:    fmt.Sprintf("failed after %d attempts: %v", maxAttempts, err),
	}
}

// doRequest creates and sends an HTTP request
func (c *Client) doRequest(path, method string, body bytes.Buffer, attempt int) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, &HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("failed to create HTTP request: %v", err),
		}
	}

	req.Header.Add("Authorization", c.authToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.userAgent)

	statusCodeAccepted := http.StatusOK
	if method == http.MethodDelete {
		statusCodeAccepted = http.StatusNoContent
	} else if method == http.MethodPost {
		statusCodeAccepted = http.StatusCreated
	}

	tflog.Trace(c.context, "sending HTTP request", map[string]interface{}{
		"path":      path,
		"method":    method,
		"body":      body.String(),
		"userAgent": c.userAgent,
		"attempt":   attempt,
	})

	// Log the request details
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &HttpError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("HTTP request failed: %v", err),
		}
	}
	defer resp.Body.Close()

	// Read the response body into a buffer
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &HttpError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("failed to read response body: %v", err),
		}
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
		return nil, &HttpError{
			StatusCode: resp.StatusCode,
			Body:       string(respBody),
			Message:    fmt.Sprintf("unexpected status code: %v - %s", resp.StatusCode, string(respBody)),
		}
	}

	return io.NopCloser(bytes.NewBuffer(respBody)), nil
}

// requestPath constructs the full request URL by appending the path to the hostname
// path: API endpoint path
func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s", c.hostname, path)
}

// HttpError represents an HTTP error with details
type HttpError struct {
	StatusCode int // HTTP status code of the error response
	Body       string
	Message    string
}

// Error returns the error message
func (e *HttpError) Error() string {
	return e.Message
}
