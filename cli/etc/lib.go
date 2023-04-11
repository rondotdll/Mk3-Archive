package etc

import (
	"bufio"
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func Write(str string, delay int) {
	for _, char := range str {
		print(string(char))
		Sleep(delay)
	}
}

func WriteLn(str string, delay int) {
	for _, char := range str {
		print(string(char))
		Sleep(delay)
	}
	println("")
}

func AdvWriteLn(str string, std_delay int) {
	sentences := strings.Split(str, ".")
	for _, sentence := range sentences {
		fragments := strings.Split(sentence, ",")
		for i, fragment := range fragments {
			Write(fragment, std_delay)
			if len(fragments) > 1 && i != len(fragments)-1 {
				fmt.Print(",")
				Sleep(250)
			}
		}
		if !(strings.HasSuffix(sentence, "?") || strings.HasPrefix(sentence, "!")) {
			fmt.Print(".")
		}
		Sleep(500)
	}
	println()
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IndexOf(slice []string, str string) int {
	i := 0

	for _, s := range slice {
		if s == str {
			break
		}
		i++
	}
	return i

	return sort.StringSlice(slice).Search(str)
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsValidURL(addr string) bool {
	_, err := url.ParseRequestURI(addr)
	if err == nil {
		return false
	}
	return true
}

func ClearConsole() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		fmt.Println("Failed; Your terminal isn't ANSI! :(")
	}
}

func DisplayArgs(args []Arg) string {
	output := " "

	for _, a := range args {
		if a.Required {
			output += "<" + a.Name + " (" + Gray + a.Datatype + White + ")> "
		} else {
			output += "[" + a.Name + " (" + Gray + a.Datatype + White + ")] "
		}
	}
	return output
}

func DisplayEnabledArgs(args []Arg, enabled []FeatureSetArg) string {
	output := " "

	x := 0
	for _, a := range args {
		enabledVal := fmt.Sprintf("%v", enabled[x].Value)

		if len(enabledVal) > 10 {
			enabledVal = enabledVal[0:10] + Gray + "..."
		}

		if a.Required {
			output += "<" + a.Name + "=" + fmt.Sprintf("\""+Green+"%v"+White+"\"", enabledVal) + " (" + Gray + a.Datatype + White + ")> "
		} else if !a.Required && x < len(enabled) {
			output += "[" + a.Name + "=" + fmt.Sprintf("\""+Green+"%v"+White+"\"", enabledVal) + " (" + Gray + a.Datatype + White + ")] "
		} else {
			output += "[" + a.Name + " (" + Gray + a.Datatype + White + ")] "
		}
		x++
	}
	return output
}

func RemoveFeature(slice []FeatureSet, s int) []FeatureSet {
	return append(slice[:s], slice[s+1:]...)
}

func FormatArgValues(args []FeatureSetArg) string {
	var output string

	for _, arg := range args {
		output += fmt.Sprintf("  %s: %v\n", arg.Arg.Name, arg.Value)
	}

	return strings.TrimSuffix(output, "\n")
}

func CountRequired(args []Arg) int {
	out := 0

	for _, arg := range args {
		if arg.Required {
			out++
		}
	}

	return out
}

func ParseArgs(args []Arg, input []string) (parsedargs []interface{}, err int) {
	var output []interface{}
	argsRaw := input
	reqLen := CountRequired(args)

	i := 1

	if reqLen > len(argsRaw)-1 {
		return output, 2
	}

	complexSuffix := ""
	offset := 0

	for _, arg := range args {
		if i > len(args) || i >= len(input) {
			break
		}
		switch arg.Datatype {
		case "string":
			appendStr := argsRaw[i+offset]
			if (strings.HasPrefix(argsRaw[i+offset], "\"") || strings.HasPrefix(argsRaw[i+offset], "'")) && !(strings.HasSuffix(argsRaw[i+offset], "\"") || strings.HasSuffix(argsRaw[i+offset], "'")) {
				complexSuffix = string((argsRaw[i+offset])[0])
				x := offset
				for {
					x++
					if i+x >= len(argsRaw) {
						return output, 3
					}
					appendStr += " " + argsRaw[i+x]
					if strings.HasSuffix(argsRaw[i+x], complexSuffix) {
						offset = x
						break
					}
				}
			}
			output = append(output, strings.ReplaceAll(strings.ReplaceAll(appendStr, "\"", ""), "'", ""))
			break
		case "int":
			parsed, err := strconv.Atoi(argsRaw[i+offset])
			if err != nil {
				return output, 1
			}
			output = append(output, parsed)
			break
		case "bool":
			parsed, err := strconv.ParseBool(argsRaw[i+offset])
			if err != nil {
				return output, 1
			}
			output = append(output, parsed)
			break
		}
		i++
	}
	return output, 0

}

func SliceItS(slice []interface{}) []string {
	output := []string{}
	for _, item := range slice {
		output = append(output, fmt.Sprintf("%v", item))
	}

	return output
}

func Input(prompt string) string {
	if prompt != "" {
		fmt.Print(prompt)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}
