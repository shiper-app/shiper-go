package shiper

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

	gen, err := genapi.NewClientWithResponses(baseURL, config.clientOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{raw: gen}, nil
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

type rawClient = genapi.ClientWithResponses

// response is satisfied by every generated *<Operation>Response.
type response interface {
	StatusCode() int
	GetBody() []byte
}

// call is the shared body of every wrapper: it validates the required
// name/value pairs, runs fn against the generated client, enforces a 2xx
// status and, when the spec declares a JSON body for that status, requires it
// to have been decoded.
func call[R response](ctx context.Context, c *Client, op string, fn func(context.Context, *rawClient) (R, error), required ...string) (R, error) {
	var zero R
	for i := 0; i+1 < len(required); i += 2 {
		if err := validateRequiredString(required[i], required[i+1]); err != nil {
			return zero, err
		}
	}

	raw, err := c.getRawClient()
	if err != nil {
		return zero, err
	}

	resp, err := fn(normalizeContext(ctx), raw)
	if err != nil {
		return zero, err
	}
	if err := require2xx(op, resp.StatusCode(), resp.GetBody()); err != nil {
		return zero, err
	}
	if body := reflect.ValueOf(resp).Elem().FieldByName("JSON" + strconv.Itoa(resp.StatusCode())); body.IsValid() && body.IsNil() {
		return zero, newAPIError(op, resp.StatusCode(), resp.GetBody())
	}

	return resp, nil
}
