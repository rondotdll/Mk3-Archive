package commands

import (
	"fmt"
	. "mk3cli/etc"
	. "mk3cli/lib/commands/base"
	. "mk3cli/lib/features/base"
	"strings"
)

func init() {

	DefaultHandler.AddCommand(Command{
		Name:        "disable",
		Description: "Disables the specified feature.",
		Args: []Arg{
			{
				Name: Name{
					"feature",
					"f",
				},
				Datatype: "string",
				Required: false,
			},
		},
		Exec: func(input []string, this Command) error {
			args := this.ParseArgs(strings.Join(input, " "))

			i := 0

			for _, efeature := range EnabledFeatures {
				if efeature.Feature.Name == args[0] {
					EnabledFeatures = RemoveFeature(EnabledFeatures, i)
					println(fmt.Sprintf("The %v payload is now disabled", args[0]))
					return nil
				}
				i++
			}

			fmt.Println(Red + "Feature not enabled. (try running 'ls')" + White)
			return nil
		},
	})
}
