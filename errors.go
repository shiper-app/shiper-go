package shiper

import (
	"fmt"
	"net/http"
	"strings"
)

type APIError struct {
	Operation  string
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	if e == nil {
		return "api error"
	}

	statusText := http.StatusText(e.StatusCode)
	if statusText == "" {
		statusText = "Unknown"
	}

	body := strings.TrimSpace(string(e.Body))
	if body == "" {
		return fmt.Sprintf("%s failed with HTTP %d %s", e.Operation, e.StatusCode, statusText)
	}

	if len(body) > 512 {
		body = body[:512] + "..."
	}

	return fmt.Sprintf("%s failed with HTTP %d %s: %s", e.Operation, e.StatusCode, statusText, body)
}

func newAPIError(operation string, statusCode int, body []byte) error {
	bodyCopy := append([]byte(nil), body...)
	return &APIError{
		Operation:  operation,
		StatusCode: statusCode,
		Body:       bodyCopy,
	}
}

func requireStatus(operation string, statusCode int, body []byte, expected ...int) error {
	for _, code := range expected {
		if statusCode == code {
			return nil
		}
	}
	return newAPIError(operation, statusCode, body)
}

func require2xx(operation string, statusCode int, body []byte) error {
	if statusCode >= 200 && statusCode < 300 {
		return nil
	}
	return newAPIError(operation, statusCode, body)
}
