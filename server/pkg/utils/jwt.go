package utils

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"

	"captioner.com.ng/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_TOKEN_EXPIRE = 1 * 60
	REFRESH_TOKEN_EXPIRE = 30 * 60
	EMAIL_TOKEN_EXPIRE   = 5 * 60
	RESET_TOKEN_EXPIRE   = 5 * 60

	REFRESH_TOKEN = "user_refresh_token"
	SESSION_TOKEN = "user_session_token"
	EMAIL_TOKEN   = "verify_email_token"
	RESET_TOKEN   = "password_reset_token"
)

type JWTPayload struct {
	UserID interface{}
	Email  string
	Type   string
}

var cfg, _ = config.NewConfig()

func CreateToken(payload JWTPayload, expire time.Time) (string, error) {

	key, err := base64.StdEncoding.DecodeString(cfg.PrivateKey)
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(key)

	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	issuer, err := bcrypt.GenerateFromPassword([]byte("api.captioner.com.ng"), 10)
	if err != nil {
		return "", err
	}

	tok, err := jwt.NewBuilder().Claim("user_id", payload.UserID).Claim("email", payload.Email).Claim("type", payload.Type).Issuer(string(issuer)).IssuedAt(time.Now()).Expiration(expire).Build()
	if err != nil {
		return "", err
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

func VerifyToken(token string) (jwt.Token, error) {
	key, err := base64.StdEncoding.DecodeString(cfg.PublicKey)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, publicKey))
	if err != nil {
		return nil, err
	}
	return verifiedToken, nil
}
