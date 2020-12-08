package cartstorage

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

/**
Create a cart, return ID if can create a item, and 0 if error
*/

func (store *cartMysql) Create(ctx context.Context, cartsCreateData []*cartmodel.Cart) error {
	// init db
	db := store.db.Begin()
	// create data to db
	if err := db.Table(cartmodel.Cart{}.TableName()).Create(cartsCreateData).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
