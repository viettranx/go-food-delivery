package authmodel

import (
	"fooddlv/token"
	"fooddlv/user/usermodel"
)

type Account struct {
	User         *usermodel.User `json:",inline,user"`
	AccessToken  *token.Token    `json:",inline,access_token"`
	RefreshToken *token.Token    `json:",inline,refresh_token"`
}

func NewAccount(user *usermodel.User, atok, rtok *token.Token) *Account {
	return &Account{
		User:         user, // remove
		AccessToken:  atok,
		RefreshToken: rtok,
	}
}
