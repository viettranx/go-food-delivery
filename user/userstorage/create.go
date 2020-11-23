package userstorage

import (
	"context"
	"fmt"
	"fooddlv/common"
	"fooddlv/user/usermodel"
)

func (store *userMySql) Create(ctx context.Context, createUserData *usermodel.CreateUser) (int, error) {
	db := store.db.Begin()

	fmt.Println("create user data", createUserData)

	if err := db.Table(usermodel.CreateUser{}.TableName()).Create(&createUserData).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	fmt.Println("after create user data", createUserData)

	return createUserData.ID, nil
}