package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db       *sql.DB
	location string
}

func (this *Storage) Init(dbpath string) {
	this.db, _ = sql.Open("sqlite3", dbpath)
	this.location = dbpath
}

func (this *Storage) CreateTable(table_name string, headers map[string]string) {
	columns := "(id INTEGER PRIMARY KEY"
	for label, datatype := range headers {
		columns += ", " + label + " " + datatype
	}
	this.db.Exec("CREATE TABLE IF NOT EXISTS" + table_name + "(id INTEGER PRIMARY KEY, name TEXT)")
}

func (this *Storage) Store(table string, values []interface{}) {
	var value_list []string
	for _, value := range values {
		value_list = append(value_list, fmt.Sprintf("%v", value))
	}
	this.db.Exec("INSERT INTO " + table + " VALUES (" + strings.Join(value_list, ",") + ")")
}

func (this *Storage) SQL(query string) interface{} {
	rows, _ := this.db.Query(query)
	return rows
}

// Removes the databse from the system
func (this *Storage) Destroy() {
	this.db.Close()
	os.Remove(this.location)
}

// Signs (RSA Encrypts) & closes the database
func (this *Storage) Sign() {
	this.db.Close()

	// read the database
	raw, _ := os.ReadFile(this.location)

	// placeholder for RSA Big Int (assigned by builder)
	var rsa_n = big.NewInt(__BIGINT_x64)

	// placeholder for public key (assigned by builder)
	key := rsa.PublicKey{
		N: rsa_n,
		E: __RSA_E,
	}

	// Encrypt the raw using the public key
	ciphertext, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&key,
		raw,
		nil,
	)

	// Write the encrypted ciphertext to the output file
	os.WriteFile(this.location, ciphertext, 0644)

}

// Helper function to save an RSA private key to a file
func savePrivateKeyToFile(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the private key in PEM format to the file
	if err := pem.Encode(file, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}); err != nil {
		return err
	}

	return nil
}

// Helper function to load an RSA private key from a file
func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the private key from the file in PEM format
	pemBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Decode the PEM bytes to get the DER-encoded private key bytes
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	// Parse the DER-encoded private key bytes to get an RSA private key
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
