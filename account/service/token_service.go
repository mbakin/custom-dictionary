package service

import (
	"crypto/rsa"
	"github.com/mbakin/custom-dictionary/account/model"
)

// TokenService used for injecting an implementation of TokenRepository
// for use in service methods along with keys and secrets for
// signing JWTs
type TokenService struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// TSConfig will hold repositories that will eventually be injected into this
// this service layer
type TSConfig struct {
	// TokenRepository model.TokenRepository
	PrivKey       *rsa.PrivateKey
	PubKey        *rsa.PublicKey
	RefreshSecret string
}

// NewTokenService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewTokenService(c *TSConfig) model.TokenService {
	return &TokenService{
		PrivKey:       c.PrivKey,
		PubKey:        c.PubKey,
		RefreshSecret: c.RefreshSecret,
	}
}
