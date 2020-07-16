package tests

import "testing"
import "github.com/kyberorg/go-utils/crypto/aesgcm"

func CryptoTest(t *testing.T) {
	cleanText := "Salasana"
	sharedSecret := "SharedSecret"

	encryptedText, err := aesgcm.EncryptString(cleanText, sharedSecret)
	if err != nil {
		t.Fail()
	}

	decryptedText, decryptError := aesgcm.DecryptString(encryptedText, sharedSecret)
	if decryptError != nil {
		t.Fail()
	}
	if decryptedText != cleanText {
		t.Fail()
	}
}
