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
		Name:        "ls",
		Description: "Lists all currently supported features.",
		Args: []Arg{
			{
				Name: Name{
					"search",
					"s",
				},
				Datatype: "string",
				Required: false,
			},
		},
		Exec: func(input []string, this Command) error {
			args, err := ParseArgs(this.Args, input)

			if err == 1 {
				fmt.Println(Red + "Invalid argument.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return nil
			} else if err == 2 {
				fmt.Println(Red + "Not enough arguments.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return nil
			} else if err == 3 {
				fmt.Println(Red + "Missing end quote.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return nil
			}

			fmt.Println("Feature List:")
		outer:
			for _, f := range Features {
				if len(args) >= 1 && !(strings.Contains(strings.ToLower(f.Name), strings.ToLower(fmt.Sprint(args[0]))) || strings.Contains(strings.ToLower(f.Description), strings.ToLower(fmt.Sprint(args[0])))) {
					continue
				}
				for _, e := range EnabledFeatures {
					if f.Name == e.Feature.Name {
						fmt.Println("\n   ┌ " + Green + f.Name + White + DisplayEnabledArgs(f.Args, e.Args) + "\n   └───> " + Gray + f.Description + White)
						continue outer
					}
				}
				fmt.Println("\n   ┌ " + f.Name + DisplayArgs(f.Args) + "\n   └───> " + Gray + f.Description + White)
			}
			return nil
		},
	})
}