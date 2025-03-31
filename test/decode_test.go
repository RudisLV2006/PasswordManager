package test

import (
	"bytes"
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
	var expected string = "my suped duper password"
	salt, err := base64.StdEncoding.DecodeString("kVAJG3VkG4punm7Lqe2ICQ==")
	if err != nil {
		t.Fatalf("Error decoding string: %v", err)
	}

	encrypted_password, err := base64.StdEncoding.DecodeString("0lSWxBnyTlpz/bsdnrn7oIv+S0+QYd9R8/16OIfvO9rFgGMrr1VUd2QJpI923op7K09I")
	if err != nil {
		t.Fatalf("Error decoding string: %v", err)
	}

	var secret_key string = "test"

	key := data_access.DeriveEncryptionKey(secret_key, salt)

	var plaintext string = data_access.DecryptIt(encrypted_password, key)

	t.Logf("Expected: %v, got: %v", expected, plaintext)

	if plaintext != expected {
		t.Errorf("Expected %v, but got %v", expected, plaintext)
	}

}
