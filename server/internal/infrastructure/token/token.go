package token

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-jose/go-jose/v4"
	"go.uber.org/zap"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

const (
	// Do not change these values, the webapp relies on these as well.
	keyID     = "capsa-server-jwk"
	algorithm = jose.RS256
)

// Token generates JWTs and JWKs with the private key passed into it in New.
type Token struct {
	privateKey  *rsa.PrivateKey
	log         *zap.SugaredLogger
	jwk         *jose.JSONWebKey
	signer      *jose.Signer
	maxLogHours int
}

// New returns a Token instance after validating the *rsa.PrivateKey.
func New(c *entities.Config, pk *rsa.PrivateKey) (*Token, error) {
	if pk == nil {
		return nil, errors.New("private key argument is required")
	}

	if err := pk.Validate(); err != nil {
		return nil, fmt.Errorf("error validating private key: %w", err)
	}

	jwk := jose.JSONWebKey{
		Key:       pk,
		Use:       "sig",
		Algorithm: string(algorithm),
		KeyID:     keyID,
	}

	signerOptions := (&jose.SignerOptions{}).
		WithHeader("alg", algorithm).
		WithHeader("jku", generateWellKnownEndpoint(c))

	jwtSigner, err := jose.NewSigner(jose.SigningKey{
		Algorithm: algorithm,
		Key:       pk,
	}, signerOptions)

	if err != nil {
		return nil, fmt.Errorf("error creating jwt signer: %w", err)
	}

	log := c.RootLogger.Named("Token").Sugar()
	jwkInstance := Token{
		privateKey:  pk,
		log:         log,
		jwk:         &jwk,
		signer:      &jwtSigner,
		maxLogHours: c.LogMaxDurationHours,
	}

	return &jwkInstance, nil
}

// GetPublicKey takes the public part of the private key and marshals this to JSON for the .well-known/jwks.json endpoint.
func (t *Token) GetPublicKey() ([]byte, error) {
	return t.jwk.Public().MarshalJSON()
}

func generateWellKnownEndpoint(c *entities.Config) string {
	if c.IsDevMode {
		return "http://" + c.ServerHostname + ":" + strconv.Itoa(c.ServerPort) + "/.well-known/jwks.json"
	}

	return "https://" + c.ServerHostname + "/.well-known/jwks.json"
}
