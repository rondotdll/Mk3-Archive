package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "killdp",
		Description: "Temporarily removes the system's desktop [kills explorer.exe]",
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
