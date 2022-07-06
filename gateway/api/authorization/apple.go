package authorization

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/oauth2"
)

const (
	ENDPOINT_APPLE_AUTH                 = "https://appleid.apple.com/auth/authorize"
	ENDPOINT_APPLE_TOKEN                = "https://appleid.apple.com/auth/token"
	idTokenVerificationKeyAppleEndpoint = "https://appleid.apple.com/auth/keys"
	AppleAudOrIss                       = "https://appleid.apple.com"
)

type authorizationApple struct {
	method
	token    string
	userInfo *UserInfo
}

func NewAuthorizationApple(token string) Authorization {
	return &authorizationApple{
		token: token,
	}
}

func ToAppleUserInfo(data interface{}) *UserInfo {
	if s, ok := data.(*UserInfo); ok {
		return s
	}
	return nil
}

type IDTokenClaims struct {
	jwt.StandardClaims
	AccessTokenHash string `json:"at_hash"`
	AuthTime        int    `json:"auth_time"`
	Email           string `json:"email"`
	IsPrivateEmail  bool   `json:"is_private_email,string"`
}

func (auth *authorizationApple) SubmitAuth() error {

	config := &oauth2.Config{
		ClientID:     os.Getenv("APPLE_APP_ID"),
		ClientSecret: os.Getenv("APPLE_APP_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  ENDPOINT_APPLE_AUTH,
			TokenURL: ENDPOINT_APPLE_TOKEN,
		},
	}

	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("client_id", config.ClientID),
		oauth2.SetAuthURLParam("client_secret", config.ClientSecret),
	}

	token, err := config.Exchange(context.TODO(), auth.token, opts...)

	if err != nil {
		return err
	}

	if !token.Valid() {
		return errors.New("invalid token received from provider")
	}

	client := config.Client(context.Background(), token)

	if idToken := token.Extra("id_token"); idToken != nil {
		idToken, err := jwt.ParseWithClaims(idToken.(string), &IDTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
			kid := t.Header["kid"].(string)
			claims := t.Claims.(*IDTokenClaims)
			vErr := new(jwt.ValidationError)
			if !claims.VerifyAudience(config.ClientID, true) {
				vErr.Inner = fmt.Errorf("audience is incorrect")
				vErr.Errors |= jwt.ValidationErrorAudience
			}
			if !claims.VerifyIssuer(AppleAudOrIss, true) {
				vErr.Inner = fmt.Errorf("issuer is incorrect")
				vErr.Errors |= jwt.ValidationErrorIssuer
			}
			if vErr.Errors > 0 {
				return nil, vErr
			}

			hash := sha256.Sum256([]byte(token.AccessToken))
			halfHash := hash[0:(len(hash) / 2)]
			encodedHalfHash := base64.RawURLEncoding.EncodeToString(halfHash)
			if encodedHalfHash != claims.AccessTokenHash {
				vErr.Inner = fmt.Errorf(`identity token invalid`)
				vErr.Errors |= jwt.ValidationErrorClaimsInvalid
				return nil, vErr
			}

			set, err := jwk.Fetch(context.Background(), idTokenVerificationKeyAppleEndpoint, jwk.WithHTTPClient(client))
			if err != nil {
				return nil, err
			}
			selectedKey, found := set.LookupKeyID(kid)
			if !found {
				return nil, errors.New("could not find matching public key")
			}
			pubKey := &rsa.PublicKey{}
			err = selectedKey.Raw(pubKey)
			if err != nil {
				return nil, err
			}
			return pubKey, nil
		})
		if err != nil {
			return err
		}

		auth.userInfo = &UserInfo{
			ID:    idToken.Claims.(*IDTokenClaims).Subject,
			Name:  "",
			Email: idToken.Claims.(*IDTokenClaims).Email,
		}

		return nil
	}

	return errors.New("get userinfo fail")
}

func (auth *authorizationApple) GetData() *UserInfo     { return auth.userInfo }
func (auth *authorizationApple) GetServiceName() string { return SERVICE_MAME_APPLE }
