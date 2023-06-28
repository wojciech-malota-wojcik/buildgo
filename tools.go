package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:     "go",
		Version:  "1.20.5",
		IsGlobal: true,
		URL:      "https://go.dev/dl/go1.20.5.linux-amd64.tar.gz",
		Hash:     "sha256:d7ec48cde0d3d2be2c69203bc3e0a44de8660b9c09a6e85c4732a3f7dc442612",
		Binaries: map[string]string{
			"go":    "go/bin/go",
			"gofmt": "go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:     "golangci",
		Version:  "1.53.3",
		IsGlobal: true,
		URL:      "https://github.com/golangci/golangci-lint/releases/download/v1.53.3/golangci-lint-1.53.3-linux-amd64.tar.gz",
		Hash:     "sha256:4f62007ca96372ccba54760e2ed39c2446b40ec24d9a90c21aad9f2fdf6cf0da",
		Binaries: map[string]string{
			"golangci-lint": "golangci-lint-1.53.3-linux-amd64/golangci-lint",
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
