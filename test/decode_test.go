package test

import (
	"bytes"
	"encoding/base64"
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
