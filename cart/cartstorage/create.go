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
func (store *cartMysql) Create(ctx context.Context, cartCreateData *cartmodel.CartCreation) (int, error) {
	// init db
	db := store.db.Begin()
	// create data to db
	fmt.Println("create cart", cartCreateData)
	if err := db.Table(cartmodel.Cart{}.TableName()).Create(&cartCreateData).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	return cartCreateData.ID, nil
}
