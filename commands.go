package buildgo

import (
	"context"

	"github.com/outofforest/build"
)

// AddCommands adds go and git commands
func AddCommands(commands map[string]build.Command) {
	commands["build/me"] = build.Command{Fn: rebuildMe, Description: "Rebuilds the building tool"}
	commands["git/fetch"] = build.Command{Fn: GitFetch, Description: "Fetches changes from repository"}
	commands["dev/lint"] = build.Command{Fn: GoLint, Description: "Lints go code"}
	commands["dev/tidy"] = build.Command{Fn: GoModTidy, Description: "Runs go mod tidy"}
	commands["dev/test"] = build.Command{Fn: func(ctx context.Context, deps build.DepsFunc) error {
		return GoTest(ctx, deps)
	}, Description: "Runs go unit tests"}
}
