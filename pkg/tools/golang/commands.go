package golang

import "github.com/outofforest/build/v2/pkg/types"

// Commands is a set of commands useful for any go environment.
var Commands = map[string]types.Command{
	"generate/go": {
		Description: "Runs generators in go code",
		Fn:          Generate,
	},
	"lint/go": {
		Description: "Lints go code",
		Fn:          Lint,
	},
	"test/go": {
		Description: "Runs go unit tests",
		Fn:          UnitTests,
	},
	"tidy/go": {
		Description: "Tidies up the go code",
		Fn:          Tidy,
	},
}
