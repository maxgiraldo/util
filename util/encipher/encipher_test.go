package encipher

import (
	"testing"
)

// TestEncryptionAndDecryption tests both encryption and decryption functions.
func TestEncryptionAndDecryption(t *testing.T) {
	encryptionKey := "1234567890abcdef1234567890abcdef" // 32-byte key for AES-128
	plainText := "Hello, World!"

	encryptedText, err := Encrypt(plainText, encryptionKey)
	if err != nil {
		t.Fatalf("Failed to encrypt text: %v", err)
	}

	decryptedText, err := Decrypt(encryptedText, encryptionKey)
	if err != nil {
		t.Fatalf("Failed to decrypt text: %v", err)
	}

	if decryptedText != plainText {
		t.Errorf("Decrypted text does not match original text. Got %s, want %s", decryptedText, plainText)
	}
}

// TestShortCipherText tests the decryption function with a short cipher text.
func TestShortCipherText(t *testing.T) {
	encryptionKey := "1234567890abcdef1234567890abcdef" // 32-byte key for AES-128
	shortCipherText := "short"

	_, err := Decrypt(shortCipherText, encryptionKey)
	if err == nil {
		t.Fatalf("Expected error for short cipher text, but got none")
	}
}

// TestInvalidKey tests the encryption and decryption functions with an invalid key length.
func TestInvalidKey(t *testing.T) {
	invalidKey := "shortkey" // Invalid key length
	plainText := "Hello, World!"

	_, err := Encrypt(plainText, invalidKey)
	if err == nil {
		t.Fatalf("Expected error for invalid key length, but got none")
	}

	encryptedText := "d0e9f30aa4f0fa5e6e1a8e3b58b9f6e7" // Some example encrypted text
	_, err = Decrypt(encryptedText, invalidKey)
	if err == nil {
		t.Fatalf("Expected error for invalid key length, but got none")
	}
}

// TestEmptyText tests the encryption and decryption functions with an empty text.
func TestEmptyText(t *testing.T) {
	encryptionKey := "1234567890abcdef1234567890abcdef" // 32-byte key for AES-128
	emptyText := ""

	encryptedText, err := Encrypt(emptyText, encryptionKey)
	if err != nil {
		t.Fatalf("Failed to encrypt text: %v", err)
	}

	decryptedText, err := Decrypt(encryptedText, encryptionKey)
	if err != nil {
		t.Fatalf("Failed to decrypt text: %v", err)
	}

	if decryptedText != emptyText {
		t.Errorf("Decrypted text does not match original text. Got %s, want %s", decryptedText, emptyText)
	}
}
