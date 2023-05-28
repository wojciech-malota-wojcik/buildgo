package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:    "go",
		Version: "1.20.4",
		URL:     "https://go.dev/dl/go1.20.4.linux-amd64.tar.gz",
		Hash:    "sha256:698ef3243972a51ddb4028e4a1ac63dc6d60821bf18e59a807e051fee0a385bd",
		Binaries: map[string]string{
			"tools/go":    "go/bin/go",
			"tools/gofmt": "go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:    "golangci",
		Version: "1.52.2",
		URL:     "https://github.com/golangci/golangci-lint/releases/download/v1.52.2/golangci-lint-1.52.2-linux-amd64.tar.gz",
		Hash:    "sha256:c9cf72d12058a131746edd409ed94ccd578fbd178899d1ed41ceae3ce5f54501",
		Binaries: map[string]string{
			"tools/golangci-lint": "golangci-lint-1.52.2-linux-amd64/golangci-lint",
		},
	},
}

// InstallAll installs all go tools
func InstallAll(ctx context.Context) error {
	return build.InstallTools(ctx, tools)
}

// EnsureGo ensures that go is installed
func EnsureGo(ctx context.Context) error {
	return build.EnsureTool(ctx, tools["go"])
}

// EnsureProtoC ensures that protoc is installed
func EnsureProtoC(ctx context.Context) error {
	return build.EnsureTool(ctx, tools["protoc"])
}

// EnsureGoProto ensures that go proto generator is installed
func EnsureGoProto(ctx context.Context, deps build.DepsFunc) error {
	deps(EnsureProtoC)

	return build.EnsureTool(ctx, tools["protoc-gen-go"])
}

// EnsureGolangCI ensures that golangci is installed
func EnsureGolangCI(ctx context.Context) error {
	return build.EnsureTool(ctx, tools["golangci"])
}
