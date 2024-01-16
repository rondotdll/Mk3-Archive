package lib

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"os"
)

func DecryptLDB(pkey *rsa.PrivateKey, filename string) []byte {
	blob, _ := os.ReadFile(filename)

	raw_bytes, _ := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		pkey,
		blob,
		nil,
	)

	return raw_bytes
}

// These 2 functions are actually a universal "Dump To File" for any datatype
// replace "rsa.PrivateKey" with "any" to make it universal
// ...and replace the custom file endings
func LoadKeyFromFile(data rsa.PrivateKey, filename string) error {
	file, err := os.Open(filename + ".lvk")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	return err
}

func DumpKeyToFile(data rsa.PrivateKey, filename string) error {
	file, err := os.Create(filename + ".lvk")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	return nil
}
