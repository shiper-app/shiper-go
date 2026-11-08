package shiper

import "context"

func (c *Client) ListVariables(ctx context.Context, projectID string, params *VariablesListParams) (*VariablesListResponse, error) {
	return call(ctx, c, "ListVariables", func(ctx context.Context, raw *rawClient) (*VariablesListResponse, error) {
		return raw.VariablesListWithResponse(ctx, projectID, params)
	}, "projectID", projectID)
}

func (c *Client) UpsertVariables(ctx context.Context, projectID string, body VariablesUpsertRequest) (*VariablesUpsertResponse, error) {
	return call(ctx, c, "UpsertVariables", func(ctx context.Context, raw *rawClient) (*VariablesUpsertResponse, error) {
		return raw.VariablesUpsertWithResponse(ctx, projectID, body)
	}, "projectID", projectID)
}

func (c *Client) DeleteVariables(ctx context.Context, projectID string, body VariablesDeleteRequest) (*VariablesDeleteResponse, error) {
	return call(ctx, c, "DeleteVariables", func(ctx context.Context, raw *rawClient) (*VariablesDeleteResponse, error) {
		return raw.VariablesDeleteWithResponse(ctx, projectID, body)
	}, "projectID", projectID)
}
