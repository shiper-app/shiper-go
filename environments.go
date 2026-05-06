package shiper

import (
	"context"
	"net/http"
)

func (c *Client) ListProjectEnvironments(ctx context.Context, projectID string) (*ProjectEnvironmentsListResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsListWithResponse(normalizeContext(ctx), projectID)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ListProjectEnvironments", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListProjectEnvironments", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) UpdateProjectEnvironment(ctx context.Context, projectID string, environmentID string, body ProjectEnvironmentUpdateRequest) (*ProjectEnvironmentsUpdateResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsUpdateWithResponse(normalizeContext(ctx), projectID, environmentID, body)
	if err != nil {
		return nil, err
	}
	if err := require2xx("UpdateProjectEnvironment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListProjectEnvironmentDeployments(ctx context.Context, projectID string, environmentID string, params *ProjectEnvironmentDeploymentsListParams) (*ProjectEnvironmentsDeploymentsListResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDeploymentsListWithResponse(normalizeContext(ctx), projectID, environmentID, params)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ListProjectEnvironmentDeployments", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListProjectEnvironmentDeployments", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) CancelProjectEnvironmentDeployment(ctx context.Context, projectID string, environmentID string, deploymentID string) (*ProjectEnvironmentsDeploymentsCancelResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("deploymentID", deploymentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDeploymentsCancelWithResponse(normalizeContext(ctx), projectID, environmentID, deploymentID)
	if err != nil {
		return nil, err
	}
	if err := require2xx("CancelProjectEnvironmentDeployment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RedeployProjectEnvironmentDeployment(ctx context.Context, projectID string, environmentID string, deploymentID string, body ProjectEnvironmentDeploymentRedeployRequest) (*ProjectEnvironmentsDeploymentsRedeployResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("deploymentID", deploymentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDeploymentsRedeployWithResponse(normalizeContext(ctx), projectID, environmentID, deploymentID, body)
	if err != nil {
		return nil, err
	}
	if err := require2xx("RedeployProjectEnvironmentDeployment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListProjectEnvironmentDomains(ctx context.Context, projectID string, environmentID string) (*ProjectEnvironmentsDomainsListResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDomainsListWithResponse(normalizeContext(ctx), projectID, environmentID)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ListProjectEnvironmentDomains", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListProjectEnvironmentDomains", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) AddProjectEnvironmentDomain(ctx context.Context, projectID string, environmentID string, body ProjectEnvironmentDomainAddRequest) (*ProjectEnvironmentsDomainsAddResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDomainsAddWithResponse(normalizeContext(ctx), projectID, environmentID, body)
	if err != nil {
		return nil, err
	}
	if err := require2xx("AddProjectEnvironmentDomain", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteProjectEnvironmentDomain(ctx context.Context, projectID string, environmentID string, domainName string) (*ProjectEnvironmentsDomainsDeleteResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("domainName", domainName); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDomainsDeleteWithResponse(normalizeContext(ctx), projectID, environmentID, domainName)
	if err != nil {
		return nil, err
	}
	if err := require2xx("DeleteProjectEnvironmentDomain", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ValidateProjectEnvironmentDomainDNS(ctx context.Context, projectID string, environmentID string, domainName string) (*ProjectEnvironmentsDomainsValidateDNSResponse, error) {
	if err := validateRequiredString("projectID", projectID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("environmentID", environmentID); err != nil {
		return nil, err
	}
	if err := validateRequiredString("domainName", domainName); err != nil {
		return nil, err
	}

	raw, err := c.getRawClient()
	if err != nil {
		return nil, err
	}

	resp, err := raw.ProjectEnvironmentsDomainsValidateDnsWithResponse(normalizeContext(ctx), projectID, environmentID, domainName)
	if err != nil {
		return nil, err
	}
	if err := requireStatus("ValidateProjectEnvironmentDomainDNS", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ValidateProjectEnvironmentDomainDNS", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}
