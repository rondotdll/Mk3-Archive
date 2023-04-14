package main

import (
	"Mk3CLI/commands"
	"Mk3CLI/etc"
	"fmt"
)

func main() {

	fmt.Println(etc.Splash)
	fmt.Println("\n" + etc.Info)
	fmt.Println("Try running 'help' for a list of commands.\n")
	for {
		commands.Handle()
	}
}
