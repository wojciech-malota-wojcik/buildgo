package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:     "go",
		Version:  "1.22.5",
		IsGlobal: true,
		URL:      "https://go.dev/dl/go1.22.5.linux-amd64.tar.gz",
		Hash:     "sha256:904b924d435eaea086515bc63235b192ea441bd8c9b198c507e85009e6e4c7f0",
		Binaries: map[string]string{
			"go":    "go/bin/go",
			"gofmt": "go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:     "golangci",
		Version:  "1.59.1",
		IsGlobal: true,
		URL:      "https://github.com/golangci/golangci-lint/releases/download/v1.59.1/golangci-lint-1.59.1-linux-amd64.tar.gz",
		Hash:     "sha256:c30696f1292cff8778a495400745f0f9c0406a3f38d8bb12cef48d599f6c7791",
		Binaries: map[string]string{
			"golangci-lint": "golangci-lint-1.59.1-linux-amd64/golangci-lint",
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
