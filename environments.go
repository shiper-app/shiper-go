package shiper

import (
	"context"
	"net/http"
)

func (c *Client) ListEnvironments(ctx context.Context, projectID string) (*ProjectEnvironmentsListResponse, error) {
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
	if err := requireStatus("ListEnvironments", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListEnvironments", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) PatchEnvironment(ctx context.Context, projectID string, environmentID string, body ProjectEnvironmentUpdateRequest) (*ProjectEnvironmentsUpdateResponse, error) {
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
	if err := require2xx("PatchEnvironment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListDeployments(ctx context.Context, projectID string, environmentID string, params *ProjectEnvironmentDeploymentsListParams) (*ProjectEnvironmentsDeploymentsListResponse, error) {
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
	if err := requireStatus("ListDeployments", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListDeployments", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) CancelDeployment(ctx context.Context, projectID string, environmentID string, deploymentID string) (*ProjectEnvironmentsDeploymentsCancelResponse, error) {
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
	if err := require2xx("CancelDeployment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) PatchDeployment(ctx context.Context, projectID string, environmentID string, deploymentID string, body ProjectEnvironmentDeploymentRedeployRequest) (*ProjectEnvironmentsDeploymentsRedeployResponse, error) {
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
	if err := require2xx("PatchDeployment", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListDomains(ctx context.Context, projectID string, environmentID string) (*ProjectEnvironmentsDomainsListResponse, error) {
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
	if err := requireStatus("ListDomains", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ListDomains", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}

func (c *Client) AddDomain(ctx context.Context, projectID string, environmentID string, body ProjectEnvironmentDomainAddRequest) (*ProjectEnvironmentsDomainsAddResponse, error) {
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
	if err := require2xx("AddDomain", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteDomain(ctx context.Context, projectID string, environmentID string, domainName string) (*ProjectEnvironmentsDomainsDeleteResponse, error) {
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
	if err := require2xx("DeleteDomain", resp.StatusCode(), resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ValidateDomainDNS(ctx context.Context, projectID string, environmentID string, domainName string) (*ProjectEnvironmentsDomainsValidateDNSResponse, error) {
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
	if err := requireStatus("ValidateDomainDNS", resp.StatusCode(), resp.Body, http.StatusOK); err != nil {
		return nil, err
	}
	if resp.JSON200 == nil {
		return nil, newAPIError("ValidateDomainDNS", resp.StatusCode(), resp.Body)
	}

	return resp, nil
}
