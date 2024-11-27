package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

const blockTypeRsaPrivateKey = "RSA PRIVATE KEY"

var (
	// ErrorPemPrivateKeyDecoding indicates that the decoding of the private-key containing PEM block failed.
	ErrorPemPrivateKeyDecoding = errors.New("failed to decode PEM block containing private key")
)

// LoadPrivateKeyFromPath reads the private key from a file and returns a rsa.PrivateKey struct.
func LoadPrivateKeyFromPath(path string) (*rsa.PrivateKey, error) {
	// Read the file
	keyBytes, err := os.ReadFile(path) //nolint:gosec // "G304: Potential file inclusion via variable", this path is set in the config.yml and should never contain user input
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != blockTypeRsaPrivateKey {
		return nil, ErrorPemPrivateKeyDecoding
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

// LoadPrivateKeyFromBase64String reads the private key from a base64 string and returns a rsa.PrivateKey struct.
func LoadPrivateKeyFromBase64String(b64 string) (*rsa.PrivateKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key base64 string: %w", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil || block.Type != blockTypeRsaPrivateKey {
		return nil, ErrorPemPrivateKeyDecoding
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}
