package main

import (
	"github.com/outofforest/build/v2"
	"github.com/outofforest/build/v2/pkg/tools/git"
	tmain "github.com/outofforest/tools"
)

func main() {
	build.RegisterCommands(
		build.Commands,
		git.Commands,
		commands,
	)
	tmain.Main()
}
