package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "shutdown",
		Description: "Triggers a system shutdown when execution finishes",
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
