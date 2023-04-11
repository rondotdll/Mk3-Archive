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
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sql_type_conversion = map[string]string{
		"string":  "TEXT",
		"[]string": "TEXT",
		"int":     "INTEGER",
		"uint64":  "INTEGER",
		"[]uint8": "BLOB",
	}
)

type Vault struct {
	db       *sql.DB
	location string
}

type Table struct {
	name    string
	headers map[string]string
	rows    []interface{}
}

// convert various structs to a table so it can be stored
func ToTable(this interface{}) Table {
	var headers map[string]string
	var rows []interface{}

	t, v := reflect.TypeOf(this), reflect.ValueOf(this)

	contains_slice := (false, 0)

	// Iterate through the struct's fields and append their names to the list
	for i := 0; i < t.NumField(); i++ {
		if (t.Field(i).Type.Name() == "[]string"){
			contains_slice = true
		}
		headers[t.Field(i).Name] = sql_type_conversion[t.Field(i).Type.Name()]
	}

	for i := 0; i < v.NumField(); i++ {
		rows = append(rows, v.Field(i).Interface())
	}

	return Table{
		name:    strings.ToLower(t.Name()),
		headers: headers,
		rows: rows,
	}
}

func (this *Vault) Init(dbpath string) *Vault {
	this.db, _ = sql.Open("sqlite3", dbpath)
	this.location = dbpath
	return this
}

func (this *Vault) PushTable(table Table) *Vault {
	columns := "(id INTEGER PRIMARY KEY"
	for label, datatype := range table.headers {
		columns += ", " + label + " " + datatype
	}
	this.db.Exec("CREATE TABLE IF NOT EXISTS" + table.name + "(id INTEGER PRIMARY KEY, name TEXT)")
	return this
}

func (this *Vault) Store(table string, values []interface{}) *Vault {
	var value_list []string
	for _, value := range values {
		value_list = append(value_list, fmt.Sprintf("%v", value))
	}
	this.db.Exec("INSERT INTO " + table + " VALUES (" + strings.Join(value_list, ",") + ")")
	return this
}

func (this *Vault) SQL(query string) interface{} {
	rows, _ := this.db.Query(query)
	return rows
}

// Removes the databse from the system
func (this *Vault) Destroy() *Vault {
	this.db.Close()
	os.Remove(this.location)
	return this
}

// Signs (RSA Encrypts) & closes the database
func (this *Vault) Sign() *Vault {
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

	return this
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
