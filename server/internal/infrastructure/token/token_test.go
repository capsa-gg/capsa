package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/capsa-gg/capsa/server/internal/entities"
)

func TestGenerateWellKnownEndpoint_Dev(t *testing.T) {
	c := &entities.Config{
		IsDevMode:      true,
		ServerHostname: "localhost",
		ServerPort:     5000,
	}

	got := generateWellKnownEndpoint(c)
	expected := "http://localhost:5000/.well-known/jwks.json"

	assert.Equal(t, expected, got)
}

func TestGenerateWellKnownEndpoint_Prod(t *testing.T) {
	c := &entities.Config{
		IsDevMode:      false,
		ServerHostname: "capsa.gg",
		ServerPort:     8080,
	}

	got := generateWellKnownEndpoint(c)
	expected := "https://capsa.gg/.well-known/jwks.json"

	assert.Equal(t, expected, got)
}
