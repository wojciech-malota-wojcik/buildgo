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
		Version: "1.23.4",
		Sources: tools.Sources{
			tools.PlatformLinuxAMD64: {
				URL:  "https://go.dev/dl/go1.23.4.linux-amd64.tar.gz",
				Hash: "sha256:6924efde5de86fe277676e929dc9917d466efa02fb934197bc2eba35d5680971",
				Links: map[string]string{
					"bin/go":    "go/bin/go",
					"bin/gofmt": "go/bin/gofmt",
				},
			},
			tools.PlatformDarwinAMD64: {
				URL:  "https://go.dev/dl/go1.23.4.darwin-amd64.tar.gz",
				Hash: "sha256:6700067389a53a1607d30aa8d6e01d198230397029faa0b109e89bc871ab5a0e",
				Links: map[string]string{
					"bin/go":    "go/bin/go",
					"bin/gofmt": "go/bin/gofmt",
				},
			},
			tools.PlatformDarwinARM64: {
				URL:  "https://go.dev/dl/go1.23.4.darwin-arm64.tar.gz",
				Hash: "sha256:87d2bb0ad4fe24d2a0685a55df321e0efe4296419a9b3de03369dbe60b8acd3a",
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
		Version: "1.62.2",
		Sources: tools.Sources{
			tools.PlatformLinuxAMD64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.2/golangci-lint-1.62.2-linux-amd64.tar.gz",
				Hash: "sha256:5101292b7925a6a14b49c5c3d845c5021399698ffd2f41bcfab8a111b5669939",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.2-linux-amd64/golangci-lint",
				},
			},
			tools.PlatformDarwinAMD64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.2/golangci-lint-1.62.2-darwin-amd64.tar.gz", //nolint:lll // breaking down urls is not beneficial
				Hash: "sha256:6c9ffd05896f0638d5c37152ac4ae337c2d301ba6c9dadf49c04e6d639f10f91",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.2-darwin-amd64/golangci-lint",
				},
			},
			tools.PlatformDarwinARM64: {
				URL:  "https://github.com/golangci/golangci-lint/releases/download/v1.62.2/golangci-lint-1.62.2-darwin-arm64.tar.gz", //nolint:lll // breaking down urls is not beneficial
				Hash: "sha256:6c76f54467ba471f7bdcd5df0f27c3fa3dbe530b771a10d384c3d8c7178f5e89",
				Links: map[string]string{
					"bin/golangci-lint": "golangci-lint-1.62.2-darwin-arm64/golangci-lint",
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
