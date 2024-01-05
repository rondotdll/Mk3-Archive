package commands

import (
	"github.com/c-bata/go-prompt"
)

type Arg struct {
	Name     Name
	Datatype string
	Required bool
}

type Args []Arg

type Name struct {
	Full  string
	Short string
}

// Command struct

type Command struct {
	Name        string
	Description string
	Args        Args
	Exec        func(input []string, this Command) error
}

type Handler struct {
	prompt     string
	commands   []Command
	completion []prompt.Suggest
}

// Displays the valid usage of a command to the terminal
func (this Command) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name.Format(false) + " [" + Cyan + a.Datatype + White + "] " + Reset
		} else {
			usage += DarkGray + "--" + a.Name.Format(true) + " [" + Cyan + a.Datatype + DarkGray + "] " + Reset
		}
	}
	println(usage)
}
