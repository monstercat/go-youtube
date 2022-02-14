package youtube

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/oauth2/jws"
	"golang.org/x/oauth2/jwt"
)

var (
	ErrInvalidPrivateKey = errors.New("invalid private key")
)

// ConvertServiceAccountJsonToJWT converts a service account JSON to JWT format.
func ConvertServiceAccountJsonToJWT(conf *jwt.Config) (string, error) {
	// Parse the PrivateKey
	key := conf.PrivateKey
	block, _ := pem.Decode(key)
	if block != nil {
		key = block.Bytes
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(key)
	if err != nil {
		parsedKey, err = x509.ParsePKCS1PrivateKey(key)
		if err != nil {
			return "", fmt.Errorf("private key should be a PEM or plain PKCS1 or PKCS8; parse error: %v", err)
		}
	}
	parsed, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", ErrInvalidPrivateKey
	}

	// Encode JWT
	header := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}
	claimSet := &jws.ClaimSet{
		Iss:           conf.Email,
		Scope:         strings.Join(conf.Scopes, " "),
		Aud:           conf.TokenURL,
		PrivateClaims: conf.PrivateClaims,
	}
	return jws.Encode(header, claimSet, parsed)
}
