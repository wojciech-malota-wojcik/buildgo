package buildgo

import (
	"context"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/outofforest/build"
	"github.com/outofforest/libexec"
	"github.com/outofforest/logger"
	"github.com/ridge/must"
	"go.uber.org/zap"
)

// GoBuildPkg builds go package
func GoBuildPkg(ctx context.Context, pkg, out string, cgo bool) error {
	logger.Get(ctx).Info("Building go package", zap.String("package", pkg), zap.String("binary", out))
	cmd := exec.Command("go", "build", "-trimpath", "-ldflags=-w -s", "-o", must.String(filepath.Abs(out)), ".")
	cmd.Dir = pkg
	if !cgo {
		cmd.Env = append([]string{"CGO_ENABLED=0"}, os.Environ()...)
	}
	return libexec.Exec(ctx, cmd)
}

// GoLint runs golangci linter, runs go mod tidy and checks that git tree is clean
func GoLint(ctx context.Context, deps build.DepsFunc) error {
	deps(EnsureGo, EnsureGolangCI)
	log := logger.Get(ctx)
	config := must.String(filepath.Abs("build/.golangci.yaml"))
	err := onModule(func(path string) error {
		log.Info("Running linter", zap.String("path", path))
		cmd := exec.Command("golangci-lint", "run", "--config", config)
		cmd.Dir = path
		return libexec.Exec(ctx, cmd)
	})
	if err != nil {
		return err
	}
	deps(GoModTidy, gitStatusClean)
	return nil
}

// GoTest runs go test
func GoTest(ctx context.Context, deps build.DepsFunc) error {
	deps(EnsureGo)
	log := logger.Get(ctx)
	return onModule(func(path string) error {
		log.Info("Running go tests", zap.String("path", path))
		cmd := exec.Command("go", "test", "-count=1", "-shuffle=on", "-race", "./...")
		cmd.Dir = path
		return libexec.Exec(ctx, cmd)
	})
}

// GoModTidy calls `go mod tidy`
func GoModTidy(ctx context.Context, deps build.DepsFunc) error {
	deps(EnsureGo)
	log := logger.Get(ctx)
	return onModule(func(path string) error {
		log.Info("Running go mod tidy", zap.String("path", path))
		cmd := exec.Command("go", "mod", "tidy")
		cmd.Dir = path
		return libexec.Exec(ctx, cmd)
	})
}

func onModule(fn func(path string) error) error {
	return filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || d.Name() != "go.mod" {
			return nil
		}
		return fn(filepath.Dir(path))
	})
}
