# go-utils
Useful Go Utilities

#AES GCM (package aesgcm)

* func Encrypt(plainText, sharedKey []byte) ([]byte, error)

* func EncryptString(plainTextData, sharedKey string) (string, error)

* func Decrypt(encryptedData, sharedKey []byte) ([]byte, error)

* func DecryptString(encryptedData, sharedKey string) (string, error)

Inspired by (ITNEXT)[https://itnext.io/encrypt-data-with-a-password-in-go-b5366384e291?gi=c4b2d7a25be9]