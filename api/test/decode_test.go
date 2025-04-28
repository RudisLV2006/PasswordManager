package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"server/api/data_access"
	"testing"
)

func TestDecode(t *testing.T) {
	salt := "ApApqVBDi5pZ2Pqg4v9h2A=="
	expected := []byte{2, 144, 41, 169, 80, 67, 139, 154, 89, 216, 250, 160, 226, 255, 97, 216}

	decoded, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		t.Fatalf("Error decoding string: %v", err)
	}

	if !bytes.Equal(decoded, expected) {
		t.Errorf("Expected %v, but got %v", expected, decoded)
	}

}

func TestDecrypt(t *testing.T) {
	var expected string = "suns123"
	salt, err := base64.StdEncoding.DecodeString("UQSgnOat13uaJ/eqnxnvMA==")
	if err != nil {
		t.Fatalf("Error decoding string: %v", err)
	}
	t.Logf("saly: %v", salt)

	encrypted_password, err := base64.StdEncoding.DecodeString("nXTFJp65JwR3JWSl9HJM5F4zBo2nJhDR/cddDezhnBRBwho=")
	if err != nil {
		t.Fatalf("Error decoding string: %v", err)
	}
	t.Logf("Encrypted password: %v", encrypted_password)

	var secret_key string = "asd"

	key := data_access.DeriveEncryptionKey(secret_key, salt)
	t.Logf("Derived key: %v", key)

	var plaintext string = data_access.DecryptIt(encrypted_password, key)
	t.Logf("Plaintext: %v", plaintext)

	t.Logf("Expected: %v, got: %v", expected, plaintext)

	if plaintext != expected {
		t.Errorf("Expected %v, but got %v", expected, plaintext)
	}

}

func TestDecryptIt(t *testing.T) {
	var expected string = "rudis"

	// Decode salt and handle errors
	salt, err := base64.StdEncoding.DecodeString("Ad8g84TRiNZje1buvhzF9w==")
	if err != nil {
		t.Fatalf("Error decoding salt string: %v", err)
	}

	// Derive the encryption key using salt
	encryptedKey := data_access.DeriveEncryptionKey("test", salt)

	// Decode ciphered text and handle errors
	ciphered, err := base64.StdEncoding.DecodeString("hY/EKLVwMRiAslkg8YRGhLOs7L6rEXm1gUkqyWXl//ur")
	if err != nil {
		t.Fatalf("Error decoding ciphered string: %v", err)
	}

	// Create AES block cipher from the derived encryption key
	aesBlock, err := aes.NewCipher(encryptedKey)
	if err != nil {
		t.Fatalf("Error creating AES cipher: %v", err)
	}

	// Create GCM cipher using AES block cipher
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		t.Fatalf("Error creating GCM cipher: %v", err)
	}

	// Ensure the ciphered text is long enough to include the nonce
	nonceSize := gcmInstance.NonceSize()
	if len(ciphered) < nonceSize {
		t.Fatalf("Ciphered text is too short to contain a nonce of size %d", nonceSize)
	}

	// Extract the nonce and the encrypted ciphertext
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	// Decrypt the message using GCM
	plaintext, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		t.Fatalf("Error decrypting message: %v", err)
	}

	// Convert the decrypted message into a string
	plaintextStr := string(plaintext)

	// Log the decrypted text for debugging
	t.Logf("Decrypted text: %s", plaintextStr)

	// Check if the decrypted text matches the expected string
	if plaintextStr != expected {
		t.Errorf("Expected %v, but got %v", expected, plaintextStr)
	}
}
