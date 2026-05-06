package shiper

// DeploymentState is the lifecycle state of a deployment.
type DeploymentState = string

const DeploymentStateSuccess DeploymentState = "success"

// BuildParameter is a key/value build-time parameter for a project.
type BuildParameter struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

// Deployment represents a single deployment of a project environment.
type Deployment struct {
	CommitSha          string        `json:"commitSha"`
	EnvironmentId      string        `json:"environmentId"`
	GithubDeploymentId string        `json:"githubDeploymentId"`
	Id                 string        `json:"id"`
	State              DeploymentState `json:"state"`
}

// Domain is a custom domain attached to a project environment.
type Domain struct {
	Name string `json:"name"`
}

// Environment is a deployment environment (production or preview) within a project.
type Environment struct {
	Branch     string `json:"branch"`
	Enabled    bool   `json:"enabled"`
	Id         string `json:"id"`
	Production bool   `json:"production"`
	ProjectId  string `json:"projectId"`
}

// Project is a Shiper project.
type Project struct {
	BasePath                *string          `json:"basePath,omitempty"`
	BuildParameters         *[]BuildParameter `json:"buildParameters,omitempty"`
	DisplayName             string           `json:"displayName"`
	Enabled                 bool             `json:"enabled"`
	Framework               string           `json:"framework"`
	GithubInstallationId    *int64           `json:"githubInstallationId,omitempty"`
	Port                    *int32           `json:"port,omitempty"`
	PreviewResourceLimit    ResourceLimit    `json:"previewResourceLimit"`
	ProductionResourceLimit ResourceLimit    `json:"productionResourceLimit"`
	ProjectId               string           `json:"projectId"`
	Replicas                *int32           `json:"replicas,omitempty"`
	Repository              string           `json:"repository"`
	TargetType              *string          `json:"targetType,omitempty"`
	UserId                  string           `json:"userId"`
}

// ResourceLimit specifies CPU and memory limits for a project environment.
type ResourceLimit struct {
	Cpu    float32 `json:"cpu"`
	Memory float32 `json:"memory"`
}

// Variable is a project-level environment variable with environment scoping.
type Variable struct {
	EnvironmentId string `json:"environmentId"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}
