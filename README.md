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
}
```

## Maintainers: Add Endpoint

1. Regenerate `internal/genapi/client.gen.go` (`go generate ./...`).
2. Add a thin `*Client` method in the matching resource file (`projects.go`, `users.go`, `variables.go`).
3. Follow wrapper rules: validate required IDs, call `raw.<Endpoint>WithResponse`, enforce status with `requireStatus`/`require2xx`, return `newAPIError` on unexpected status.
4. Add aliases in `types.go` if new request/response types are needed.
5. Add/extend a focused test for success and non-2xx behavior.

## Release a new version

1. Download api-1.json from https://api.shiper.com/v2
2. run `openapi-down-convert --input "api-1.json" --output spec-3.0.yaml`
3. manually edit spec-3.0.yaml to remove the "Null" from all oneOf schemas.
4. run `go generate ./...`
5. commit the generated code and handwritten wrapper changes and push to main
6. Create a new release with the tag name `vX.Y.Z` where X.Y.Z is the new version number.
