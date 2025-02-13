package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/splightplatform/terraform-provider-splight/splight/client/models"
)

// RetrieveEmail fetches the email of the current user.
func (c *Client) RetrieveEmail() (string, error) {
	body, err := c.HttpRequest("auth/account/user/profile/", "GET", bytes.Buffer{})
	if err != nil {
		return "", fmt.Errorf("error making profile request: %w", err)
	}
	defer body.Close()

	profile := make(map[string]interface{})
	if err := json.NewDecoder(body).Decode(&profile); err != nil {
		return "", fmt.Errorf("error decoding profile response: %w", err)
	}

	email, ok := profile["name"].(string)
	if !ok {
		return "", fmt.Errorf("user email not found in profile response")
	}
	return email, nil
}

type FileDetails struct {
	Checksum string `json:"checksum"`
}

// UpdateFileChecksum fetches the checksum of a file and unescapes it
func (c *Client) UpdateFileChecksum(model *models.File) error {
	// Make the HTTP request to fetch file details
	body, httpErr := c.HttpRequest(fmt.Sprintf("%s%s/details", model.ResourcePath(), model.GetId()), "GET", bytes.Buffer{})
	if httpErr != nil {
		return fmt.Errorf("error making file details request: %w", httpErr)
	}
	defer body.Close()

	var fileDetails FileDetails

	// Decode the JSON response into fileDetails struct
	if err := json.NewDecoder(body).Decode(&fileDetails); err != nil {
		return fmt.Errorf("error decoding file details response: %w", err)
	}

	// Unescape the checksum string
	unescapedChecksum, err := strconv.Unquote(fileDetails.Checksum)
	if err != nil {
		return fmt.Errorf("error unescaping checksum: %w", httpErr)
	}

	// Assign the unescaped checksum to the model
	model.Checksum = unescapedChecksum

	return nil
}

// UploadFile uploads a file by retrieving the upload URL first.
func (c *Client) UploadFile(model *models.File) error {
	// Step 1: Retrieve the upload URL
	body, httpErr := c.HttpRequest(fmt.Sprintf("%s%s/upload_url/", model.ResourcePath(), model.GetId()), "GET", bytes.Buffer{})
	if httpErr != nil {
		return fmt.Errorf("error retrieving upload URL: %w", httpErr)
	}
	defer body.Close()

	// Step 2: Decode the response directly into a map to extract the upload URL
	var response map[string]interface{}
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding upload URL response: %w", err)
	}

	uploadURL, ok := response["url"].(string)
	if !ok {
		return fmt.Errorf("upload URL not found in the response")
	}

	// Step 3: Open the file to be uploaded
	file, err := os.Open(model.Path)
	if err != nil {
		return fmt.Errorf("error opening file for upload: %w", err)
	}
	defer file.Close()

	// Step 4: Get file size for the upload request
	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error retrieving file stats: %w", err)
	}
	fileSize := fileStat.Size()

	// Step 5: Create a PUT request to upload the file
	req, err := http.NewRequest("PUT", uploadURL, file)
	if err != nil {
		return fmt.Errorf("error creating PUT request: %w", err)
	}
	req.ContentLength = fileSize

	startTime := time.Now()

	// Step 6: Perform the file upload
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error performing file upload request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("file upload failed with status code: %d", resp.StatusCode)
	}

	uploadTime := time.Since(startTime)

	// Step 7: Log successful upload with more context
	tflog.Debug(c.context, "File uploaded successfully", map[string]interface{}{
		"filePath":   model.Path,
		"fileSize":   fileSize,
		"uploadURL":  uploadURL,
		"uploadTime": uploadTime,
	})

	return nil
}
