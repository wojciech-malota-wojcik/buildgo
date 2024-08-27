package tools

import (
	"runtime"
	"runtime/debug"
	"strings"
)

// Version returns the version of the module.
func Version() string {
	module := module()

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("reading build info failed")
	}

	for _, m := range append([]*debug.Module{&bi.Main}, bi.Deps...) {
		if m.Path != module {
			continue
		}
		if m.Replace != nil {
			m = m.Replace
		}

		if m.Version == "(devel)" {
			return "devel"
		}

		return m.Version
	}

	panic("impossible condition: module not found")
}

func module() string {
	_, file, _, _ := runtime.Caller(0)
	module := strings.Join(strings.Split(file, "/")[:3], "/")
	index := strings.Index(module, "@")
	if index > 0 {
		module = module[:index]
	}
	return module
}
