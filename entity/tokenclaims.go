package entity

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	User         string   `json:"user"`
	Subordinates []string `json:"subordinates"`
	BelongsToHR  bool     `json:"hr"`
	jwt.StandardClaims
}
