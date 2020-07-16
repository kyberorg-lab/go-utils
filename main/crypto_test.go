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

	encryptedText, encryptError := aesgcm.EncryptString(cleanText, sharedSecret)
	if encryptError != nil {
		fmt.Println(encryptError)
		panic("Encrypted Error")
	}

	decryptedText, decryptError := aesgcm.DecryptString(encryptedText, sharedSecret)
	if decryptError != nil {
		fmt.Println(decryptError)
		panic("Decrypt Error")
	}
	if decryptedText != cleanText {
		t.Fail()
	}
}
