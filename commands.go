package buildgo

// AddCommands adds go and git commands
func AddCommands(commands map[string]interface{}) {
	commands["git/fetch"] = GitFetch
	commands["dev/lint"] = GoLint
	commands["dev/tidy"] = GoModTidy
	commands["dev/test"] = GoTest
}
