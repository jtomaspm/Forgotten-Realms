package models

import "github.com/golang-jwt/jwt/v4"

type NewUserClaims struct {
	jwt.RegisteredClaims
	ExternalId string `json:"external_id"`
	Source     string `json:"source"`
	Email      string `json:"email"`
}
