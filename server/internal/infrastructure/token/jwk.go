package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/go-jose/go-jose/v4"
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/entities"
)

const keyID = "capsa-server-jwk"

// Jwk generates jwks with the private key passed into it in New.
type Jwk struct {
	privateKey *rsa.PrivateKey
	log        *zap.SugaredLogger
	jwk        *jose.JSONWebKey
}

// New returns a Jwk instance after validating the *rsa.PrivateKey.
func New(c *entities.Config, pk *rsa.PrivateKey) (*Jwk, error) {
	if pk == nil {
		return nil, fmt.Errorf("private key argument is required")
	}

	if err := pk.Validate(); err != nil {
		return nil, fmt.Errorf("error validating private key: %w", err)
	}

	jwk := jose.JSONWebKey{
		Key:       pk,
		Use:       "sig",
		Algorithm: string(jose.RS256),
		KeyID:     keyID,
	}

	log := c.RootLogger.Named("jwk").Sugar()
	jwkInstance := Jwk{
		privateKey: pk,
		log:        log,
		jwk:        &jwk,
	}

	return &jwkInstance, nil
}

// GetPublicKey takes the public part of the private key and marshals this to JSON for the .well-known/jwks.json endpoint.
func (jwk *Jwk) GetPublicKey() ([]byte, error) {
	return jwk.jwk.MarshalJSON()
}
