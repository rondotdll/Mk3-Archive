package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/tawesoft/golib/v2/dialog"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

func GetSysInfo() *SysInfo {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	System := new(SysInfo)

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

func GetExternIP() string {
	resp, _ := http.Get("https://myexternalip.com/raw")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func DisplayErrorMsg(title string, description string) {
	dialog.Message{
		Title:  title,
		Format: description,
		Icon:   dialog.IconError,
	}.Raise()
}
