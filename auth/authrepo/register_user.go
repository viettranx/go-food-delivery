package authrepo

import (
	"context"
	"fooddlv/auth/authmodel"
	"fooddlv/common"
	"fooddlv/hash"
	"fooddlv/randx"
	"fooddlv/user/usermodel"
)

type RegisterStorage interface {
	FindUserByCondition(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error)
	Create(ctx context.Context, createUserData *usermodel.CreateUser) (int, error)
}

type registerRepo struct {
	store RegisterStorage
}

func NewAuthRepo(store RegisterStorage) *registerRepo {
	return &registerRepo{store: store}
}

func (repo *registerRepo) Register(ctx context.Context, createUserData *authmodel.CreateUser) (userId int, err error) {
	user, err := repo.store.FindUserByCondition(ctx, map[string]interface{}{"email": createUserData.Email})

	if user != nil {
		return 0, common.ErrEntityExisted(authmodel.EntityName, err)
	}

	md5Hash := hash.NewMd5Hash(createUserData.Password, randx.GenSalt(50))

	userId, err = repo.store.Create(ctx, createUserData.ToCreateUser(md5Hash))

	if err != nil {
		return 0, common.ErrCannotCreateEntity(authmodel.EntityName, err)
	}

	return userId, nil

}
