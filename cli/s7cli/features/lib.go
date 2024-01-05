package features

import (
	"fmt"
	"strings"
)

func RemoveFeature(slice []FeatureSet, s int) []FeatureSet {
	return append(slice[:s], slice[s+1:]...)
}

func FormatArgValues(args []FeatureSetArg) string {
	var output string

	for _, arg := range args {
		output += fmt.Sprintf("  %s: %v\n", arg.Arg.Name, arg.Value)
	}

	return strings.TrimSuffix(output, "\n")
}
