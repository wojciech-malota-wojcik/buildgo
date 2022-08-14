package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	// https://go.dev/dl/
	"go": {
		Name:    "go",
		Version: "1.19",
		URL:     "https://go.dev/dl/go1.19.linux-amd64.tar.gz",
		Hash:    "sha256:464b6b66591f6cf055bc5df90a9750bf5fbc9d038722bb84a9d56a2bea974be6",
		Binaries: []string{
			"go/bin/go",
			"go/bin/gofmt",
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	"golangci": {
		Name:    "golangci",
		Version: "1.48.0",
		URL:     "https://github.com/golangci/golangci-lint/releases/download/v1.48.0/golangci-lint-1.48.0-linux-amd64.tar.gz",
		Hash:    "sha256:127c5c9d47cf3a3cf4128815dea1d9623d57a83a22005e91b986b0cbceb09233",
		Binaries: []string{
			"golangci-lint-1.48.0-linux-amd64/golangci-lint",
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

// EnsureGolangCI ensures that golangci is installed
func EnsureGolangCI(ctx context.Context) error {
	return build.EnsureTool(ctx, tools["golangci"])
}
