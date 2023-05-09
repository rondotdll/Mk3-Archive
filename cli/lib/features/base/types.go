package features

import (
	"mk3cli/lib/commands/base"
)

type Feature struct {
	Name         string
	Description  string
	Args         []commands.Arg
	Dependencies []string
}

type FeatureSet struct {
	Feature Feature
	Enabled bool
	Args    []FeatureSetArg
}

type FeatureSetArg struct {
	Arg   commands.Arg
	Value interface{}
}

func (this Feature) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name.Format(false) + " [" + a.Datatype + "] "
		} else {
			usage += "--" + commands.DarkGray + a.Name.Format(true) + " [" + commands.Cyan + a.Datatype + commands.DarkGray + "] "
		}
	}
	println(usage)
}
