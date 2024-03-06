package features

import (
	. "mk3cli/s7cli/commands"
	. "mk3cli/s7cli/features"
	"os"
)

func init() {
	FeatureList = append(FeatureList, Feature{
		Name:        "error",
		Description: "Displays a custom fake error message when execution finishes",
		ReturnsData: false,
		Args: []Arg{
			{
				Name: Name{
					"title",
					"t",
				},
				Datatype: "string",
				Required: true,
			}, {
				Name: Name{
					"description",
					"d",
				},
				Datatype: "string",
				Required: true,
			},
		},
		Dependencies: []string{
			"system.lib.go",
		},
		GenerateCode: func(args FeatureSetArgsList) (string, error) {
			title, e := args.Find("description")
			if e != nil {
				println(Red + "An internal error has occured, and Mk3 needs to exit.\n Please open a new issue on github: " + e.Error() + Reset)
				os.Exit(1)
			}

			description, e2 := args.Find("title")
			if e2 != nil {
				println(Red + "An internal error has occured, and Mk3 needs to exit.\n Please open a new issue on github: " + e2.Error() + Reset)
				os.Exit(1)
			}

			output := "DisplayErrorMsg(\"" + title.Value.(string) + "\", \"" + description.Value.(string) + "\")\n"

			return output, nil
		},
	})
}
