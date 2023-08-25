package commands

import (
	"fmt"
	"os"
)

// initalizes the 2 default lib (help & exit)
func (this *Handler) Init() Handler {

	this.AddCommand(Command{
		Name:        "help",
		Description: "Displays the list of lib.",
		Args:        []Arg{},
		Exec: func(args []string, command Command) error {
			if len(args) > 1 {
				for _, c := range this.commands {
					if args[0] == c.Name {
						c.DisplayUsage()
						println()
					}
				}
			}
			fmt.Println("List of all currently supported commands:\n")
			for _, c := range this.commands {
				print("  ")
				c.DisplayUsage()
			}
			return nil
		},
	})

	this.AddCommand(Command{
		Name:        "exit",
		Description: "Exits this application.",
		Args:        []Arg{},
		Exec: func(args []string, command Command) error {
			os.Exit(0)
			return nil
		},
	})

	return *this
}
