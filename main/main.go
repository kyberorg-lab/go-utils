package main

import (
	"fmt"
	"github.com/kyberorg/go-utils/crypto/aesgcm"
	"github.com/kyberorg/go-utils/osutils"
	"os"
)

func main() {
	cryptoTest()
	getEnvTest()
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

func getEnvTest() {
	envKey := "MY_VAR"
	envValue := "myValue"
	defaultValue := "defaultValue"

	os.Setenv(envKey, envValue)

	valueFromGetEnv, exists := osutils.GetEnv(envKey, defaultValue)

	if !exists {
		fmt.Println("Failed. ENV", envKey, "was set, but func reported false")
	} else if valueFromGetEnv != envValue {
		fmt.Println("Failed. Value from GetEnv returns mismatch value. Original", envValue, "We have",
			valueFromGetEnv)
	}

	nonExistingEnv := "SOME_VAR"

	valueFromGetEnvNonExists, valExists := osutils.GetEnv(nonExistingEnv, defaultValue)
	if valExists {
		fmt.Println("Fail. Function reported exist on non-exist ENV")
	}
	if valueFromGetEnvNonExists != defaultValue {
		fmt.Println("Fail. Function should return default value")
	}
}
