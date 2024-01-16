package commands

import (
	"fmt"

	s7cli "mk3cli/s7cli/commands"
	feat "mk3cli/s7cli/features"
)

func init() {

	s7cli.DefaultHandler.AddCommand(s7cli.Command{
		Name:        "enable",
		Description: "Enables the specified feat.",
		Args: s7cli.Args{
			{
				Name: s7cli.Name{
					"feat",
					"f",
				},
				Datatype: "string",
				Required: true,
			},
			{
				Name: s7cli.Name{
					"args",
					"a",
				},
				Datatype: "...string",
				Required: false,
			},
		},
		Exec: func(input []string, this s7cli.Command) error {
			args, err := this.Args.Parse(input)

			if err == 1 {
				fmt.Println(s7cli.Red + "Invalid argument.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			} else if err == 2 {
				fmt.Println(s7cli.Red + "Not enough arguments.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			} else if err == 3 {
				fmt.Println(s7cli.Red + "Missing end quote.\n" + s7cli.White + "Usage:\n  enable " + s7cli.DisplayArgs(this.Args))
				return nil
			}

			initial := len(feat.EnabledFeatures)

		outer:
			for _, feature := range feat.FeatureList {
				if args[0] == feature.Name {
					for _, efeature := range feat.EnabledFeatures {
						if efeature.Feature.Name == feature.Name {
							fmt.Println(s7cli.Red + "Feature already enabled. (run 'ls' to see more details)" + s7cli.White)
							return nil
						}
					}

					f_args, f_err := feature.Args.Parse(input[1:])

					if f_err == 1 {
						fmt.Println(s7cli.Red + "Invalid feat argument.\n" + s7cli.White + "Usage:\n  enable " + feature.Name + " " + s7cli.DisplayArgs(feature.Args))
						return nil
					} else if f_err == 2 {
						fmt.Println(s7cli.Red + "Not enough feat arguments.\n" + s7cli.White + "Usage:\n  enable " + feature.Name + " " + s7cli.DisplayArgs(feature.Args))
						return nil
					} else if f_err == 3 {
						fmt.Println(s7cli.Red + "Missing end quote.\n" + s7cli.White + "Usage:\n  enable " + feature.Name + " " + s7cli.DisplayArgs(feature.Args))
						return nil
					}

					var FeatureArgs []feat.FeatureSetArg
					i := 0
					for _, arg := range feature.Args {
						if !arg.Required && (s7cli.CountRequired(feature.Args) < len(args)) {
							if len(args)-1 == 0 {
								FeatureArgs = append(FeatureArgs, feat.FeatureSetArg{
									Arg:   arg,
									Value: "",
								})
								continue
							} else if len(args)-1 < 0 {
								continue outer
							}

						}
						FeatureArgs = append(FeatureArgs, feat.FeatureSetArg{
							Arg:   arg,
							Value: f_args[i],
						})
						i++
					}

					feat.EnabledFeatures = append(feat.EnabledFeatures, feat.FeatureSet{
						Feature: feature,
						Enabled: true,
						Args:    FeatureArgs,
					})
					break
				}
			}

			if len(feat.EnabledFeatures) == initial {
				fmt.Println(s7cli.Red + "Invalid feat. (try running 'ls')" + s7cli.White)
				return nil
			}

			println(fmt.Sprintf("The %v payload is now enabled", args[0]))
			fmtargs := feat.FormatArgValues(feat.EnabledFeatures[len(feat.EnabledFeatures)-1].Args)
			if fmtargs != "" {
				println(fmtargs)
			}
			return nil

		}})
}
