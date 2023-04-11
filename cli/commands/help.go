package commands

import (
	. "Mk3CLI/etc"
	"fmt"
)

func init() {
	AddCommand(Command{
		Name:        "help",
		Description: "Displays the list of commands.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) {
			fmt.Println("Command List:")
			for _, cmd := range Commands {
				fmt.Println("\n   ┌ " + cmd.Name + DisplayArgs(cmd.Args) + "\n   └───> " + Gray + cmd.Description + White)
			}
		}})
}
