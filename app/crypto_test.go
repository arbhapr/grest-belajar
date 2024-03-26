package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "afa86391a4794da1b45b64381bcb9374"
	encrypted, err := Crypto().Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Crypto().Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}
