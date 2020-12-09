package jwthdl

import (
	"context"
	"fooddlv/token"
)

type JwtRepo interface {
	Validate(ctx context.Context, payload *token.JwtPayload) error
}

type jwtHdl struct {
	repo JwtRepo
}

func NewJwtHdl(repo JwtRepo) *jwtHdl {
	return &jwtHdl{
		repo: repo,
	}
}

func (hdl *jwtHdl) Validate(ctx context.Context, payload *token.JwtPayload) error {
	return hdl.repo.Validate(ctx, payload)
}
