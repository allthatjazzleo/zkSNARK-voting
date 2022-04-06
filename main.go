package main

import (
	"crypto/sha256"
	"log"
	"math/big"

	"github.com/sethvargo/go-password/password"
)

func main() {
	// Generate a password that is 64 characters long with 10 digits, 0 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(64, 10, 0, false, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("secret: %v", res)
	input := []byte(res)
	input1 := new(big.Int)
	input2 := new(big.Int)
	input3 := new(big.Int)
	input4 := new(big.Int)
	input1.SetBytes(input[:16])
	input2.SetBytes(input[16:32])
	input3.SetBytes(input[32:48])
	input4.SetBytes(input[48:])
	log.Printf("input1: %v", input1)
	log.Printf("input2: %v", input2)
	log.Printf("input3: %v", input3)
	log.Printf("input4: %v", input4)

	// // generate hash
	hash := sha256.Sum256(input)
	output1 := new(big.Int)
	output2 := new(big.Int)
	output1.SetBytes(hash[:16])
	output2.SetBytes(hash[16:])

	log.Printf("output1: %v", output1)
	log.Printf("output2: %v", output2)
}
