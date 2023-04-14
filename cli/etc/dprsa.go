package etc

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// Generate a new RSA key pair with 2048 bits
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Save the private key to a local file
	SavePrivKey(privateKey, "private.key")

	// Load the private key from the local file
	loadedPrivateKey := LoadPrivKey("private.key")

	// Encrypt a message using the loaded private key
	message := "Hello, world!"
	encryptedMessage, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, &loadedPrivateKey.PublicKey, []byte(message), nil)

	// Print the encrypted message
	fmt.Println("Encrypted message:", base64.StdEncoding.EncodeToString(encryptedMessage))

	// Decrypt the message using the original private key
	decryptedMessage, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedMessage, nil)

	// Print the decrypted message
	fmt.Println("Decrypted message:", string(decryptedMessage))
}

func DecryptDB(pkey *rsa.PrivateKey, filename string) []byte {
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

func SavePrivKey(privateKey *rsa.PrivateKey, filename string) {
	// Create a PEM block for the private key
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Create the output file
	outputFile, _ := os.Create(filename)
	defer outputFile.Close()

	// Write the PEM block to the output file
	pem.Encode(outputFile, privateKeyBlock)
}

func LoadPrivKey(filename string) *rsa.PrivateKey {

	pemBytes, _ := os.ReadFile(filename)
	privateKeyBlock, _ := pem.Decode(pemBytes)

	// Parse the RSA private key from the PEM block
	privateKey, _ := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)

	return privateKey
}
