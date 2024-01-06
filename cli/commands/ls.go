package commands

import (
	"fmt"
	"strings"

	s7cli "mk3cli/s7cli/commands"
	feat "mk3cli/s7cli/features"
)

func init() {
	s7cli.DefaultHandler.AddCommand(s7cli.Command{
		Name:        "ls",
		Description: "Lists all currently supported feat.",
		Args: s7cli.Args{
			{
				Name: s7cli.Name{
					"search",
					"s",
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

			fmt.Println("Feature List:")
		outer:
			for _, f := range feat.FeatureList {
				featureName := strings.ToLower(f.Name)
				featureDescription := strings.ToLower(f.Description)

				// if the search argument is provided, only display feat that match the search
				if len(args) >= 1 {
					arg_0 := strings.ToLower(fmt.Sprint(args[0]))
					if !(strings.Contains(featureName, arg_0) || strings.Contains(featureDescription, arg_0)) {
						continue
					}
				}

				// check if the feat is enabled
				// ik this isn't the most efficient way to do this
				for _, e := range feat.EnabledFeatures {
					if f.Name == e.Feature.Name {
						fmt.Println("\n   ┌ " + s7cli.Green + f.Name + s7cli.White + " " + feat.DisplayEnabledArgs(f.Args, e.Args) + "\n   └───> " + s7cli.Gray + f.Description + s7cli.Reset)
						continue outer
					}
				}
				fmt.Println("\n   ┌ " + f.Name + " " + s7cli.DisplayArgs(f.Args) + "\n   └───> " + s7cli.Gray + f.Description + s7cli.Reset)
			}
			return nil
		},
	})
}
