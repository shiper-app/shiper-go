package shiper

import "github.com/shiper-app/shiper-go/internal/genapi"

// HTTP plumbing re-exports.
type RequestEditorFn = genapi.RequestEditorFn
type HttpRequestDoer = genapi.HttpRequestDoer

// Domain types (mirror components/schemas in the OpenAPI spec).
type Project = genapi.Project
type ProjectListItem = genapi.ProjectListItem
type ProjectWithCounts = genapi.ProjectWithCounts
type ResourceLimit = genapi.ResourceLimit
type BuildParameter = genapi.BuildParameter
type BuildParameterValue = genapi.BuildParameter_Value
type VariableEnvironmentId = genapi.Variable_EnvironmentId
type Environment = genapi.Environment
type Deployment = genapi.Deployment
type DeploymentState = genapi.DeploymentState
type Domain = genapi.Domain
type Variable = genapi.Variable
type Framework = genapi.Framework
type BuildOption = genapi.BuildOption
type ResourceTemplate = genapi.ResourceTemplate
type InitializationNotice = genapi.InitializationNotice
type ProjectAcl = genapi.ProjectAcl
type AclRole = genapi.AclRole
type User = genapi.User
type UserRole = genapi.UserRole
type QosPolicy = genapi.QosPolicy

// DeploymentState constants.
const (
	DeploymentStatePending    = genapi.DeploymentStatePending
	DeploymentStateSuccess    = genapi.DeploymentStateSuccess
	DeploymentStateFailure    = genapi.DeploymentStateFailure
	DeploymentStateInProgress = genapi.DeploymentStateInProgress
	DeploymentStateQueued     = genapi.DeploymentStateQueued
	DeploymentStateError      = genapi.DeploymentStateError
	DeploymentStateInactive   = genapi.DeploymentStateInactive
)

// Request/parameter types.
type ProjectListParams = genapi.ProjectListParams
type ProjectCreateRequest = genapi.ProjectCreateJSONRequestBody
type ProjectUpdateRequest = genapi.ProjectUpdateJSONRequestBody
type ProjectEnvironmentUpdateRequest = genapi.ProjectEnvironmentsUpdateJSONRequestBody
type ProjectEnvironmentDeploymentsListParams = genapi.ProjectEnvironmentsDeploymentsListParams
type ProjectEnvironmentDeploymentsPatchRequest = genapi.ProjectEnvironmentsDeploymentsPatchJSONRequestBody
type ProjectEnvironmentDeploymentsPatchResponse = genapi.ProjectEnvironmentsDeploymentsPatchResponse
type ProjectEnvironmentDeploymentRedeployRequest = genapi.ProjectEnvironmentsDeploymentsRedeployJSONRequestBody
type ProjectEnvironmentDomainAddRequest = genapi.ProjectEnvironmentsDomainsAddJSONRequestBody
type UserSearchParams = genapi.UserSearchParams
type VariablesListParams = genapi.ProjectVariablesListParams
type VariablesDeleteRequest = genapi.ProjectVariablesDeleteJSONRequestBody
type VariablesUpsertRequest = genapi.ProjectVariablesUpsertJSONRequestBody

// Response types.
type ProjectListResponse = genapi.ProjectListResponse
type ProjectGetResponse = genapi.ProjectGetResponse
type ProjectCreateResponse = genapi.ProjectCreateResponse
type ProjectUpdateResponse = genapi.ProjectUpdateResponse
type ProjectDeleteResponse = genapi.ProjectDeleteResponse
type ProjectEnvironmentsListResponse = genapi.ProjectEnvironmentsListResponse
type ProjectEnvironmentsUpdateResponse = genapi.ProjectEnvironmentsUpdateResponse
type ProjectEnvironmentsDeploymentsListResponse = genapi.ProjectEnvironmentsDeploymentsListResponse
type ProjectEnvironmentsDeploymentsCancelResponse = genapi.ProjectEnvironmentsDeploymentsCancelResponse
type ProjectEnvironmentsDeploymentsRedeployResponse = genapi.ProjectEnvironmentsDeploymentsRedeployResponse
type ProjectEnvironmentsDomainsListResponse = genapi.ProjectEnvironmentsDomainsListResponse
type ProjectEnvironmentsDomainsAddResponse = genapi.ProjectEnvironmentsDomainsAddResponse
type ProjectEnvironmentsDomainsDeleteResponse = genapi.ProjectEnvironmentsDomainsDeleteResponse
type ProjectEnvironmentsDomainsValidateDNSResponse = genapi.ProjectEnvironmentsDomainsValidateDnsResponse

type UserMeResponse = genapi.UserMeResponse
type UserGetResponse = genapi.UserGetResponse
type UserSearchResponse = genapi.UserSearchResponse

type VariablesListResponse = genapi.ProjectVariablesListResponse
type VariablesUpsertResponse = genapi.ProjectVariablesUpsertResponse
type VariablesDeleteResponse = genapi.ProjectVariablesDeleteResponse
