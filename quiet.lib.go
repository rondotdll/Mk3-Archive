package main

import (
	"fmt"
	"github.com/tawesoft/golib/v2/dialog"
	"io"
	"log"
	"net/http"
	"os/exec"
	"regexp"
)

func DisplayErrorMsg(title string, description string) {
	dialog.Message{
		Title:  title,
		Format: description,
		Icon:   dialog.IconError,
	}.Raise()
}

func GetExternIP() string {
	resp, _ := http.Get("https://myexternalip.com/raw")
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return string(body)
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
