package shiper

import "github.com/shiper-app/shiper-go/internal/genapi"

// HTTP plumbing re-exports.
type RequestEditorFn = genapi.RequestEditorFn
type HttpRequestDoer = genapi.HttpRequestDoer

// Domain types. The spec only names User in components/schemas; every other
// payload is an inline struct reachable through the response types below.
type User = genapi.User

// Union types callers must name to build request bodies.
type ProjectCreateBuildParameterValue = genapi.ProjectCreateJSONBody_BuildParameters_Value
type ProjectUpdateBuildParameterValue = genapi.ProjectUpdateJSONBody_BuildParameters_Value
type VariableEnvironmentId = genapi.VariablesUpsertJSONBody_Variables_EnvironmentId

// Request/parameter types.
type ProjectListParams = genapi.ProjectListParams
type ProjectCreateRequest = genapi.ProjectCreateJSONRequestBody
type ProjectUpdateRequest = genapi.ProjectUpdateJSONRequestBody
type ProjectAclCreateRequest = genapi.ProjectAclCreateJSONRequestBody
type ProjectAclUpdateRequest = genapi.ProjectAclUpdateJSONRequestBody
type EnvironmentHealthParams = genapi.EnvironmentHealthParams
type EnvironmentSetEnabledRequest = genapi.EnvironmentSetEnabledJSONRequestBody
type DeploymentListOwnParams = genapi.DeploymentListOwnParams
type DeploymentListParams = genapi.DeploymentListParams
type DeploymentRedeployRequest = genapi.DeploymentRedeployJSONRequestBody
type DomainAddRequest = genapi.DomainAddJSONRequestBody
type AnalyticsRequestsParams = genapi.AnalyticsRequestsParams
type AnalyticsBrowsersParams = genapi.AnalyticsBrowsersParams
type AnalyticsRoutesParams = genapi.AnalyticsRoutesParams
type AnalyticsCpuParams = genapi.AnalyticsCpuParams
type AnalyticsMemoryParams = genapi.AnalyticsMemoryParams
type UserSearchParams = genapi.UserSearchParams
type VariablesListParams = genapi.VariablesListParams
type VariablesDeleteRequest = genapi.VariablesDeleteJSONRequestBody
type VariablesUpsertRequest = genapi.VariablesUpsertJSONRequestBody

// Response types.
type ProjectListResponse = genapi.ProjectListResponse
type ProjectGetResponse = genapi.ProjectGetResponse
type ProjectCreateResponse = genapi.ProjectCreateResponse
type ProjectUpdateResponse = genapi.ProjectUpdateResponse
type ProjectDeleteResponse = genapi.ProjectDeleteResponse
type ProjectAclListResponse = genapi.ProjectAclListResponse
type ProjectAclGetResponse = genapi.ProjectAclGetResponse
type ProjectAclCreateResponse = genapi.ProjectAclCreateResponse
type ProjectAclUpdateResponse = genapi.ProjectAclUpdateResponse
type ProjectAclDeleteResponse = genapi.ProjectAclDeleteResponse
type ResourcesLimitsResponse = genapi.ResourcesLimitsResponse
type ResourcesRecommendationsResponse = genapi.ResourcesRecommendationsResponse
type ResourcesUsageResponse = genapi.ResourcesUsageResponse

type FrameworkListResponse = genapi.FrameworkListResponse
type FrameworkGetResponse = genapi.FrameworkGetResponse

type EnvironmentListResponse = genapi.EnvironmentListResponse
type EnvironmentHealthResponse = genapi.EnvironmentHealthResponse
type EnvironmentSetEnabledResponse = genapi.EnvironmentSetEnabledResponse
type DeploymentListOwnResponse = genapi.DeploymentListOwnResponse
type DeploymentListResponse = genapi.DeploymentListResponse
type DeploymentCancelResponse = genapi.DeploymentCancelResponse
type DeploymentRedeployResponse = genapi.DeploymentRedeployResponse
type DomainListResponse = genapi.DomainListResponse
type DomainConfigResponse = genapi.DomainConfigResponse
type DomainAddResponse = genapi.DomainAddResponse
type DomainDeleteResponse = genapi.DomainDeleteResponse
type AnalyticsRequestsResponse = genapi.AnalyticsRequestsResponse
type AnalyticsBrowsersResponse = genapi.AnalyticsBrowsersResponse
type AnalyticsRoutesResponse = genapi.AnalyticsRoutesResponse
type AnalyticsCpuResponse = genapi.AnalyticsCpuResponse
type AnalyticsMemoryResponse = genapi.AnalyticsMemoryResponse

type UserMeResponse = genapi.UserMeResponse
type UserSearchResponse = genapi.UserSearchResponse

type VariablesListResponse = genapi.VariablesListResponse
type VariablesUpsertResponse = genapi.VariablesUpsertResponse
type VariablesDeleteResponse = genapi.VariablesDeleteResponse
