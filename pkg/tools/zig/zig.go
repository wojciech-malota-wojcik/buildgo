package zig

import (
	"context"
	_ "embed"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/outofforest/build/v2/pkg/tools"
	"github.com/outofforest/build/v2/pkg/types"
	"github.com/outofforest/libexec"
	"github.com/outofforest/logger"
)

// BuildConfig is the configuration for building binaries.
type BuildConfig struct {
	// PackagePath is the path to package to build.
	PackagePath string

	// Step to build.
	Step string

	// OutputPath is the path for compiled artifacts.
	OutputPath string
}

// Build builds go binary.
func Build(ctx context.Context, deps types.DepsFunc, config BuildConfig) error {
	deps(EnsureZig)

	cacheDir := filepath.Join(tools.DevDir(ctx), "zig", "cache", "build")
	outputPath, err := filepath.Abs(config.OutputPath)
	if err != nil {
		return errors.WithStack(err)
	}

	args := []string{"build"}
	if config.Step != "" {
		args = append(args, config.Step)
	}
	args = append(args,
		"--prefix", outputPath,
		"--prefix-lib-dir", outputPath,
		"--prefix-exe-dir", outputPath,
		"--cache-dir", cacheDir,
		"--global-cache-dir", cacheDir,
		"--summary", "all",
	)

	cmd := exec.Command(tools.Bin(ctx, "bin/zig", tools.PlatformLocal), args...)
	cmd.Dir = config.PackagePath

	logger.Get(ctx).Info(
		"Building zig package",
		zap.String("package", config.PackagePath),
		zap.String("output", config.OutputPath),
		zap.String("command", cmd.String()),
	)
	if err := libexec.Exec(ctx, cmd); err != nil {
		return errors.Wrapf(err, "building zig package '%s' failed", config.PackagePath)
	}
	return nil
}
