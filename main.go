package tools

import (
	"github.com/outofforest/build/v2"
	"github.com/outofforest/tools/pkg/tools"
)

// Main is the entrypoint for builders.
func Main() {
	build.Main("outofforest", tools.Version())
}
