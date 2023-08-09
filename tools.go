package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:     "go",
		Version:  "1.21.0",
		IsGlobal: true,
		URL:      "https://go.dev/dl/go1.21.0.linux-amd64.tar.gz",
		Hash:     "sha256:d0398903a16ba2232b389fb31032ddf57cac34efda306a0eebac34f0965a0742",
		Binaries: map[string]string{
			"go":    "go/bin/go",
			"gofmt": "go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:     "golangci",
		Version:  "1.54.0",
		IsGlobal: true,
		URL:      "https://github.com/golangci/golangci-lint/releases/download/v1.54.0/golangci-lint-1.54.0-linux-amd64.tar.gz",
		Hash:     "sha256:a694f19dbfab3ea4d3956cb105d2e74c1dc49cb4c06ece903a3c534bce86b3dc",
		Binaries: map[string]string{
			"golangci-lint": "golangci-lint-1.54.0-linux-amd64/golangci-lint",
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
