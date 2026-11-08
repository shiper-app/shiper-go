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

	log.Printf("environments: %d", len(*envs.JSON200))
}
```

## Maintainers: Add Endpoint

1. Regenerate `internal/genapi/client.gen.go` (`go generate ./...`).
2. Add a thin `*Client` method in the matching resource file (`projects.go`, `users.go`, `variables.go`, `environments.go`, `frameworks.go`).
3. Implement it as a single `call(...)` (see `client.go`): pass the operation name, a closure calling `raw.<Endpoint>WithResponse`, and the required IDs as name/value pairs. `call` validates the IDs, enforces a 2xx status and a decoded JSON body, and returns `newAPIError` otherwise.
4. Add aliases in `types.go` if new request/response types are needed.
5. Add/extend a focused test for success and non-2xx behavior.

## Release a new version

1. Download `api-1.json` from https://api.shiper.app/v2.
2. Regenerate the client: `go generate ./...`.
3. Commit the generated code and any handwritten wrapper changes and push to main.
4. Create a new release with the tag name `vX.Y.Z` where X.Y.Z is the new version number.
