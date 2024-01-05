package commands

import (
	"fmt"

	s7cli "mk3cli/s7cli/commands"
	feat "mk3cli/s7cli/features"
)

func init() {

	s7cli.DefaultHandler.AddCommand(s7cli.Command{
		Name:        "disable",
		Description: "Disables the specified features.",
		Args: s7cli.Args{
			{
				Name: s7cli.Name{
					"features",
					"f",
				},
				Datatype: "string",
				Required: false,
			},
		},
		Exec: func(input []string, this s7cli.Command) error {
			args, err := this.Args.Parse(input)

			if err == 1 {
				fmt.Println(s7cli.Red + "Invalid argument.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			} else if err == 2 {
				fmt.Println(s7cli.Red + "Not enough arguments.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			} else if err == 3 {
				fmt.Println(s7cli.Red + "Missing end quote.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			}
			i := 0

			for _, efeature := range feat.EnabledFeatures {
				if efeature.Feature.Name == args[0] {
					feat.EnabledFeatures = feat.RemoveFeature(feat.EnabledFeatures, i)
					println(fmt.Sprintf("The %v payload is now disabled", args[0]))
					return nil
				}
				i++
			}

			fmt.Println(s7cli.Red + "Feature not enabled. (try running 'ls')" + s7cli.White)
			return nil
		},
	})
}
