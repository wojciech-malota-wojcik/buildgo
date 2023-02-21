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
		Binaries: []string{
			"go/bin/go",
			"go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:    "golangci",
		Version: "1.51.2",
		URL:     "https://github.com/golangci/golangci-lint/releases/download/v1.51.2/golangci-lint-1.51.2-linux-amd64.tar.gz",
		Hash:    "sha256:4de479eb9d9bc29da51aec1834e7c255b333723d38dbd56781c68e5dddc6a90b",
		Binaries: []string{
			"golangci-lint-1.51.2-linux-amd64/golangci-lint",
		},
	},

	// https://github.com/protocolbuffers/protobuf/releases/
	"protoc": {
		Name:    "protoc",
		Version: "21.9",
		URL:     "https://github.com/protocolbuffers/protobuf/releases/download/v21.9/protoc-21.9-linux-x86_64.zip",
		Hash:    "sha256:3cd951aff8ce713b94cde55e12378f505f2b89d47bf080508cf77e3934f680b6",
		Binaries: []string{
			"bin/protoc",
		},
	},

	// https://github.com/protocolbuffers/protobuf-go/releases
	"protoc-gen-go": {
		Name:    "protoc-gen-go",
		Version: "1.28.1",
		URL:     "https://github.com/protocolbuffers/protobuf-go/releases/download/v1.28.1/protoc-gen-go.v1.28.1.linux.amd64.tar.gz",
		Hash:    "sha256:5c5802081fb9998c26cdfe607017a677c3ceaa19aae7895dbb1eef9518ebcb7f",
		Binaries: []string{
			"protoc-gen-go",
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
