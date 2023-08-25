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
	UnderlineRed             = "\033[31;4m"
	UnderlineGreen           = "\033[32;4m"
	UnderlineYellow          = "\033[33;4m"
	UnderlineBlue            = "\033[34;4m"
	UnderlinePurple          = "\033[35;4m"
	UnderlineCyan            = "\033[36;4m"
	UnderlineGray            = "\033[37;4m"
	UnderlineDarkGray        = "\033[90;4m"
	UnderlineWhite           = "\033[97;4m"
	Splash            string = Green + `         ___           ___
        /\  \         /|  |
       |::\  \       |:|  |
       |:|:\  \      |:|  |     /\\\\\\\\\\\  /\\\\\\\\\\\  /\\\\\\\\\\\
     __|:|\:\  \   __|:|  |     \/////\\\///  \/////\\\///  \/////\\\///
    /::::|_\:\__\ /\ |:|__|____      \/\\\         \/\\\         \/\\\
    \:\~~\  \/__/ \:\/:::::/__/       \/\\\         \/\\\         \/\\\
     \:\  \        \::/~~/~            \/\\\         \/\\\         \/\\\
      \:\  \        \:\~~\              \/\\\         \/\\\         \/\\\
       \:\__\        \:\__\           /\\\\\\\\\\\  /\\\\\\\\\\\  /\\\\\\\\\\\
        \/__/         \/__/           \///////////  \///////////  \///////////


    ` + Yellow + `COMMUNITY EDITION` + Gray + `                                  by ` + White + `Studio 7 Development`

	Info string = Gray + `	Thanks for downloading! Sorry for the wait... the biggest challenge
	was figuring out how the hell we could bypass defender again since
	every C# payload was patched a few years back. You are using the 
	Community Edition of this software, which has been constructed with
	CyberSecurity Specialists and Pen-Testers in mind.
	
	visit ` + Cyan + `https://studio7.dev/mk3` + Gray + ` for more cool software.

								- S7 Dev Team

    Change Log [v1.0.0 ` + Green + `L` + Gray + `u` + Green + `c` + Gray + `k` + Green + `y ` + Yellow + `7` + Gray + `]:
      [` + Green + `+` + White + `] Initial Release
`
)
