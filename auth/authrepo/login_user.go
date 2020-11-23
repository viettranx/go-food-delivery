package authrepo

import (
	"context"
	"errors"
	"fooddlv/auth/authmodel"
	"fooddlv/common"
	"fooddlv/hash"
	"fooddlv/token"
	"fooddlv/user/usermodel"
	"time"
)

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(errors.New("Username or Password invalid"), "", "")
)

type LoginUserStorage interface {
	FindUserByCondition(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error)
}

type loginUserRepo struct {
	store         LoginUserStorage
	tokenProvider token.Provider
}

func NewLoginUserRepo(store LoginUserStorage, tokProvider token.Provider) *loginUserRepo {
	return &loginUserRepo{
		store:         store,
		tokenProvider: tokProvider,
	}
}

func (repo *loginUserRepo) LoginUser(ctx context.Context, loginUserData *authmodel.LoginUser) (*authmodel.Account, error) {
	user, err := repo.store.FindUserByCondition(ctx, map[string]interface{}{"email": loginUserData.Email})

	if err != nil {
		return nil, common.ErrUserAlreadyExists(authmodel.EntityName, err)
	}

	md5Hash := hash.NewMd5Hash(loginUserData.Password, user.Salt)

	if ok := user.ComparePassword(md5Hash); !ok {
		return nil, ErrUsernameOrPasswordInvalid
	}

	accessToken, err := repo.tokenProvider.Generate(*user, token.WithExpiry(15*time.Minute))
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := repo.tokenProvider.Generate(*user, token.WithExpiry(48*time.Hour))
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := authmodel.NewAccount(user, accessToken, refreshToken)

	return account, nil
}
