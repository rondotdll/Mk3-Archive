package features

import (
	"mk3cli/commands/lib"
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
	"strconv"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "bsod",
		Description: "Triggers a Blue Screen of Death when execution finishes.",
		ReturnsData: false,
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
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			stopcode, e := args.Find("stopcode")
			lib.Handle(e)

			// ignore this, go is stupid and won't auto-compress int64 to int
			return "DoBSoD(" + strconv.Itoa(int(stopcode.Value.(int64))) + ")", nil
		},
	})
}
