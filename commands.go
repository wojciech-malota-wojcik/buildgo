package buildgo

import "github.com/outofforest/build"

// AddCommands adds go and git commands
func AddCommands(commands map[string]build.Command) {
	commands["git/fetch"] = build.Command{Fn: GitFetch, Description: "Fetches changes from repository"}
	commands["dev/lint"] = build.Command{Fn: GoLint, Description: "Lints go code"}
	commands["dev/tidy"] = build.Command{Fn: GoModTidy, Description: "Runs go mod tidy"}
	commands["dev/test"] = build.Command{Fn: GoTest, Description: "Runs go unit tests"}
}
