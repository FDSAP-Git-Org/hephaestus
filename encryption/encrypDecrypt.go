package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 13, 5}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

<<<<<<<<<<<<<<  ✨ Codeium Command ⭐  >>>>>>>>>>>>>>>>
// decodeBase64 decodes a given base64 string into a byte slice.
// Panics if the string is not a valid base64 string.
<<<<<<<  dc178924-238b-4cd3-8c39-7dc1ff3ef0b3  >>>>>>>
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeBase64(cipherText), nil
}

// Encrypt method is to encrypt or hide any classified request and repsonse body
func EncryptRequest(text interface{}, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	reqMS, _ := json.Marshal(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(reqMS))
	cfb.XORKeyStream(cipherText, reqMS)
	return encodeBase64(cipherText), nil
}

func CreateSeal(text interface{}, secretKey string) (string, error) {
	log.Println("SBODY:", text)
	msString, _ := json.Marshal(text)

	fmt.Println("MARSHAL:", string(msString))
	// GenerateHash(text)
	return "", nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
