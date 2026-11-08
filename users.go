package shiper

import "context"

func (c *Client) GetMe(ctx context.Context) (*UserMeResponse, error) {
	return call(ctx, c, "GetMe", func(ctx context.Context, raw *rawClient) (*UserMeResponse, error) {
		return raw.UserMeWithResponse(ctx)
	})
}

func (c *Client) SearchUsers(ctx context.Context, params *UserSearchParams) (*UserSearchResponse, error) {
	return call(ctx, c, "SearchUsers", func(ctx context.Context, raw *rawClient) (*UserSearchResponse, error) {
		return raw.UserSearchWithResponse(ctx, params)
	})
}
