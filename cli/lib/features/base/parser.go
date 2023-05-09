package features

import (
	commands "mk3cli/lib/commands/base"
	"os"
	"strconv"
	"strings"
)

func (this Feature) ParseArgs(input string) []interface{} {
	// convert our args into arg/value "chunks"
	chunks := strings.Split(input, " ")
	in_string, buffer, terminator := false, "", ""
	parsed_chunks := []string{}

	// merges strings into a single chunk
	for _, chunk := range chunks {
		if strings.HasPrefix(chunk, "\"") || strings.HasPrefix(chunk, "'") {
			in_string = true
			terminator = string(chunk[0])
			buffer = buffer + strings.TrimLeft(chunk, string(chunk[0]))
		} else if in_string {
			if strings.HasSuffix(chunk, terminator) {
				in_string = false
				terminator = ""
				buffer = buffer + strings.TrimRight(chunk, terminator)
				parsed_chunks = append(parsed_chunks, buffer)
			}
		} else {
			parsed_chunks = append(parsed_chunks, chunk)
		}
	}

	output := make([]interface{}, len(this.Args))

	// parse our chunks based on command argument Datatypes
	for c_index, c := range parsed_chunks {
		if strings.HasPrefix(c, "-") {
			arg_name := strings.Replace(c, "-", "", 2)
			for arg_index, arg := range this.Args {
				// if the inputted arg and the required args match
				if arg.Name.Full == arg_name || arg_name == arg.Name.Short {
					switch arg.Datatype {
					case "string":
						output[arg_index] = parsed_chunks[c_index+1]
						break
					case "int":
						p_int, e := strconv.Atoi(parsed_chunks[c_index+1])
						if e != nil {
							println(commands.Red + "Invalid argument type supplied \"" + c + "\" (expected Integer)\n\tArg: " + arg.Name.Full + commands.Reset)
							this.DisplayUsage()
						}
						output[arg_index] = p_int
						break
					case "bool":
						p_bool, e := strconv.ParseBool(parsed_chunks[c_index+1])
						if e != nil {
							println(commands.Red + "Invalid argument type supplied \"" + c + "\" (expected Boolean)\n\tArg: " + arg.Name.Full + commands.Reset)
							this.DisplayUsage()
							return nil
						}
						output[arg_index] = p_bool
						break
					case "...string":
						break
					default:
						panic(commands.Red + "Invalid argument type in feature \"" + this.Name + "\", expected either 'string', 'bool', or 'int'\n (found " + arg.Datatype + ") " + commands.Reset)
						os.Exit(1)
					}
				}
			}
		}
	}

	// populate any non-required arguments with "nil"
	for i := 0; len(output) < len(this.Args); i++ {
		if this.Args[i].Required {
			println(commands.Red + "Not enough arguments for feature \"" + this.Name + "\", expected at least " + strconv.Itoa(len(this.Args)))
			return nil
		}
		output = append(output, nil)
	}

	return output
}
