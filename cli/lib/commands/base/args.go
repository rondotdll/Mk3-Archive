package commands

func (this Name) Format(gray bool) string {
	output := ""
	char_found := false

	format := []string{Reset, UnderlineGray}
	if gray {
		format = []string{Reset + DarkGray, UnderlineDarkGray}
	}

	for _, c := range this.Full {
		if string(c) == this.Short && !char_found {
			output += format[1] + string(c) + format[0]
			char_found = true
			continue
		}
		output += string(c)
	}

	return output
}

func (this Command) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name.Format(false) + " [" + a.Datatype + "] "
		} else {
			usage += DarkGray + "--" + a.Name.Format(true) + " [" + Cyan + a.Datatype + DarkGray + "] " + Reset
		}
	}
	println(usage)
}

func DisplayArgs(args []Arg) string {
	output := " "

	for _, a := range args {
		if a.Required {
			output += "<" + a.Name.Format(true) + " (" + Gray + a.Datatype + White + ")> "
		} else {
			output += "[" + a.Name.Format(true) + " (" + Gray + a.Datatype + White + ")] "
		}
	}
	return output
}
