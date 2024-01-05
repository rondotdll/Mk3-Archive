package commands

func (this Name) Format(gray bool) string {
	output := ""
	char_found := false

	format := []string{Reset + Gray, UnderlineGray}
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

func DisplayArgs(args []Arg) string {
	output := ""

	for _, a := range args {
		if a.Required {
			output += Reset + "<" + Gray + a.Name.Format(false) + Reset + " (" + Cyan + a.Datatype + Reset + ")> "
		} else {
			output += Gray + "[" + DarkGray + a.Name.Format(true) + Gray + " (" + Cyan + a.Datatype + Gray + ")] "
		}
	}
	return output
}
