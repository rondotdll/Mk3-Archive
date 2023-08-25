package main

import (
	"fmt"
	"mk3cli/etc"
	"mk3cli/lib/commands"
	. "mk3cli/lib/commands/base"
	. "mk3cli/lib/features/base"
	"strconv"
)

func main() {
	print(commands.Garbage)
	fmt.Println(Splash)
	fmt.Println("\n" + Info)
	fmt.Println("Try running 'help' for a list of commands.\n")
	DefaultHandler.SetPrompt("[III] " + strconv.Itoa(len(EnabledFeatures)) + "/" + strconv.Itoa(len(etc.Features)) + " ~> ")
	for {
		DefaultHandler.Handle(DefaultHandler.GetInput())
		println()
	}
}
