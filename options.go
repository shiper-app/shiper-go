package shiper

import (
	"context"
	"net/http"
	"strings"

	"github.com/shiper-app/shiper-go/internal/genapi"
)

type Option func(*clientOptions)

type clientOptions struct {
	clientOptions []genapi.ClientOption
}

func (o *clientOptions) add(option genapi.ClientOption) {
	if option != nil {
		o.clientOptions = append(o.clientOptions, option)
	}
}

func WithHTTPClient(doer HttpRequestDoer) Option {
	return func(o *clientOptions) {
		o.add(genapi.WithHTTPClient(doer))
	}
}

func WithRequestEditor(fn RequestEditorFn) Option {
	return func(o *clientOptions) {
		o.add(genapi.WithRequestEditorFn(fn))
	}
}

func WithToken(token string) Option {
	token = strings.TrimSpace(token)

	return func(o *clientOptions) {
		if token == "" {
			return
		}

		o.add(genapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			return nil
		}))
	}
}
