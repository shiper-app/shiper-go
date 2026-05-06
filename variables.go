package shiper

import (
	"context"
	"net/http"
)

func (c *Client) ListVariables(ctx context.Context, projectID string, params *VariablesListParams) (*VariablesListResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectVariablesListWithResponse(normalizeContext(ctx), projectID, params)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ListVariables", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListVariables", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) UpsertVariables(ctx context.Context, projectID string, body VariablesUpsertRequest) (*VariablesUpsertResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectVariablesUpsertWithResponse(normalizeContext(ctx), projectID, body)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("UpsertVariables", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("UpsertVariables", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) DeleteVariables(ctx context.Context, projectID string, body VariablesDeleteRequest) (*VariablesDeleteResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectVariablesDeleteWithResponse(normalizeContext(ctx), projectID, body)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("DeleteVariables", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("DeleteVariables", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}
