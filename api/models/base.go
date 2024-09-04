package models

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/splightplatform/terraform-provider-splight/api/client"
)

type SplightDatabaseBaseModel struct {
	ID     string `json:"id"`
	Client *client.Client
}

func (b *SplightDatabaseBaseModel) ResourcePath() (string, error) {
	return "", errors.New("ResourcePath() must be implemented by the embedding model")
}

// TODO: List method
func (b *SplightDatabaseBaseModel) Save() error {
	url, err := b.ResourcePath()
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if err := json.NewEncoder(&buf).Encode(b); err != nil {
		return err
	}

	method := http.MethodPost
	if b.ID != "" {
		method = http.MethodPatch
	}

	body, err := b.Client.HttpRequest(url, method, buf)
	if err != nil {
		return err
	}

	return json.NewDecoder(body).Decode(b)
}

func (b *SplightDatabaseBaseModel) Retrieve(id string) error {
	url, err := b.ResourcePath()
	if err != nil {
		return err
	}

	body, err := b.Client.HttpRequest(fmt.Sprintf("%s%s/", url, id), "GET", bytes.Buffer{})
	if err != nil {
		return err
	}

	return json.NewDecoder(body).Decode(b)
}

func (b *SplightDatabaseBaseModel) Delete(id string) error {
	// TODO: remove the ID from the struct
	url, err := b.ResourcePath()
	if err != nil {
		return err
	}

	_, err = b.Client.HttpRequest(fmt.Sprintf("%s%s/", url, id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func NewSplightDatabaseBaseModel(ctx context.Context) (*SplightDatabaseBaseModel, error) {
	// TODO: set a default like "splight database model, and version dev"
	userAgentOptions := client.UserAgent{
		ProductName:    ctx.Value("ProductName"),
		ProductVersion: ctx.Value("Version"),
	}

	client, err := client.NewClient(ctx, userAgentOptions)

	if err != nil {
		return nil, fmt.Errorf("Error creating client: %v", err)
	}

	return &SplightDatabaseBaseModel{
		Client: client,
	}, nil
}
