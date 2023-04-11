package commands

import . "Mk3CLI/etc"

func init() {
	AddCommand(Command{
		Name:        "clear",
		Description: "Clears the console.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) {
			ClearConsole()
		},
	})
}
