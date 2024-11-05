package token_test

import (
	"crypto/rand"
	"crypto/rsa"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/lucianonooijen/capsa/server/internal/entities"
	"github.com/lucianonooijen/capsa/server/internal/infrastructure/token"
)

func generateRsaKeyset(t *testing.T) *rsa.PrivateKey {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)

	require.NoError(t, err)
	require.NotEmpty(t, privateKey)

	return privateKey
}

func getTestTokenInstance(t *testing.T) *token.Token {
	t.Helper()

	c := &entities.Config{
		IsDevMode:           true,
		ServerPort:          5000,
		ServerHostname:      "localhost",
		LogMaxDurationHours: 48,
		RootLogger:          zap.NewNop(),
	}

	k := generateRsaKeyset(t)

	tok, err := token.New(c, k)

	require.NoError(t, err)

	return tok
}

func TestJwtFlow_Client_HappyPath(t *testing.T) {
	tok := getTestTokenInstance(t)
	sub := "TestSubjectHappyClient"

	key, err := tok.GenerateClientJwt(sub)

	require.NoError(t, err)
	require.NotEmpty(t, key)

	claims, err := tok.ValidateJwt(key)

	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.Equal(t, sub, claims.Subject)
}

func TestJwtFlow_User_HappyPath(t *testing.T) {
	tok := getTestTokenInstance(t)
	sub := "TestSubjectHappyUser"
	tid := "uuid"
	name := "TestName"

	key, err := tok.GenerateUserJwt(sub, tid, name)

	require.NoError(t, err)
	require.NotEmpty(t, key)

	claims, err := tok.ValidateJwt(key)

	require.NoError(t, err)
	require.NotEmpty(t, claims)
	require.Equal(t, sub, claims.Subject)
	require.Equal(t, tid, claims.JwtID)
	require.Equal(t, name, claims.Name)
}

func TestJwt_SignedByOtherKey(t *testing.T) {
	tok := getTestTokenInstance(t)
	sub := "TestSubjectOtherKey"

	key, err := tok.GenerateClientJwt(sub)

	require.NoError(t, err)
	require.NotEmpty(t, key)

	tokOtherKey := getTestTokenInstance(t)

	claims, err := tokOtherKey.ValidateJwt(key)

	require.Error(t, err)
	require.Nil(t, claims)
}

func TestJwt_InvalidSignature(t *testing.T) {
	tok := getTestTokenInstance(t)
	sub := "TestSubjectInvalidSignature"

	key, err := tok.GenerateClientJwt(sub)

	require.NoError(t, err)
	require.NotEmpty(t, key)

	keyParts := strings.Split(key, ".")
	exampleSignature := "VKPicz1jQzeysLyvjPxAJAJYzc0zHFVuMqabop9ovxc"
	keyParts[2] = exampleSignature
	keyWithInvalidSignature := strings.Join(keyParts, ".")

	claims, err := tok.ValidateJwt(keyWithInvalidSignature)

	require.Error(t, err)
	require.Nil(t, claims)
}

func TestJwt_ExpiredKey(t *testing.T) {
	tok := getTestTokenInstance(t)
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	twoMonthAgo := now.AddDate(0, -2, 0)

	claims := token.JwtClaims{
		Issuer:    "tests",
		Subject:   "testsub",
		Audience:  "gotest",
		Expiry:    oneMonthAgo.Unix(),
		NotBefore: twoMonthAgo.Unix(),
		IssuedAt:  twoMonthAgo.Unix(),
	}

	expiredKey, err := tok.GenerateTokenForClaims(claims)

	require.NoError(t, err)
	require.NotEmpty(t, expiredKey)

	claimsOut, err := tok.ValidateJwt(expiredKey)

	require.Error(t, err)
	require.Nil(t, claimsOut)
}

func TestJwt_NotBeforeInFuture(t *testing.T) {
	tok := getTestTokenInstance(t)
	now := time.Now()
	oneMonthFromNow := now.AddDate(0, 1, 0)
	twoMonthsFromNow := now.AddDate(0, 2, 0)

	claims := token.JwtClaims{
		Issuer:    "tests",
		Subject:   "testsub",
		Audience:  "gotest",
		Expiry:    twoMonthsFromNow.Unix(),
		NotBefore: oneMonthFromNow.Unix(),
		IssuedAt:  now.Unix(),
	}

	expiredKey, err := tok.GenerateTokenForClaims(claims)

	require.NoError(t, err)
	require.NotEmpty(t, expiredKey)

	claimsOut, err := tok.ValidateJwt(expiredKey)

	require.Error(t, err)
	require.Nil(t, claimsOut)
}
