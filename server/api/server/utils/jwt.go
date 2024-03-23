package utils

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"captioner.com.ng/config"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	UserID primitive.ObjectID
	Type   string
}
type JWTClaims struct {
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

	tok, err := jwt.NewBuilder().Claim("user_id", payload.UserID).Claim("type", payload.Type).Issuer(`api.captioner.com.ng`).IssuedAt(time.Now()).Expiration(time.Now().Add(365 * 24 * 24 * time.Hour)).Build()
	if err != nil {
		return "", err
	}
	signed, err := jwt.Sign(tok, jwa.RS256, privateKey)
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

	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	verifiedToken, err := jwt.Parse([]byte(token), jwt.WithVerify(jwa.RS256, publicKey))
	if err != nil {
		fmt.Println("Couldn't verify token")
		return nil, err
	}
	return verifiedToken, nil
}
