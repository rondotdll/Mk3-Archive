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
		"string":   "TEXT",
		"[]string": "TEXT",
		"int":      "INTEGER",
		"uint64":   "INTEGER",
		"[]uint8":  "BLOB",
	}
)

type Vault struct {
	db       *sql.DB
	location string
}

type Table struct {
	name    string
	headers map[string]string
	rows    [][]interface{}
}

// Convert (most) structs and struct slices to a table so they can be stored
func ToTable(this interface{}) Table {
	output := Table{
		headers: make(map[string]string),
		rows:    make([][]interface{}, 0),
	}

	t, v := reflect.TypeOf(this), reflect.ValueOf(this)
	output.name = t.Name()

	// if the input object is a slice
	if output.name == "" && t.Kind() == reflect.Slice {
		// iterate through the outer slice (each row)
		for i := 0; i < v.Len(); i++ {
			row := make([]interface{}, 0)
			// iterate through row columns
			for n := 0; n < v.Index(i).Type().NumField(); n++ {
				// if the row is the first row
				if i == 0 {
					sqlType := sql_type_conversion[v.Index(0).Type().Field(n).Type.Name()]
					output.headers[v.Index(0).Type().Field(n).Name] = sqlType
				}

				println(v.Index(i).Field(n).Interface())
				// append the current column to the current row
				row = append(row, v.Index(i).Field(n).Interface())
			}
			output.rows = append(output.rows, row)
		}

	} else {
		output.rows = make([][]interface{}, 0)

		// iterate through the struct's property names
		for n := 0; n < v.Type().NumField(); n++ {
			sqlType := sql_type_conversion[v.Type().Field(n).Type.Name()]
			output.headers[v.Type().Field(n).Name] = sqlType
			output.rows[0] = append(output.rows[0], v.Interface())
		}
	}

	return output
}

func sformat(val interface{}) string {
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s ", value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d ", value.Int())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f ", value.Float())
	case reflect.Bool:
		return fmt.Sprintf("%t ", value.Bool())
	default:
		return fmt.Sprintf("%v ", value.Interface())
	}
}

func (this *Vault) Init(dbpath string) *Vault {
	this.db, _ = sql.Open("sqlite3", dbpath)
	this.location = dbpath
	return this
}

// Pushes and stores a Table in the database
func (this *Vault) StoreTable(table Table) *Vault {
	columns := "(id INTEGER PRIMARY KEY"
	for hname, htype := range table.headers {
		columns += ", " + hname + " " + htype
	}
	this.db.Exec("CREATE TABLE IF NOT EXISTS" + table.name + columns + ")")

	for _, row := range table.rows {
		var values []string
		for _, col := range row {
			values = append(values, sformat(col))
		}

		this.db.Exec("INSERT INTO " + table.name + " VALUES (" + strings.Join(values, ",") + ")")
	}
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
