package main

import (
	"Mk3CLI/commands"
	. "Mk3CLI/etc"
	"fmt"
)

func main() {

	fmt.Println(Splash)
	fmt.Println("\n" + Info)
	fmt.Println("Try running 'help' for a list of commands.\n")

	for {
		commands.Handle()
	}
}

// saved for later:
// Save the private key to a file
//func main2() {
//  privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

//	privateKeyFile := "private.pem"
//	if err := savePrivateKeyToFile(privateKeyFile, privateKey); err != nil {
//		panic(err)
//	}
//
//	// Load the private key from the file
//	loadedPrivateKey, err := loadPrivateKeyFromFile(privateKeyFile)
//	if err != nil {
//		panic(err)
//	}
//
//	// Decrypt the ciphertext using the loaded private key
//	decryptedText, err := rsa.DecryptOAEP(
//		sha256.New(),
//		rand.Reader,
//		loadedPrivateKey,
//		ciphertext,
//		nil,
//	)
//
//	// Verify that the decrypted text matches the original raw
//	if string(decryptedText) != string(raw) {
//		panic("decrypted text does not match original raw")
//	}
//
//	fmt.Println("Encryption and decryption successful!")
//}
