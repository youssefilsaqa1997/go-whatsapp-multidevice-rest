package types

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthJWTClaims struct {
	JID string `json:"jid"`
	jwt.RegisteredClaims
}
