package encipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// GenerateEncryptionKey generates a new random 128-bit (16-byte) encryption key
// and returns it as a hexadecimal string.
//
// Returns:
// - string: The generated encryption key as a hexadecimal string.
// - error: An error if there is a failure in generating the random key.
//
// Example usage:
//
//	key, err := GenerateEncryptionKey()
//	if err != nil {
//		log.Fatalf("Error generating encryption key: %v", err)
//	}
//	fmt.Println("Generated encryption key:", key)
func GenerateEncryptionKey() (string, error) {
	key := make([]byte, 16)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	hexKey := hex.EncodeToString(key)

	return hexKey
}

// Encryption key is a 32-byte key
func Encrypt(text string, encryptionKey string) (string, error) {
	key := []byte(encryptionKey) // Must be 32-bytes
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := []byte(text)
	cipherText := make([]byte, aes.BlockSize+len(b))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], b)

	return hex.EncodeToString(cipherText), nil
}

func Decrypt(cryptoText string, encryptionKey string) (string, error) {
	key := []byte(encryptionKey)
	cipherText, _ := hex.DecodeString(cryptoText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
