package commands

import (
	"bufio"
	"fmt"
	"net/mail"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"sort"
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
		fmt.Println("\n")
	}
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
