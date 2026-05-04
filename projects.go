package shiper

import (
	"context"
	"net/http"
)

func (c *Client) ListProjects(ctx context.Context, params *ProjectListParams) (*ProjectListResponse, error) {
	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectListWithResponse(normalizeContext(ctx), params)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ListProjects", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListProjects", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) GetProject(ctx context.Context, projectID string) (*ProjectGetResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectGetWithResponse(normalizeContext(ctx), projectID)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("GetProject", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("GetProject", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) CreateProject(ctx context.Context, body ProjectCreateRequest) (*ProjectCreateResponse, error) {
	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectCreateWithResponse(normalizeContext(ctx), body)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("CreateProject", resp.StatusCode(), resp.Body, http.StatusCreated); err != nil {
		return nil, err
	}
	if resp.JSON201 == nil {
		return nil, newAPIError("CreateProject", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) UpdateProject(ctx context.Context, projectID string, body ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectUpdateWithResponse(normalizeContext(ctx), projectID, body)
	if err != nil {
		return nil, err
	}
	if err := require2xx("UpdateProject", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteProject(ctx context.Context, projectID string) (*ProjectDeleteResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectDeleteWithResponse(normalizeContext(ctx), projectID)
	if err != nil {
		return nil, err
	}
	if err := require2xx("DeleteProject", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}
