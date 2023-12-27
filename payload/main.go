package main

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
	"strconv"
)

func main() {

	k, _ := rsa.GenerateKey(rand.Reader, 2048)

	var bigInt *big.Int = big.NewInt(16)

	println(bigInt)

	print("Public Key Modulus: ")
	println(k.PublicKey.N.Int64())
	println("Public Key Exponent: " + strconv.Itoa(k.PublicKey.E))

}
