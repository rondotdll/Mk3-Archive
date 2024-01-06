package features

import (
	. "mk3cli/s7cli/commands"
)

// Feature Structs

type Feature struct {
	Name         string
	Description  string
	Args         Args
	Dependencies []string
	ReturnsData  bool
}

type Features []Feature

type FeatureSet struct {
	Feature Feature
	Enabled bool
	Args    []FeatureSetArg
}

type FeatureSetArg struct {
	Arg   Arg
	Value interface{}
}

func (this Feature) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name.Format(false) + " [" + a.Datatype + "] "
		} else {
			usage += "--" + DarkGray + a.Name.Format(true) + " [" + Cyan + a.Datatype + DarkGray + "] "
		}
	}
	println(usage)
}
