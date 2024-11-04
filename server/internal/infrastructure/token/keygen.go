package token

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// RsaKeySet is a set of generated public and private keys.
type RsaKeySet struct {
	PrivateKey []byte
	PublicKey  []byte
}

// GenerateRsaKeySet generates a private/public key set for signing JWKs.
func GenerateRsaKeySet() (*RsaKeySet, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, fmt.Errorf("error generating private key: %w", err)
	}

	return EncodePrivateKeyToBytes(privateKey)
}

// EncodePrivateKeyToBytes takes in a *rsa.PrivateKey and encodes this to the private and public PEM key bytes.
func EncodePrivateKeyToBytes(privateKey *rsa.PrivateKey) (*RsaKeySet, error) {
	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	publicKeyBytes, err := encodePublicKeyToPem(privateKey)
	if err != nil {
		return nil, fmt.Errorf("error generating poblic key: %w", err)
	}

	keyData := RsaKeySet{
		PrivateKey: privateKeyBytes,
		PublicKey:  publicKeyBytes,
	}

	return &keyData, nil
}

func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get PKCS #1, ASN.1 DER format
	privateKeyContents := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyContents,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privateBlock)

	return privatePEM
}

func encodePublicKeyToPem(privateKey *rsa.PrivateKey) ([]byte, error) {
	// Extract the public key
	publicKey := &privateKey.PublicKey

	// Marshal the public key to PKIX, ASN.1 DER form
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, fmt.Errorf("error marshaling public key: %w", err)
	}

	// Create the PEM block
	publicKeyPEMBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	// Encode the PEM block to []byte
	publicKeyPEM := pem.EncodeToMemory(&publicKeyPEMBlock)

	return publicKeyPEM, nil
}
