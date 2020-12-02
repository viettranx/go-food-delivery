package cartstorage

import (
	"context"
	"fmt"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

/**
Create a cart, return ID if can create a item, and 0 if error
*/

func (store *cartMysql) Create(ctx context.Context, cartsCreateData *[]cartmodel.Cart) (int, error) {
	// init db
	db := store.db.Begin()
	// create data to db
	fmt.Println("create cart", cartsCreateData)
	if err := db.Table(cartmodel.Cart{}.TableName()).Create(&cartsCreateData).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return 1, nil
}
