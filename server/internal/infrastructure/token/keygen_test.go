package token_test

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"testing"

	"github.com/capsa-gg/capsa/server/internal/infrastructure/token"
)

// TestGenerateRsaKeySet tests the GenerateRsaKeySet function.
func TestGenerateRsaKeySet(t *testing.T) {
	keySet, err := token.GenerateRsaKeySet()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(keySet.PrivateKey) == 0 {
		t.Error("Expected private key to be generated")
	}

	if len(keySet.PublicKey) == 0 {
		t.Error("Expected public key to be generated")
	}
}

// TestEncryptDecrypt tests the ability to encrypt and decrypt data using the generated key pair.
func TestEncryptDecrypt(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	publicKey := &privateKey.PublicKey

	// Test data
	message := []byte("This is a secret message")

	// Create a new SHA256 hash
	hash := sha256.New()

	// Encrypt the message with the public key
	encryptedMessage, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, nil)
	if err != nil {
		t.Fatalf("Expected no error during encryption, got %v", err)
	}

	// Decrypt the message with the private key
	decryptedMessage, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, encryptedMessage, nil)
	if err != nil {
		t.Fatalf("Expected no error during decryption, got %v", err)
	}

	// Verify the decrypted message matches the original message
	if !bytes.Equal(decryptedMessage, message) {
		t.Errorf("Expected decrypted message to be %s, got %s", message, decryptedMessage)
	}
}
