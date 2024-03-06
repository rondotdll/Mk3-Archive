package lib

import (
	. "mk3cli/s7cli/commands"
	"os"
)

func Handle(e error) {
	if e != nil {
		println(Red + "An internal error has occured, and Mk3 needs to exit.\n Please open a new issue on github: " + e.Error() + Reset)
		os.Exit(1)
	}
}
