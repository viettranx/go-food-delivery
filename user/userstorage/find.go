package userstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/user/usermodel"
)

func (store *userMySql) FindUserByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	relations ...string) (*usermodel.User, error) {

	var user usermodel.User
	db := store.db.Table(usermodel.User{}.TableName())

	if err := db.First(&user).Where(conditions).Where("status = 1").Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
