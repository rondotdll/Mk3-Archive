package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "geo+",
		Description: "Advanced Geo Location payload, grabs precise geo coordinates of the connected router",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "CODE GOES HERE"

			return output, nil
		},
	})
}
