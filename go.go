package buildgo

import (
	"context"
	"os"
	"os/exec"

	"github.com/wojciech-malota-wojcik/build"
	"github.com/wojciech-malota-wojcik/libexec"
)

// GoBuildPkg builds go package
func GoBuildPkg(ctx context.Context, pkg, out string, cgo bool) error {
	cmd := exec.Command("go", "build", "-o", out, "./"+pkg)
	if !cgo {
		cmd.Env = append([]string{"CGO_ENABLED=0"}, os.Environ()...)
	}
	return libexec.Exec(ctx, cmd)
}

// GoLint runs golangci linter, runs go mod tidy and checks that git tree is clean
func GoLint(ctx context.Context, deps build.DepsFunc) error {
	if err := libexec.Exec(ctx, exec.Command("golangci-lint", "run", "--config", "build/.golangci.yaml")); err != nil {
		return err
	}
	deps(goModTidy, gitStatusClean)
	return nil
}

// GoImports run goimports
func GoImports(ctx context.Context) error {
	return libexec.Exec(ctx, exec.Command("goimports", "-w", "."))
}

// GoTest runs go test
func GoTest(ctx context.Context) error {
	return libexec.Exec(ctx, exec.Command("go", "test", "-count=1", "-race", "./..."))
}

func goModTidy(ctx context.Context) error {
	return libexec.Exec(ctx, exec.Command("go", "mod", "tidy"))
}
