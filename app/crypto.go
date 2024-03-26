package app

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"grest.dev/grest"
)

// Crypto returns a pointer to the cryptoUtil instance (crpto).
// If crpto is not initialized, it creates a new cryptoUtil instance, configures it, and assigns it to crpto.
// It ensures that only one instance of cryptoUtil is created and reused.
func Crypto() *cryptoUtil {
	if crpto == nil {
		crpto = &cryptoUtil{}
		crpto.configure()
	}
	return crpto
}

// crpto is a pointer to a cryptoUtil instance.
// It is used to store and access the singleton instance of cryptoUtil.
var crpto *cryptoUtil

// cryptoUtil represents a crypto utility.
// It embeds grest.Crypto, indicating that cryptoUtil inherits from grest.Crypto.
type cryptoUtil struct {
	grest.Crypto
}

// configure configures the crypto utility instance.
// It sets the encryption key (c.Key), salt (c.Salt), info (c.Info), and JWT key (c.JWTKey) to the corresponding environment variables.
func (c *cryptoUtil) configure() {
	c.Key = CRYPTO_KEY
	c.Salt = CRYPTO_SALT
	c.Info = CRYPTO_INFO
	c.JWTKey = JWT_KEY
}

// NewToken generates a new token using UUID.
// It replaces all dashes in the UUID string with an empty string and returns the resulting token.
func (c *cryptoUtil) NewToken() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}

// NewCrypto creates a new cryptoUtil instance with custom keys.
// It initializes the instance, configures it, and assigns the custom keys (if provided) to the corresponding fields (c.Key, c.Salt, c.Info, c.JWTKey).
// It returns the created cryptoUtil instance.
func NewCrypto(keys ...string) *cryptoUtil {
	c := &cryptoUtil{}
	c.configure()
	if len(keys) > 0 {
		c.Key = keys[0]
	}
	if len(keys) > 1 {
		c.Salt = keys[1]
	}
	if len(keys) > 2 {
		c.Info = keys[2]
	}
	if len(keys) > 3 {
		c.JWTKey = keys[3]
	}
	return c
}

// RegisteredJWTClaim represents registered claim names in the IANA "JSON Web Token Claims" registry.
//
// See: https://tools.ietf.org/html/rfc7519#section-4.1
type RegisteredJWTClaim struct {

	// ID claim provides a unique identifier for the JWT.
	ID string `json:"jti,omitempty"`

	// Audience claim identifies the recipients that the JWT is intended for.
	Audience string `json:"aud,omitempty"`

	// Issuer claim identifies the principal that issued the JWT.
	Issuer string `json:"iss,omitempty"`

	// Subject claim identifies the principal that is the subject of the JWT.
	Subject string `json:"sub,omitempty"`

	// IssuedAt claim identifies the time at which the JWT was issued.
	// This claim can be used to determine the age of the JWT.
	IssuedAt NullUnixTime `json:"iat,omitempty"`

	// ExpiresAt claim identifies the expiration time on or after which the JWT MUST NOT be accepted for processing.
	ExpiresAt NullUnixTime `json:"exp,omitempty"`

	// NotBefore claim identifies the time before which the JWT MUST NOT be accepted for processing.
	NotBefore NullUnixTime `json:"nbf,omitempty"`
}

// IsValidExpiresAt reports whether a token isn't expired at a given time.
func (rc RegisteredJWTClaim) IsValidExpiresAt(now time.Time) bool {
	return !rc.ExpiresAt.Valid || rc.ExpiresAt.Time.After(now)
}

// IsValidNotBefore reports whether a token isn't used before a given time.
func (rc RegisteredJWTClaim) IsValidNotBefore(now time.Time) bool {
	return !rc.NotBefore.Valid || rc.NotBefore.Time.Before(now)
}

// IsValidIssuedAt reports whether a token was created before a given time.
func (rc RegisteredJWTClaim) IsValidIssuedAt(now time.Time) bool {
	return !rc.IssuedAt.Valid || rc.IssuedAt.Time.Before(now)
}

// IsValidAt reports whether a token is valid at a given time.
func (rc RegisteredJWTClaim) IsValidAt(now time.Time) bool {
	return rc.IsValidExpiresAt(now) && rc.IsValidNotBefore(now) && rc.IsValidIssuedAt(now)
}
