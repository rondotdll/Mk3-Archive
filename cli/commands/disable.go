package commands

import (
	. "Mk3CLI/etc"
	"fmt"
)

func init() {
	AddCommand(Command{
		Name:        "disable",
		Description: "Disables the specified feature.",
		Args: []Arg{
			{Name: "feature", Datatype: "string", Required: true},
		},
		Exec: func(input []string, this Command) {
			args, err := ParseArgs(this.Args, input)

			if err == 1 {
				fmt.Println(Red + "Invalid argument.\n" + White + "Usage:\n  enable" + DisplayArgs(Commands[1].Args))
				return
			} else if err == 2 {
				fmt.Println(Red + "Not enough arguments.\n" + White + "Usage:\n enable" + DisplayArgs(Commands[1].Args))
				return
			} else if err == 3 {
				fmt.Println(Red + "Missing end quote.\n" + White + "Usage:\n enable" + DisplayArgs(Commands[1].Args))
				return
			}

			i := 0

			for _, efeature := range EnabledFeatures {
				if efeature.Feature.Name == args[0] {
					EnabledFeatures = RemoveFeature(EnabledFeatures, i)
					println(fmt.Sprintf("The %v payload is now disabled", args[0]))
					return
				}
				i++
			}

			fmt.Println(Red + "Feature not enabled. (try running 'ls')" + White)

		}})
}
