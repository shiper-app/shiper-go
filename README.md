# Shiper Go

Go SDK for the Shiper API.

## Usage

```go
package main

import (
    "context"
    "log"

    shiper "github.com/shiper-app/shiper-go"
)

func main() {
	client, err := shiper.NewClient(shiper.WithToken("YOUR_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

    resp, err := client.ListProjects(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

	log.Printf("projects: %d", len(resp.JSON200.Data))

	envs, err := client.ListEnvironments(context.Background(), "project-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("environments: %d", len(envs.JSON200.Environments))
}
```

## Maintainers: Add Endpoint

1. Regenerate `internal/genapi/client.gen.go` (`go generate ./...`).
2. Add a thin `*Client` method in the matching resource file (`projects.go`, `users.go`, `variables.go`, `environments.go`).
3. Follow wrapper rules: validate required IDs, call `raw.<Endpoint>WithResponse`, enforce status with `requireStatus`/`require2xx`, return `newAPIError` on unexpected status.
4. Add aliases in `types.go` if new request/response types are needed.
5. Add/extend a focused test for success and non-2xx behavior.

## Release a new version

1. Download `api-1.json` from https://api.shiper.app/v2.
2. Down-convert to OpenAPI 3.0: `npx @apiture/openapi-down-convert --input api-1.json --output spec-3.0.yaml`.
3. Patch the spec for oapi-codegen compatibility (collapses `anyOf: [X, null]` into `nullable: true` and sets `type: string` on bare string enums): `node hack/fix-nullables.mjs spec-3.0.yaml`. Run `cd hack && npm install` once to fetch the YAML dep.
4. Regenerate the client: `go generate ./...`.
5. Commit the generated code and any handwritten wrapper changes and push to main.
6. Create a new release with the tag name `vX.Y.Z` where X.Y.Z is the new version number.
