package shiper

import "context"

func (c *Client) ListFrameworks(ctx context.Context) (*FrameworkListResponse, error) {
	return call(ctx, c, "ListFrameworks", func(ctx context.Context, raw *rawClient) (*FrameworkListResponse, error) {
		return raw.FrameworkListWithResponse(ctx)
	})
}

func (c *Client) GetFramework(ctx context.Context, frameworkID string) (*FrameworkGetResponse, error) {
	return call(ctx, c, "GetFramework", func(ctx context.Context, raw *rawClient) (*FrameworkGetResponse, error) {
		return raw.FrameworkGetWithResponse(ctx, frameworkID)
	}, "frameworkID", frameworkID)
}
