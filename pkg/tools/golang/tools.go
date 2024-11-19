package golang

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/outofforest/build/v2/pkg/tools"
	"github.com/outofforest/build/v2/pkg/types"
	"github.com/outofforest/libexec"
	"github.com/outofforest/logger"
)

// Tool names.
const (
	Go        tools.Name = "go"
	GolangCI  tools.Name = "golangci"
	LibEVMOne tools.Name = "libevmone"
)

var t = []tools.Tool{
	// https://go.dev/dl/
	tools.BinaryTool{
		Name:    Go,
		Version: "1.23.3",
		Sources: tools.Sources{
			tools.PlatformLinuxAMD64: {
				URL:  "https://go.dev/dl/go1.23.3.linux-amd64.tar.gz",
				Hash: "sha256:a0afb9744c00648bafb1b90b4aba5bdb86f424f02f9275399ce0c20b93a2c3a8",
				Links: map[string]string{
					"bin/go":    "go/bin/go",
					"bin/gofmt": "go/bin/gofmt",
				},
			},
			tools.PlatformDarwinAMD64: {
				URL:  "https://go.dev/dl/go1.23.3.darwin-amd64.tar.gz",
				Hash: "sha256:c7e024d5c0bc81845070f23598caf02f05b8ae88fd4ad2cd3e236ddbea833ad2",
				Links: map[string]string{
					"bin/go":    "go/bin/go",
					"bin/gofmt": "go/bin/gofmt",
				},
			},
			tools.PlatformDarwinARM64: {
				URL:  "https://go.dev/dl/go1.23.3.darwin-arm64.tar.gz",
				Hash: "sha256:31e119fe9bde6e105407a32558d5b5fa6ca11e2bd17f8b7b2f8a06aba16a0632",
				Links: map[string]string{
					"bin/go":    "go/bin/go",
					"bin/gofmt": "go/bin/gofmt",
				},
			},
		},
	},

	// https://github.com/golangci/golangci-lint/releases/
	tools.BinaryTool{
		Name:    GolangCI,
		Version: "1.62.0",
		Sources: tools.Sources{
			tools.PlatformLinuxAMD64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.0/golangci-lint-1.62.0-linux-amd64.tar.gz",
				Hash: "sha256:53695531eeb824b6883c703335cef6f07882f8ba6fedc00ed43853ea07fa1fbd",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.0-linux-amd64/golangci-lint",
				},
			},
			tools.PlatformDarwinAMD64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.0/golangci-lint-1.62.0-darwin-amd64.tar.gz", //nolint:lll // breaking down urls is not beneficial
				Hash: "sha256:0ed6f1a216ddb62e293858196799608d63894bd2ec178114484363ca45cde84b",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.0-darwin-amd64/golangci-lint",
				},
			},
			tools.PlatformDarwinARM64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.0/golangci-lint-1.62.0-darwin-arm64.tar.gz", //nolint:lll // breaking down urls is not beneficial
				Hash: "sha256:dde51958f0f24d442062b5709b6912d91e235115dfe5887e80b3e5602c9cc09b",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.0-darwin-arm64/golangci-lint",
				},
			},
		},
	},

	// https://github.com/ethereum/evmone/releases
	tools.BinaryTool{
		Name:    LibEVMOne,
		Version: "0.12.0",
		Sources: tools.Sources{
			tools.PlatformDockerAMD64: {
				URL:  "https://github.com/ethereum/evmone/releases/download/v0.12.0/evmone-0.12.0-linux-x86_64.tar.gz",
				Hash: "sha256:1c7b5eba0c8c3b3b2a7a05101e2d01a13a2f84b323989a29be66285dba4136ce",
				Links: map[string]string{
					"lib/libevmone.so": "lib/libevmone.so",
				},
			},
		},
	},
}

// GoPackageTool is the tool installed using go install command.
type GoPackageTool struct {
	Name    tools.Name
	Version string
	Package string
}

// GetName returns the name of the tool.
func (gpt GoPackageTool) GetName() tools.Name {
	return gpt.Name
}

// GetVersion returns the version of the tool.
func (gpt GoPackageTool) GetVersion() string {
	return gpt.Version
}

// IsCompatible tells if tool is defined for the platform.
func (gpt GoPackageTool) IsCompatible(platform tools.Platform) (bool, error) {
	golang, err := tools.Get(Go)
	if err != nil {
		return false, err
	}
	return golang.IsCompatible(platform)
}

// Verify verifies the cheksums.
func (gpt GoPackageTool) Verify(ctx context.Context) ([]error, error) {
	return nil, nil
}

// Ensure ensures that tool is installed.
func (gpt GoPackageTool) Ensure(ctx context.Context, platform tools.Platform) error {
	binName := filepath.Base(gpt.Package)
	downloadDir := tools.ToolDownloadDir(ctx, platform, gpt)
	dst := filepath.Join("bin", binName)

	//nolint:nestif // complexity comes from trivial error-handling ifs.
	if tools.ShouldReinstall(ctx, platform, gpt, dst, binName) {
		if err := tools.Ensure(ctx, Go, platform); err != nil {
			return errors.Wrapf(err, "ensuring go failed")
		}

		cmd := exec.Command(tools.Bin(ctx, "bin/go", platform), "install", gpt.Package+"@"+gpt.Version)
		cmd.Env = append(env(ctx), "GOBIN="+downloadDir)

		if err := libexec.Exec(ctx, cmd); err != nil {
			return err
		}

		srcPath := filepath.Join(downloadDir, binName)

		binChecksum, err := tools.Checksum(srcPath)
		if err != nil {
			return err
		}

		linksDir := tools.ToolLinksDir(ctx, platform, gpt)
		dstPath := filepath.Join(linksDir, dst)
		dstPathChecksum := dstPath + ":" + binChecksum

		if err := os.Remove(dstPath); err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		if err := os.Remove(dstPathChecksum); err != nil && !os.IsNotExist(err) {
			return errors.WithStack(err)
		}

		if err := os.MkdirAll(filepath.Dir(dstPath), 0o700); err != nil {
			return errors.WithStack(err)
		}
		if err := os.Chmod(srcPath, 0o700); err != nil {
			return errors.WithStack(err)
		}
		srcLinkPath, err := filepath.Rel(filepath.Dir(dstPathChecksum), filepath.Join(downloadDir, binName))
		if err != nil {
			return errors.WithStack(err)
		}
		if err := os.Symlink(srcLinkPath, dstPathChecksum); err != nil {
			return errors.WithStack(err)
		}
		if err := os.Symlink(filepath.Base(dstPathChecksum), dstPath); err != nil {
			return errors.WithStack(err)
		}
		if _, err := filepath.EvalSymlinks(dstPath); err != nil {
			return errors.WithStack(err)
		}

		logger.Get(ctx).Info("Binary installed to path", zap.String("path", dstPath))
	}

	return tools.LinkFiles(ctx, platform, gpt, []string{dst})
}

// EnsureGo ensures that go is available.
func EnsureGo(ctx context.Context, _ types.DepsFunc) error {
	return tools.Ensure(ctx, Go, tools.PlatformLocal)
}

// EnsureGolangCI ensures that go linter is available.
func EnsureGolangCI(ctx context.Context, _ types.DepsFunc) error {
	return tools.Ensure(ctx, GolangCI, tools.PlatformLocal)
}

// EnsureLibEVMOne ensures that libevmone is available.
func EnsureLibEVMOne(ctx context.Context, _ types.DepsFunc) error {
	return tools.Ensure(ctx, LibEVMOne, tools.PlatformDockerAMD64)
}

func init() {
	tools.Add(t...)
}
