package utils

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"captioner.com.ng/config"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	UserID primitive.ObjectID
	Type   string
}

func CreateToken(payload JWTPayload, expire time.Time) (string, error) {

	key, err := base64.StdEncoding.DecodeString(config.EnvPrivateKey())
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(key)

	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	tok, err := jwt.NewBuilder().Claim("user_id", payload.UserID).Claim("type", payload.Type).Issuer(`api.captioner.com.ng`).IssuedAt(time.Now()).Expiration(expire).Build()
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
	key, err := base64.StdEncoding.DecodeString(config.EnvPublicKey())

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(key)

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("Couldn't Load Pubkey")
		return nil, err
	}

	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.RS256, publicKey))
	if err != nil {
		fmt.Println("Couldn't verify token")
		return nil, err
	}
	return verifiedToken, nil
}
