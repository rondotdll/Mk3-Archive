package commands

import "github.com/c-bata/go-prompt"

type Arg struct {
	Name     Name
	Datatype string
	Required bool
}

type Handler struct {
	prompt     string
	commands   []Command
	completion []prompt.Suggest
}

type Name struct {
	Full  string
	Short string
}

type Command struct {
	Name        string
	Description string
	Args        []Arg
	Exec        func(input []string, this Command) error
}

// ansi definitions
const ()
