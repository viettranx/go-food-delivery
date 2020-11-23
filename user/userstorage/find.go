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
	db := store.db.Table(usermodel.User{}.TableName()).Where("status = 1")

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
