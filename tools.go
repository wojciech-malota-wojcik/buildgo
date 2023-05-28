package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:    "go",
		Version: "1.20.1",
		URL:     "https://go.dev/dl/go1.20.1.linux-amd64.tar.gz",
		Hash:    "sha256:000a5b1fca4f75895f78befeb2eecf10bfff3c428597f3f1e69133b63b911b02",
		Binaries: map[string]string{
			"tools/go":    "go/bin/go",
			"tools/gofmt": "go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:    "golangci",
		Version: "1.51.2",
		URL:     "https://github.com/golangci/golangci-lint/releases/download/v1.51.2/golangci-lint-1.51.2-linux-amd64.tar.gz",
		Hash:    "sha256:4de479eb9d9bc29da51aec1834e7c255b333723d38dbd56781c68e5dddc6a90b",
		Binaries: map[string]string{
			"tools/golangci-lint": "golangci-lint-1.51.2-linux-amd64/golangci-lint",
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
