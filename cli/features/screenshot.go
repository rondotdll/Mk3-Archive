package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "screenshot",
		Description: "Takes a screenshot of the system",
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
