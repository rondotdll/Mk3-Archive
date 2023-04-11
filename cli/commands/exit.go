package commands

import (
	. "Mk3CLI/etc"
	"os"
)

func init() {
	AddCommand(Command{
		Name:        "exit",
		Description: "Closes the Mk3 CLI.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) {
			println(Purple + "Goodbye for now 👋" + Reset)
			os.Exit(0)
		}})
}
