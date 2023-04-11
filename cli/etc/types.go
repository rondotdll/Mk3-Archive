package etc

type Arg struct {
	Name     string
	Datatype string
	Required bool
}

type Command struct {
	Name        string
	Description string
	Args        []Arg
	Exec        func(input []string, this Command)
}

type Feature struct {
	Name        string
	Description string
	Args        []Arg
}

type FeatureSet struct {
	Feature Feature
	Enabled bool
	Args    []FeatureSetArg
}

type FeatureSetArg struct {
	Arg   Arg
	Value interface{}
}
