package commands

import "C"
import (
	. "Mk3CLI/etc"
	"fmt"
	"strconv"
)

func init() {
	AddCommand(Command{
		Name:        "enable",
		Description: "Enables the specified feature.",
		Args: []Arg{
			{Name: "feature", Datatype: "string", Required: true},
			{Name: "args", Datatype: "string...", Required: false},
		},
		Exec: func(input []string, this Command) {
			args, err := ParseArgs(this.Args, input)

			if err == 1 {
				fmt.Println(Red + "Invalid argument.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return
			} else if err == 2 {
				fmt.Println(Red + "Not enough arguments.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return
			} else if err == 3 {
				fmt.Println(Red + "Missing end quote.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
				return
			}

			initial := len(EnabledFeatures)

		outer:
			for _, feature := range Features {
				if args[0] == feature.Name {
					for _, efeature := range EnabledFeatures {
						if efeature.Feature.Name == feature.Name {
							fmt.Println(Red + "Feature already enabled. (run 'ls' to see more details)" + White)
							return
						}
					}

					args2, err := ParseArgs(feature.Args, input[1:])

					if err == 1 {
						fmt.Println(Red + "Bad argument type.\n" + White + "Arguments for '" + feature.Name + "':\n   " + DisplayArgs(feature.Args))
						return
					} else if err == 2 {
						fmt.Println(Red + "Not enough arguments.\n" + White + "'" + feature.Name + "' requires at least " + strconv.Itoa(CountRequired(feature.Args)) + " Arguments:\n  " + DisplayArgs(feature.Args))
						return
					} else if err == 3 {
						fmt.Println(Red + "Missing end quote.\n" + White + "Usage:\n  enable" + DisplayArgs(this.Args))
						return
					}
					var FeatureArgs []FeatureSetArg
					i := 0
					for _, f_arg := range feature.Args {
						if !f_arg.Required && (CountRequired(feature.Args) < len(args)) {
							if len(args)-1 == 0 {
								FeatureArgs = append(FeatureArgs, FeatureSetArg{
									Arg:   f_arg,
									Value: "",
								})
								continue
							} else if len(args)-1 < 0 {
								continue outer
							}

						}
						FeatureArgs = append(FeatureArgs, FeatureSetArg{
							Arg:   f_arg,
							Value: args2[i],
						})
						i++
					}

					EnabledFeatures = append(EnabledFeatures, FeatureSet{
						Feature: feature,
						Enabled: true,
						Args:    FeatureArgs,
					})
					break
				}
			}

			if len(EnabledFeatures) == initial {
				fmt.Println(Red + "Invalid feature. (try running 'ls')" + White)
				return
			}

			println(fmt.Sprintf("The %v payload is now enabled", args[0]))
			fmtargs := FormatArgValues(EnabledFeatures[len(EnabledFeatures)-1].Args)
			if fmtargs != "" {
				println(fmtargs)
			}
		}})
}
