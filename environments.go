package shiper

import "context"

func (c *Client) ListEnvironments(ctx context.Context, projectID string) (*EnvironmentListResponse, error) {
	return call(ctx, c, "ListEnvironments", func(ctx context.Context, raw *rawClient) (*EnvironmentListResponse, error) {
		return raw.EnvironmentListWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) GetEnvironmentHealth(ctx context.Context, projectID string, params *EnvironmentHealthParams) (*EnvironmentHealthResponse, error) {
	return call(ctx, c, "GetEnvironmentHealth", func(ctx context.Context, raw *rawClient) (*EnvironmentHealthResponse, error) {
		return raw.EnvironmentHealthWithResponse(ctx, projectID, params)
	}, "projectID", projectID)
}

func (c *Client) SetEnvironmentEnabled(ctx context.Context, projectID string, environmentID string, body EnvironmentSetEnabledRequest) (*EnvironmentSetEnabledResponse, error) {
	return call(ctx, c, "SetEnvironmentEnabled", func(ctx context.Context, raw *rawClient) (*EnvironmentSetEnabledResponse, error) {
		return raw.EnvironmentSetEnabledWithResponse(ctx, projectID, environmentID, body)
	}, "projectID", projectID, "environmentID", environmentID)
}

// ListOwnDeployments lists deployments across all projects of the current user.
func (c *Client) ListOwnDeployments(ctx context.Context, params *DeploymentListOwnParams) (*DeploymentListOwnResponse, error) {
	return call(ctx, c, "ListOwnDeployments", func(ctx context.Context, raw *rawClient) (*DeploymentListOwnResponse, error) {
		return raw.DeploymentListOwnWithResponse(ctx, params)
	})
}

func (c *Client) ListDeployments(ctx context.Context, projectID string, params *DeploymentListParams) (*DeploymentListResponse, error) {
	return call(ctx, c, "ListDeployments", func(ctx context.Context, raw *rawClient) (*DeploymentListResponse, error) {
		return raw.DeploymentListWithResponse(ctx, projectID, params)
	}, "projectID", projectID)
}

func (c *Client) CancelDeployment(ctx context.Context, projectID string, deploymentID string) (*DeploymentCancelResponse, error) {
	return call(ctx, c, "CancelDeployment", func(ctx context.Context, raw *rawClient) (*DeploymentCancelResponse, error) {
		return raw.DeploymentCancelWithResponse(ctx, projectID, deploymentID)
	}, "projectID", projectID, "deploymentID", deploymentID)
}

func (c *Client) RedeployDeployment(ctx context.Context, projectID string, deploymentID string, body DeploymentRedeployRequest) (*DeploymentRedeployResponse, error) {
	return call(ctx, c, "RedeployDeployment", func(ctx context.Context, raw *rawClient) (*DeploymentRedeployResponse, error) {
		return raw.DeploymentRedeployWithResponse(ctx, projectID, deploymentID, body)
	}, "projectID", projectID, "deploymentID", deploymentID)
}

func (c *Client) ListDomains(ctx context.Context, projectID string) (*DomainListResponse, error) {
	return call(ctx, c, "ListDomains", func(ctx context.Context, raw *rawClient) (*DomainListResponse, error) {
		return raw.DomainListWithResponse(ctx, projectID)
	}, "projectID", projectID)
}

func (c *Client) GetDomainConfig(ctx context.Context, projectID string, domain string) (*DomainConfigResponse, error) {
	return call(ctx, c, "GetDomainConfig", func(ctx context.Context, raw *rawClient) (*DomainConfigResponse, error) {
		return raw.DomainConfigWithResponse(ctx, projectID, domain)
	}, "projectID", projectID, "domain", domain)
}

func (c *Client) AddDomain(ctx context.Context, projectID string, environmentID string, body DomainAddRequest) (*DomainAddResponse, error) {
	return call(ctx, c, "AddDomain", func(ctx context.Context, raw *rawClient) (*DomainAddResponse, error) {
		return raw.DomainAddWithResponse(ctx, projectID, environmentID, body)
	}, "projectID", projectID, "environmentID", environmentID)
}

func (c *Client) DeleteDomain(ctx context.Context, projectID string, environmentID string, domain string) (*DomainDeleteResponse, error) {
	return call(ctx, c, "DeleteDomain", func(ctx context.Context, raw *rawClient) (*DomainDeleteResponse, error) {
		return raw.DomainDeleteWithResponse(ctx, projectID, environmentID, domain)
	}, "projectID", projectID, "environmentID", environmentID, "domain", domain)
}

func (c *Client) GetAnalyticsRequests(ctx context.Context, projectID string, environmentID string, params *AnalyticsRequestsParams) (*AnalyticsRequestsResponse, error) {
	return call(ctx, c, "GetAnalyticsRequests", func(ctx context.Context, raw *rawClient) (*AnalyticsRequestsResponse, error) {
		return raw.AnalyticsRequestsWithResponse(ctx, projectID, environmentID, params)
	}, "projectID", projectID, "environmentID", environmentID)
}

func (c *Client) GetAnalyticsBrowsers(ctx context.Context, projectID string, environmentID string, params *AnalyticsBrowsersParams) (*AnalyticsBrowsersResponse, error) {
	return call(ctx, c, "GetAnalyticsBrowsers", func(ctx context.Context, raw *rawClient) (*AnalyticsBrowsersResponse, error) {
		return raw.AnalyticsBrowsersWithResponse(ctx, projectID, environmentID, params)
	}, "projectID", projectID, "environmentID", environmentID)
}

func (c *Client) GetAnalyticsRoutes(ctx context.Context, projectID string, environmentID string, params *AnalyticsRoutesParams) (*AnalyticsRoutesResponse, error) {
	return call(ctx, c, "GetAnalyticsRoutes", func(ctx context.Context, raw *rawClient) (*AnalyticsRoutesResponse, error) {
		return raw.AnalyticsRoutesWithResponse(ctx, projectID, environmentID, params)
	}, "projectID", projectID, "environmentID", environmentID)
}

func (c *Client) GetAnalyticsCPU(ctx context.Context, projectID string, environmentID string, params *AnalyticsCpuParams) (*AnalyticsCpuResponse, error) {
	return call(ctx, c, "GetAnalyticsCPU", func(ctx context.Context, raw *rawClient) (*AnalyticsCpuResponse, error) {
		return raw.AnalyticsCpuWithResponse(ctx, projectID, environmentID, params)
	}, "projectID", projectID, "environmentID", environmentID)
}

func (c *Client) GetAnalyticsMemory(ctx context.Context, projectID string, environmentID string, params *AnalyticsMemoryParams) (*AnalyticsMemoryResponse, error) {
	return call(ctx, c, "GetAnalyticsMemory", func(ctx context.Context, raw *rawClient) (*AnalyticsMemoryResponse, error) {
		return raw.AnalyticsMemoryWithResponse(ctx, projectID, environmentID, params)
	}, "projectID", projectID, "environmentID", environmentID)
}
