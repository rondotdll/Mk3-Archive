package commands

import (
	"strconv"
	"strings"
)

// Returns the number of required arguments in a slice of Args
func CountRequired(args Args) int {
	out := 0

	for _, arg := range args {
		if arg.Required {
			out++
		}
	}

	return out
}

// Parses the arguments [for a Command object] from the raw input string
// and returns them as a slice of interface{}'s, with each index corresponding
// to the index & type of the argument in the "Command" object's "Args" slice.
// Error Codes:
//
//	0: Success (No Error)
//	1: Invalid argument (Mismatched datatype)
//	2: Not enough arguments (Failed to satisfy required args)
//	3: Missing end quote (Complex string missing end quote)
func (this Args) Parse(input []string) (parsed_args []interface{}, err int) {

	var output []interface{}
	argsRaw := input // should be the input split by spaces
	reqLen := CountRequired(this)

	// verify all required arguments have been satisfied
	if reqLen > len(argsRaw)-1 {
		return output, 2
	}

	terminator := ""
	base_arg_pos := 1 // current argument's position (starting point for scanning)
	arg_offset := 0   // current position being scanned for complex strings (represents everything already scanned)

	for _, arg := range this {
		if base_arg_pos > len(this) || base_arg_pos >= len(input) {
			break
		}
		switch arg.Datatype {
		case "string":
			buffer := argsRaw[base_arg_pos+arg_offset] // the buffer is the string we're appending to the output

			if (strings.HasPrefix(argsRaw[base_arg_pos+arg_offset], "\"") || strings.HasPrefix(argsRaw[base_arg_pos+arg_offset], "'")) &&
				!(strings.HasSuffix(argsRaw[base_arg_pos+arg_offset], "\"") || strings.HasSuffix(argsRaw[base_arg_pos+arg_offset], "'")) {
				terminator = string((argsRaw[base_arg_pos+arg_offset])[0]) // if we found a complex string, set the terminator to the first character
				carat_pos := arg_offset                                    // we're handing off the arg_offset to the carat_pos so we can increment it without affecting the arg_offset
				for {
					carat_pos++
					// if we haven't found a terminator by the end of the input, return an error
					if base_arg_pos+carat_pos >= len(argsRaw) {
						return output, 3
					}
					buffer += " " + argsRaw[base_arg_pos+carat_pos]
					if strings.HasSuffix(argsRaw[base_arg_pos+carat_pos], terminator) {
						arg_offset = carat_pos
						break
					}
				}
			}
			output = append(output, strings.ReplaceAll(strings.ReplaceAll(buffer, "\"", ""), "'", ""))
			break
		case "int":
			var err error
			var parsed int64
			var current_arg = argsRaw[base_arg_pos+arg_offset]

			// Attempt to parse the integer in base 10 (regular number), if that fails, try base 16 (hex)
			parsed, err = strconv.ParseInt(current_arg, 10, strconv.IntSize)
			if err != nil {
				parsed, err = strconv.ParseInt(strings.TrimPrefix(current_arg, "0x"), 16, strconv.IntSize)
				if err != nil {
					return output, 1
				}
			}
			output = append(output, parsed)
			break
		case "bool":
			parsed, err := strconv.ParseBool(argsRaw[base_arg_pos+arg_offset])
			if err != nil {
				return output, 1
			}
			output = append(output, parsed)
			break
		}
		base_arg_pos++
	}
	return output, 0

}
