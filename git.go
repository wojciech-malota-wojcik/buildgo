package buildgo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/wojciech-malota-wojcik/libexec"
)

// GitFetch fetches changes from repo
func GitFetch(ctx context.Context) error {
	return libexec.Exec(ctx, exec.Command("git", "fetch", "-p"))
}

func gitStatusClean(ctx context.Context) error {
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
