package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "bsod",
		Description: "Triggers a Blue Screen of Death when execution finishes.",
		Args: []Arg{
			{
				Name: Name{
					"stopcode",
					"c",
				},
				Datatype: "int",
				Required: true,
			},
		},
		Dependencies: []string{
			"system.lib.go",
		},
	})
}
