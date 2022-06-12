package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

var tools = map[string]build.Tool{
	"go": {
		Name:    "go",
		Version: "1.18.3",
		URL:     "https://go.dev/dl/go1.18.3.linux-amd64.tar.gz",
		Hash:    "sha256:956f8507b302ab0bb747613695cdae10af99bbd39a90cae522b7c0302cc27245",
		Binaries: []string{
			"go/bin/go",
			"go/bin/gofmt",
		},
	},
	"golangci": {
		Name:    "golangci",
		Version: "1.46.2",
		URL:     "https://github.com/golangci/golangci-lint/releases/download/v1.46.2/golangci-lint-1.46.2-linux-amd64.tar.gz",
		Hash:    "sha256:242cd4f2d6ac0556e315192e8555784d13da5d1874e51304711570769c4f2b9b",
		Binaries: []string{
			"golangci-lint-1.46.2-linux-amd64/golangci-lint",
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
