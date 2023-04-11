package commands

import (
	. "Mk3CLI/etc"
	"fmt"
)

func init() {
	AddCommand(Command{
		Name:        "landon",
		Description: "???",
		Args:        []Arg{},
		Exec: func(input []string, this Command) {
			fmt.Println("You've entered Landon mode.")
			AdvWriteLn("What's up chucklenuts, I'm landon. Now what were you looking to do?", 15)
			Input("[BETA] >> ")
		}})
}
