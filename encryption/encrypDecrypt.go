package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 13, 5}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

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

// VERSION 2
// Encrypt using AES-GCM
func EncryptV2(plaintext, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// Decrypt using AES-GCM
func DecryptV2(ciphertext, nonce []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// createAESGCM initializes AES-GCM using the provided key.
func createAESGCM(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

// EncryptV2 encrypts the plaintext using AES-GCM.
func EncryptV3(plaintext, key []byte) ([]byte, []byte, error) {
	aesGCM, err := createAESGCM(key)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	return aesGCM.Seal(nil, nonce, plaintext, nil), nonce, nil
}

// DecryptV2 decrypts the ciphertext using AES-GCM.
func DecryptV3(ciphertext, key []byte) ([]byte, error) {
	aesGCM, err := createAESGCM(key)
	if err != nil {
		return nil, err
	}
	_, nonce, _ := EncryptV3(ciphertext, key)
	return aesGCM.Open(nil, nonce, ciphertext, nil)
}
