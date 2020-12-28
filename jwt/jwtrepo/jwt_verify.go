package jwtrepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/user/usermodel"
	"fooddlv/token"
)

type JwtVerifyStorage interface {
	FindUserByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*usermodel.User, error)
}

type jwtVerifyRepo struct {
	store JwtVerifyStorage
}

func NewJwtVerifyRepo(store JwtVerifyStorage) *jwtVerifyRepo {
	return &jwtVerifyRepo{store: store}
}

func (repo *jwtVerifyRepo) Validate(ctx context.Context, payload *token.JwtPayload) (*common.SimpleUser, error) {
	user, err := repo.store.FindUserByCondition(ctx, map[string]interface{}{"id": payload.UserId})


	if ok := user.IsActive(); !ok {
		return nil,common.NewUnauthorized(err, "user is deactivate", "ErrUserIsNotActive")
	}

	simpleUser := user.ToSimpleUser()
	return simpleUser, nil
}
