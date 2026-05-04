package shiper

import "github.com/shiper-app/shiper-go/internal/genapi"

type RequestEditorFn = genapi.RequestEditorFn
type HttpRequestDoer = genapi.HttpRequestDoer

type ProjectListParams = genapi.ProjectListParams
type ProjectCreateRequest = genapi.ProjectCreateJSONRequestBody
type ProjectUpdateRequest = genapi.ProjectUpdateJSONRequestBody
type UserSearchParams = genapi.UserSearchParams
type VariablesListParams = genapi.VariablesListParams
type VariablesDeleteRequest = genapi.VariablesDeleteJSONRequestBody
type VariablesUpsertRequest = genapi.VariablesUpsertJSONRequestBody

type ProjectListResponse = genapi.ProjectListResponse
type ProjectGetResponse = genapi.ProjectGetResponse
type ProjectCreateResponse = genapi.ProjectCreateResponse
type ProjectUpdateResponse = genapi.ProjectUpdateResponse
type ProjectDeleteResponse = genapi.ProjectDeleteResponse

type UserMeResponse = genapi.UserMeResponse
type UserGetResponse = genapi.UserGetResponse
type UserSearchResponse = genapi.UserSearchResponse

type VariablesListResponse = genapi.VariablesListResponse
type VariablesUpsertResponse = genapi.VariablesUpsertResponse
type VariablesDeleteResponse = genapi.VariablesDeleteResponse
