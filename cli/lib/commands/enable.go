package commands

import "C"
import (
	"fmt"
	. "mk3cli/etc"
	. "mk3cli/lib/commands/base"
	. "mk3cli/lib/features/base"
	"strings"
)

func init() {
	DefaultHandler.AddCommand(Command{
		Name:        "enable",
		Description: "Enables the specified feature.",
		Args: []Arg{
			{
				Name: Name{
					"feature",
					"f",
				},
				Datatype: "string",
				Required: true,
			},
			{
				Name: Name{
					"args",
					"a",
				},
				Datatype: "...string",
				Required: false,
			},
		},
		Exec: func(input []string, this Command) error {
			args := this.ParseArgs(strings.Join(input, " "))
			if args == nil {
				return nil
			}

			initial := len(EnabledFeatures)

		outer:
			for _, feature := range Features {
				if args[0] == feature.Name {
					for _, efeature := range EnabledFeatures {
						if efeature.Feature.Name == feature.Name {
							fmt.Println(Red + "Feature already enabled. (run 'ls' to see more details)" + White)
							return nil
						}
					}

					fargs := feature.ParseArgs(args[0].(string))
					if fargs == nil {
						return nil
					}

					var FeatureArgs []FeatureSetArg
					i := 0
					for _, arg := range feature.Args {
						if !arg.Required && (CountRequired(feature.Args) < len(args)) {
							if len(args)-1 == 0 {
								FeatureArgs = append(FeatureArgs, FeatureSetArg{
									Arg:   arg,
									Value: "",
								})
								continue
							} else if len(args)-1 < 0 {
								continue outer
							}

						}
						FeatureArgs = append(FeatureArgs, FeatureSetArg{
							Arg:   arg,
							Value: fargs[i],
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
				return nil
			}

			println(fmt.Sprintf("The %v payload is now enabled", args[0]))
			fmtargs := FormatArgValues(EnabledFeatures[len(EnabledFeatures)-1].Args)
			if fmtargs != "" {
				println(fmtargs)
			}
			return nil

		}})
}
