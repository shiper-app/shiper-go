package shiper

import (
	"context"
	"net/http"
)

func (c *Client) GetMe(ctx context.Context) (*UserMeResponse, error) {
	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.UserMeWithResponse(normalizeContext(ctx))
	if err != nil {
		return nil, err
	}
	if err := requireStatus("GetMe", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("GetMe", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) GetUser(ctx context.Context, userID string) (*UserGetResponse, error) {
	if err := validateRequiredString("userID", userID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.UserGetWithResponse(normalizeContext(ctx), userID)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("GetUser", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("GetUser", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) SearchUsers(ctx context.Context, params *UserSearchParams) (*UserSearchResponse, error) {
	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.UserSearchWithResponse(normalizeContext(ctx), params)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("SearchUsers", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("SearchUsers", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}
