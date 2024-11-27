package passhash_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/capsa-gg/capsa/server/internal/infrastructure/passhash"
)

func TestPasswordHasher_PlainTextToHash(t *testing.T) {
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash("password")
	require.NoError(t, err)
	require.NotEmpty(t, hash)
}

func TestPasswordHasher_ComparePassToHash_Correct(t *testing.T) {
	plaintextPassword := "geitje123"
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash(plaintextPassword)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = hasher.ComparePassToHash(plaintextPassword, hash)
	require.NoError(t, err)
}

func TestPasswordHasher_ComparePassToHash_Incorrect(t *testing.T) {
	hasher := passhash.New()
	hash, err := hasher.PlainTextToHash("correctpass")
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = hasher.ComparePassToHash("incorrectpass", hash)
	require.Error(t, err)
}
