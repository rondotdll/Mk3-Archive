package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/c-bata/go-prompt"
)

func NewHandler(prompt string, commands ...Command) Handler {
	return Handler{prompt: prompt, commands: commands}
}

func (this Handler) completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(this.completion, d.GetWordBeforeCursor(), true)
}

func (this Handler) GetInput() string {
	return prompt.Input(this.prompt, this.completer)
}

func (this Handler) Handle(input string) {
	args := strings.Split(input, " ")
	notfound := true

	for _, c := range this.commands {
		if args[0] == c.Name {
			c.Exec(args, c)
			this.completion = append(this.completion, prompt.Suggest{
				Text: strings.Join(args, " "),
			})
			notfound = false
		}
	}

	if notfound {
		fmt.Println(Red + "Command not found: '" + args[0] + "'" + White)
	}
}

func (this *Handler) SetPrompt(prompt string) {
	this.prompt = prompt
}

func (this *Handler) AddCommand(command Command) {
	// check to verify required args come before any non-required
	in_required := true
	for _, arg := range command.Args {
		if !arg.Required && in_required {
			in_required = false
		} else if arg.Required && !in_required {
			panic("Command \"" + command.Name + "\" has required argument after non-required arguments!\n\tArgument: " + arg.Name.Full)
			os.Exit(1)
		}
	}
	this.commands = append(this.commands, command)
	this.completion = append(this.completion, prompt.Suggest{
		Text:        command.Name,
		Description: command.Description,
	})
}

// initalizes the CLI with 3 default commands (help, clear, exit)
func (this *Handler) Init() Handler {
	// Add the default commands

	// Exit command
	this.AddCommand(Command{
		Name:        "exit",
		Description: "Exits this application.",
		Args:        []Arg{},
		Exec: func(args []string, command Command) error {
			os.Exit(0)
			return nil
		},
	})

	// Help command
	this.AddCommand(Command{
		Name:        "help",
		Description: "Displays the list of ",
		Args:        []Arg{},
		Exec: func(args []string, command Command) error {
			if len(args) > 1 {
				for _, c := range this.commands {
					if args[0] == c.Name {
						c.DisplayUsage()
						println()
					}
				}
			}
			fmt.Println("List of all currently supported commands:\n")
			for _, c := range this.commands {
				print("  ")
				c.DisplayUsage()
			}
			return nil
		},
	})

	// Clear command
	this.AddCommand(Command{
		Name:        "clear",
		Description: "Clears the console.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) error {
			os_switch := make(map[string]func()) //Initialize it
			os_switch["linux"] = func() {
				cmd := exec.Command("clear") //Linux example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
			os_switch["windows"] = func() {
				cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
				cmd.Stdout = os.Stdout
				cmd.Run()
			}

			value, ok := os_switch[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
			if ok {                              //if we defined a clear func for that platform:
				value() //we execute it
			} else { //unsupported platform
				fmt.Println("Failed; Your terminal isn't ANSI! :(")
			}
			return nil
		},
	})
	return *this
}
