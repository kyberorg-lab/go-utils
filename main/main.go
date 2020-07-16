package main

import (
	"fmt"
	"github.com/kyberorg/go-utils/crypto/aesgcm"
)

func main() {
	cryptoTest()
}

func cryptoTest() {
	cleanText := "Salasana"
	sharedSecret := "SharedSecret"

	fmt.Println("Clean Text: ", cleanText, "Shared secret (aka secret key password)", sharedSecret)

	encryptedText, encryptError := aesgcm.EncryptString(cleanText, sharedSecret)
	if encryptError != nil {
		fmt.Println(encryptError)
		panic("Encrypted Error")
	}
	fmt.Println("Encrypted text", encryptedText)

	decryptedText, decryptError := aesgcm.DecryptString(encryptedText, sharedSecret)
	if decryptError != nil {
		fmt.Println(decryptError)
		panic("Decrypt Error")
	}
	fmt.Println("Decrypted text", decryptedText)

	if decryptedText != cleanText {
		fmt.Println("Failed. Text mismatch", decryptedText, cleanText)
	}
}
