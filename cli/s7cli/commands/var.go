package commands

const (
	Reset = "\033[0m"

	// standard colors
	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Blue     = "\033[34m"
	Purple   = "\033[35m"
	Cyan     = "\033[36m"
	Gray     = "\033[37m"
	DarkGray = "\033[90m"
	White    = "\033[97m"

	// bold colors
	BoldRed    = "\033[31;1m"
	BoldGreen  = "\033[32;1m"
	BoldYellow = "\033[33;1m"
	BoldBlue   = "\033[34;1m"
	BoldPurple = "\033[35;1m"
	BoldCyan   = "\033[36;1m"
	BoldGray   = "\033[37;1m"
	BoldWhite  = "\033[97;1m"

	// underline
	UnderlineRed      = "\033[31;4m"
	UnderlineGreen    = "\033[32;4m"
	UnderlineYellow   = "\033[33;4m"
	UnderlineBlue     = "\033[34;4m"
	UnderlinePurple   = "\033[35;4m"
	UnderlineCyan     = "\033[36;4m"
	UnderlineGray     = "\033[37;4m"
	UnderlineDarkGray = "\033[90;4m"
	UnderlineWhite    = "\033[97;4m"
)

var (
	DefaultHandler = NewHandler("~> ")
)
