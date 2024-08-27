package rust

import "github.com/outofforest/build/v2/pkg/types"

// Commands is a set of commands useful for any rust environment.
var Commands = map[string]types.Command{
	"lint/rust": {
		Description: "Lints rust code",
		Fn:          Lint,
	},
	"test/rust": {
		Description: "Runs rust unit tests",
		Fn:          UnitTests,
	},
}
