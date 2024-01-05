package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "pcinfo",
		Description: "Grabs general basic system information",
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
