package shiper

import "context"

func (c *Client) ListProjects(ctx context.Context, params *ProjectListParams) (*ProjectListResponse, error) {
	return call(ctx, c, "ListProjects", func(ctx context.Context, raw *rawClient) (*ProjectListResponse, error) {
		return raw.ProjectListWithResponse(ctx, params)
	})
}

func (c *Client) GetProject(ctx context.Context, projectID string) (*ProjectGetResponse, error) {
	return call(ctx, c, "GetProject", func(ctx context.Context, raw *rawClient) (*ProjectGetResponse, error) {
		return raw.ProjectGetWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) CreateProject(ctx context.Context, body ProjectCreateRequest) (*ProjectCreateResponse, error) {
	return call(ctx, c, "CreateProject", func(ctx context.Context, raw *rawClient) (*ProjectCreateResponse, error) {
		return raw.ProjectCreateWithResponse(ctx, body)
	})
}

func (c *Client) UpdateProject(ctx context.Context, projectID string, body ProjectUpdateRequest) (*ProjectUpdateResponse, error) {
	return call(ctx, c, "UpdateProject", func(ctx context.Context, raw *rawClient) (*ProjectUpdateResponse, error) {
		return raw.ProjectUpdateWithResponse(ctx, projectID, body)
	}, "projectID", projectID)
}

func (c *Client) DeleteProject(ctx context.Context, projectID string) (*ProjectDeleteResponse, error) {
	return call(ctx, c, "DeleteProject", func(ctx context.Context, raw *rawClient) (*ProjectDeleteResponse, error) {
		return raw.ProjectDeleteWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) ListProjectACLs(ctx context.Context, projectID string) (*ProjectAclListResponse, error) {
	return call(ctx, c, "ListProjectACLs", func(ctx context.Context, raw *rawClient) (*ProjectAclListResponse, error) {
		return raw.ProjectAclListWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) GetProjectACL(ctx context.Context, projectID string, userID string) (*ProjectAclGetResponse, error) {
	return call(ctx, c, "GetProjectACL", func(ctx context.Context, raw *rawClient) (*ProjectAclGetResponse, error) {
		return raw.ProjectAclGetWithResponse(ctx, projectID, userID)
	}, "projectID", projectID, "userID", userID)
}

func (c *Client) CreateProjectACL(ctx context.Context, projectID string, body ProjectAclCreateRequest) (*ProjectAclCreateResponse, error) {
	return call(ctx, c, "CreateProjectACL", func(ctx context.Context, raw *rawClient) (*ProjectAclCreateResponse, error) {
		return raw.ProjectAclCreateWithResponse(ctx, projectID, body)
	}, "projectID", projectID)
}

func (c *Client) UpdateProjectACL(ctx context.Context, projectID string, userID string, body ProjectAclUpdateRequest) (*ProjectAclUpdateResponse, error) {
	return call(ctx, c, "UpdateProjectACL", func(ctx context.Context, raw *rawClient) (*ProjectAclUpdateResponse, error) {
		return raw.ProjectAclUpdateWithResponse(ctx, projectID, userID, body)
	}, "projectID", projectID, "userID", userID)
}

func (c *Client) DeleteProjectACL(ctx context.Context, projectID string, userID string) (*ProjectAclDeleteResponse, error) {
	return call(ctx, c, "DeleteProjectACL", func(ctx context.Context, raw *rawClient) (*ProjectAclDeleteResponse, error) {
		return raw.ProjectAclDeleteWithResponse(ctx, projectID, userID)
	}, "projectID", projectID, "userID", userID)
}

func (c *Client) GetResourceLimits(ctx context.Context, projectID string) (*ResourcesLimitsResponse, error) {
	return call(ctx, c, "GetResourceLimits", func(ctx context.Context, raw *rawClient) (*ResourcesLimitsResponse, error) {
		return raw.ResourcesLimitsWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) GetResourceRecommendations(ctx context.Context, projectID string) (*ResourcesRecommendationsResponse, error) {
	return call(ctx, c, "GetResourceRecommendations", func(ctx context.Context, raw *rawClient) (*ResourcesRecommendationsResponse, error) {
		return raw.ResourcesRecommendationsWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) GetResourceUsage(ctx context.Context, projectID string) (*ResourcesUsageResponse, error) {
	return call(ctx, c, "GetResourceUsage", func(ctx context.Context, raw *rawClient) (*ResourcesUsageResponse, error) {
		return raw.ResourcesUsageWithResponse(ctx, projectID)
	}, "projectID", projectID)
}
