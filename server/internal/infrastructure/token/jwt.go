package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

const (
	// AudienceClient is the audience for log senders.
	AudienceClient = "capsa-client"

	// AudienceUser is the audience for log readers.
	// Note: do not change this value, the webapp depends on this for middleware.
	AudienceUser = "capsa-user"

	// Issuer is the identification of the server when signing JWTs.
	// Note: do not change this value, the webapp depends on this for middleware.
	Issuer = "capsa-server"
)

var (
	// ErrorJwtParsing indicates that the signature of the JWT is not valid.
	ErrorJwtParsing = errors.New("parsing JWT failed")

	// ErrorJwtValidation indicates that the token claims could not be validated.
	ErrorJwtValidation = errors.New("claim validation for JWT failed")

	// ErrorJwtConversion indicates that the token claims to struct conversion failed.
	ErrorJwtConversion = errors.New("claim conversion for JWT failed")
)

// GenerateClientJwt generates a JWT for clients (log senders).
func (t *Token) GenerateClientJwt(subject string) (string, error) {
	now := time.Now()
	expiryHours := time.Duration(t.maxLogHours) * time.Hour

	claims := JwtClaims{
		Issuer:    Issuer,
		Subject:   subject,
		Audience:  AudienceClient,
		Expiry:    now.Add(expiryHours).Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
	}

	tok, err := t.GenerateTokenForClaims(claims)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	t.log.Infof("generated client jwt with subject: %s", subject)

	return tok, nil
}

// GenerateUserJwt generates a JWT for users (log readers).
func (t *Token) GenerateUserJwt(subject, tokenID, name, role string) (string, error) {
	now := time.Now()
	oneMonthFromNow := now.AddDate(0, 1, 0)

	claims := JwtClaims{
		Issuer:    Issuer,
		Subject:   subject,
		Audience:  AudienceUser,
		Expiry:    oneMonthFromNow.Unix(),
		NotBefore: now.Unix(),
		IssuedAt:  now.Unix(),
		JwtID:     tokenID,
		Name:      name,
		Role:      role,
	}

	tok, err := t.GenerateTokenForClaims(claims)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	t.log.Infof("generated user jwt with subject: %s and token id: %s", subject, tokenID)

	return tok, nil
}

// GenerateTokenForClaims is a function that will sign the JwtClaims passed in.
// WARNING: This method should only be used by this module itself and tests for this module.
// For generating tokens in production code, always use the audience-specific methods.
func (t *Token) GenerateTokenForClaims(claims JwtClaims) (string, error) { //nolint:gocritic // Jose needs val, not ref
	return jwt.Signed(*t.signer).Claims(claims).Serialize()
}

// ValidateJwt validates a jwt string with the public key on the server and validates claims.
func (t *Token) ValidateJwt(token string) (*JwtClaims, error) {
	parsedToken, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{algorithm})
	if err != nil {
		t.log.Warnf("error parsing signed token: %s", err)

		return nil, ErrorJwtParsing
	}

	claims := Claims{}

	// Note: dereference here is very important!
	// The jose code checks for *rsa.PublicKey specifically, and does not accept rsa.PublicKey
	err = parsedToken.Claims(&t.privateKey.PublicKey, &claims)
	if err != nil {
		t.log.Warnf("error validating token claims: %s", err)

		return nil, ErrorJwtValidation
	}

	c, err := claims.jwtClaims()
	if err != nil {
		t.log.Warnf("error converting token claims: %s", err)

		return nil, ErrorJwtConversion
	}

	err = validateJwtClaims(c)
	if err != nil {
		t.log.Warnf("error validating token claims: %s", err)

		return nil, err
	}

	return c, nil
}
