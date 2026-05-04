# Shiper Go

Go SDK for the Shiper API

## Release a new version

1. Download api-1.json from https://api.shiper.com/v2
2. run `openapi-down-convert --input "api-1.json" --output spec-3.0.yaml`
3. manually edit spec-3.0.yaml to remove the "Null" from all oneOf schemas.
4. run `oapi-codegen --config oapi-codegen.yaml spec-3.0.yaml`
5. commit the generated code and push to main
6. Create a new release with the tag name `vX.Y.Z` where X.Y.Z is the new version number.
