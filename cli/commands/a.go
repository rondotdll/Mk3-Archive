package commands

import (
	. "Mk3CLI/etc"
	"fmt"
	"strconv"
	"strings"
)

var Commands []Command = []Command{}

func Handle() {
	args := strings.Split(Input("[III] ("+strconv.Itoa(len(EnabledFeatures))+"/"+strconv.Itoa(len(Features))+") ~> "), " ")
	notfound := true

	for _, cmd := range Commands {
		if args[0] == cmd.Name {
			cmd.Exec(args, cmd)
			notfound = false
		}
	}

	if notfound {
		fmt.Println(Red + "command not found: '" + args[0] + "'" + White)
	}

	println()
}

func AddCommand(cmd Command) {
	Commands = append(Commands, cmd)
}
