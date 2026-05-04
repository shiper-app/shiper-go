package shiper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/shiper-app/shiper-go/internal/genapi"
)

var errClientNotInitialized = errors.New("shiper client is not initialized")

const defaultBaseURL = "https://api.shiper.app/v2"
const envBaseURL = "SHIPER_BASE_URL"

type Client struct {
	raw *genapi.ClientWithResponses
}

func NewClient(opts ...Option) (*Client, error) {
	config := clientOptions{}
	for _, opt := range opts {
		if opt != nil {
			opt(&config)
		}
	}

	baseURL := strings.TrimSpace(os.Getenv(envBaseURL))
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	rawClient, err := genapi.NewClientWithResponses(baseURL, config.clientOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{raw: rawClient}, nil
}

func (c *Client) RawClient() *genapi.ClientWithResponses {
	if c == nil {
		return nil
	}
	return c.raw
}

func (c *Client) getRawClient() (*genapi.ClientWithResponses, error) {
	if c == nil || c.raw == nil {
		return nil, errClientNotInitialized
	}
	return c.raw, nil
}

func validateRequiredString(fieldName, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func normalizeContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
