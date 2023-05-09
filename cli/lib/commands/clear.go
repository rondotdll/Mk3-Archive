package commands

import (
	. "mk3cli/etc"
	. "mk3cli/lib/commands/base"
)

func init() {
	DefaultHandler.AddCommand(Command{
		Name:        "clear",
		Description: "Clears the console.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) error {
			ClearConsole()
			return nil
		},
	})
}
