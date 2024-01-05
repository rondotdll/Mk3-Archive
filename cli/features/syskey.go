package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "syskey",
		Description: "Grabs the OS' activation key & key type",
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
