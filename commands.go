package buildgo

// AddCommands adds go and git commands
func AddCommands(commands map[string]interface{}) {
	commands["git/fetch"] = GitFetch
	commands["dev/goimports"] = GoImports
	commands["dev/lint"] = GoLint
	commands["dev/test"] = GoTest
}
