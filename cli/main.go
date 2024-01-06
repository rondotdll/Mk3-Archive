package main

import (
	"fmt"
	"strconv"

	_ "mk3cli/commands"
	_ "mk3cli/features"

	s7cli "mk3cli/s7cli/commands"
	feat "mk3cli/s7cli/features"
)

var (
	Splash string = s7cli.Green + `
         ___           ___
        /\  \         /|  |
       |::\  \       |:|  |
       |:|:\  \      |:|  |     /\\\\\\\\\\\  /\\\\\\\\\\\  /\\\\\\\\\\\
     __|:|\:\  \   __|:|  |     \/////\\\///  \/////\\\///  \/////\\\///
    /::::|_\:\__\ /\ |:|__|____      \/\\\         \/\\\         \/\\\
    \:\~~\  \/__/ \:\/:::::/__/       \/\\\         \/\\\         \/\\\
     \:\  \        \::/~~/~            \/\\\         \/\\\         \/\\\
      \:\  \        \:\~~\              \/\\\         \/\\\         \/\\\
       \:\__\        \:\__\           /\\\\\\\\\\\  /\\\\\\\\\\\  /\\\\\\\\\\\
        \/__/         \/__/           \///////////  \///////////  \///////////


    ` + s7cli.Yellow + `COMMUNITY EDITION` + s7cli.Gray + `                                  by ` + s7cli.White + `Studio 7 Development`

	Info string = s7cli.Gray + `	Thanks for downloading! Sorry for the wait... the biggest challenge
	was figuring out how the hell we could bypass defender again since
	every C# payload was patched a few years back. You are using the 
	Community Edition of this software, which has been constructed with
	CyberSecurity Specialists and Pen-Testers in mind.
	
	visit ` + s7cli.Cyan + `https://studio7.dev/mk3` + s7cli.Gray + ` for more cool software.

								- S7 Dev Team

    Change Log [v1.0.0 ` + s7cli.Green + `L` + s7cli.Gray + `u` + s7cli.Green + `c` + s7cli.Gray + `k` + s7cli.Green + `y ` + s7cli.Yellow + `7` + s7cli.Gray + `]:
      [` + s7cli.Green + `+` + s7cli.White + `] Initial Release
`
)

func main() {
	fmt.Println(Splash)
	fmt.Println("\n" + Info)
	fmt.Println("Try running 'help' for a list of commands.\n")
	for {
		s7cli.DefaultHandler.SetPrompt("[III] " + strconv.Itoa(len(feat.EnabledFeatures)) + "/" + strconv.Itoa(len(feat.FeatureList)) + " ~> ")
		s7cli.DefaultHandler.Handle(s7cli.DefaultHandler.GetInput())
		println()
	}
}
