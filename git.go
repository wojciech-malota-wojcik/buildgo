package buildgo

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/outofforest/build"
	"github.com/outofforest/libexec"
	"github.com/pkg/errors"
)

// GitFetch fetches changes from repo
func GitFetch(ctx context.Context, _ build.DepsFunc) error {
	return libexec.Exec(ctx, exec.Command("git", "fetch", "-p"))
}

func gitStatusClean(ctx context.Context, _ build.DepsFunc) error {
	buf := &bytes.Buffer{}
	cmd := exec.Command("git", "status", "-s")
	cmd.Stdout = buf
	if err := libexec.Exec(ctx, cmd); err != nil {
		return err
	}
	if buf.Len() > 0 {
		fmt.Println("git status:")
		fmt.Println(buf)
		return errors.New("git status is not empty")
	}
	return nil
}
