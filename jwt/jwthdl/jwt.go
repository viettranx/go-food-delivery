package jwthdl

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/user/usermodel"
	"fooddlv/token"
)

type JwtRepo interface {
	Validate(ctx context.Context, payload *token.JwtPayload) (*common.SimpleUser, error)
}

type jwtHdl struct {
	repo JwtRepo
}

func NewJwtHdl(repo JwtRepo) *jwtHdl {
	return &jwtHdl{
		repo: repo,
	}
}

func (hdl *jwtHdl) Validate(ctx context.Context, payload *token.JwtPayload) (*common.SimpleUser, error) {
	user, err := hdl.repo.Validate(ctx, payload)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return user, nil
}
