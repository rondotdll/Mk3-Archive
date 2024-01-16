package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
	"os"
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
			if e != nil {
				println(Red + "An internal error has occured, and Mk3 needs to exit.\n Please open a new issue on github: " + e.Error() + Reset)
				os.Exit(1)
			}

			// ignore this, go is stupid and won't auto-convert int64 to int
			return "DoBSoD(" + strconv.Itoa(int(stopcode.Value.(int64))) + ")", nil
		},
	})
}
