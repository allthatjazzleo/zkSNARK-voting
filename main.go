package main

import (
	"crypto/sha256"
	"encoding/hex"
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
	hex_input := hex.EncodeToString(input[:])
	input1 := new(big.Int)
	input2 := new(big.Int)
	input3 := new(big.Int)
	input4 := new(big.Int)
	input1.SetString(hex_input[:32], 16)
	input2.SetString(hex_input[32:64], 16)
	input3.SetString(hex_input[64:96], 16)
	input4.SetString(hex_input[96:], 16)
	log.Printf("input1: %v", input1)
	log.Printf("input2: %v", input2)
	log.Printf("input3: %v", input3)
	log.Printf("input4: %v", input4)

	// // generate hash
	hash := sha256.Sum256(input)
	hex_hash := hex.EncodeToString(hash[:])
	output1 := new(big.Int)
	output2 := new(big.Int)
	output1.SetString(hex_hash[:32], 16)
	output2.SetString(hex_hash[32:], 16)

	log.Printf("output1: %v", output1)
	log.Printf("output2: %v", output2)
}
