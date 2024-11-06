package token

import (
	"errors"
	"time"
)

var (
	// ErrorJwtInvalidIssuer indicates an invalid Issuer.
	ErrorJwtInvalidIssuer = errors.New("token issuer is invalid")

	// ErrorJwtInvalidAudience indicates that the audience is not valid.
	ErrorJwtInvalidAudience = errors.New("token audience is invalid")

	// ErrorJwtIncorrectJwtID indicates that the jwt id value does not match the expected value.
	// This happens when a user has changed their password and attempts to use an old token.
	ErrorJwtIncorrectJwtID = errors.New("jwt id does not match expected value")

	// ErrorJwtExpired indicates that the token is expired.
	ErrorJwtExpired = errors.New("token is expired")

	// ErrorJwtNotValidYet indicates that the token is not yet valid.
	ErrorJwtNotValidYet = errors.New("token is not valid yet")
)

// Example for such logic: https://github.com/go-jose/go-jose/blob/v4.0.4/jwt/validation.go#L61
func validateJwtClaims(claims *JwtClaims) error {
	now := time.Now().Unix()

	if claims.Issuer != Issuer {
		return ErrorJwtInvalidIssuer
	}

	if claims.Audience != AudienceClient && claims.Audience != AudienceUser {
		return ErrorJwtInvalidAudience
	}

	if claims.Expiry < now {
		return ErrorJwtExpired
	}

	if claims.NotBefore > now {
		return ErrorJwtNotValidYet
	}

	return nil
}
