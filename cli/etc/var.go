package etc

var (
	EnabledFeatures []FeatureSet
)

const (
	Reset         = "\033[0m"
	Red           = "\033[31m"
	Green         = "\033[32m"
	Yellow        = "\033[33m"
	Blue          = "\033[34m"
	Purple        = "\033[35m"
	Cyan          = "\033[36m"
	Gray          = "\033[37m"
	White         = "\033[97m"
	Splash string = Green + `         ___           ___
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
