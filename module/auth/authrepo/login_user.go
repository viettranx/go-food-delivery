package authrepo

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/hash"
	"fooddlv/module/auth/authmodel"
	"fooddlv/module/user/usermodel"
	"fooddlv/token"
	"time"
)

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
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
		return nil, common.ErrEntityExisted(authmodel.EntityName, err)
	}

	md5Hash := hash.NewMd5Hash(loginUserData.Password, user.Salt)

	if ok := user.ComparePassword(md5Hash); !ok {
		return nil, ErrUsernameOrPasswordInvalid
	}

	accessToken, err := repo.tokenProvider.Generate(*user, token.WithExpiry(24*30*time.Hour))
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := repo.tokenProvider.Generate(*user, token.WithExpiry(24*60*time.Hour))
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := authmodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
