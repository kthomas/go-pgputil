package pgputil

import (
	"reflect"
	"testing"
)

func TestPGPPubEncryptDecrypt(t *testing.T) {
	plaintext := []byte("encrypt me!!!")
	ciphertext, err := PGPPubEncrypt(plaintext)
	if err != nil {
		t.Errorf("PGPPubEncrypt() error = %v", err)
		return
	}

	decryptedPlaintext, _ := PGPPubDecrypt(ciphertext)
	if !reflect.DeepEqual(decryptedPlaintext, plaintext) {
		t.Errorf("PGPPubDecrypt() = %s, want %s", decryptedPlaintext, plaintext)
	}
}

func TestPGPPubSymmetricallyEncryptDecrypt(t *testing.T) {
	plaintext := []byte("encrypt me!!!")
	ciphertext, err := PGPPubSymmetricallyEncrypt(plaintext)
	if err != nil {
		t.Errorf("PGPPubSymmetricallyEncrypt() error = %v", err)
		return
	}

	decryptedPlaintext, _ := PGPPubDecrypt(ciphertext)
	if !reflect.DeepEqual(decryptedPlaintext, plaintext) {
		t.Errorf("PGPPubDecrypt() = %s, want %s", decryptedPlaintext, plaintext)
	}
}
