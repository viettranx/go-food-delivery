package authmodel

import (
	"fooddlv/token"
)

type Account struct {
	AccessToken  *token.Token    `json:"access_token"`
	RefreshToken *token.Token    `json:"refresh_token"`
}

func NewAccount(atok, rtok *token.Token) *Account {
	return &Account{
		AccessToken:  atok,
		RefreshToken: rtok,
	}
}
