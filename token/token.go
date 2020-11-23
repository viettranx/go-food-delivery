package token

import (
	"errors"
	"fooddlv/common"
	"fooddlv/user/usermodel"
	"time"
)

var (
	ErrNotFound      = common.NewCustomError(errors.New("token not found"), "", "")
	ErrEncodingToken = common.NewCustomError(errors.New("error encoding the token"), "", "")
	ErrInvalidToken  = common.NewCustomError(errors.New("invalid token provided"), "", "")
)

type JwtPayload struct {
	UserId int `json:"user_id"`
	Role   string `json:"role"`
}

type Provider interface {
	Generate(user usermodel.User, opts ...GenerateOption) (*Token, error)
	Inspect(token string) (*JwtPayload, error)
	String() string
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  time.Time `json:"expiry"`
}
