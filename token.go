package appstoreconnect

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type claims struct {
	jwt.RegisteredClaims
	Scope []string `json:"scope,omitempty"`
}

type Token struct {
	v *jwt.Token
}

func NewToken(keyID string, issuerID string, lifetime time.Duration, scope []string) *Token {
	iat := time.Now()
	exp := iat.Add(lifetime)

	c := &claims{
		jwt.RegisteredClaims{
			Issuer:    issuerID,
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
			Audience:  []string{"appstoreconnect-v1"},
		},
		scope,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	t.Header["kid"] = keyID

	return &Token{t}
}

func (t *Token) Sign(p8key []byte) (string, error) {
	block, _ := pem.Decode(p8key)
	if block == nil {
		return "", fmt.Errorf("pem data not found")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	pk, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("not ecdsa private key")
	}

	sstr, err := t.v.SignedString(pk)
	if err != nil {
		return "", fmt.Errorf("failed to create signed string: %w", err)
	}

	return sstr, nil
}
