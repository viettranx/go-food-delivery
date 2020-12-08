package cartstorage

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

func (store *cartMysql) Delete(ctx context.Context, cartId int, userId int) error {
	// init db
	// create data to db
	if err := store.db.Table(cartmodel.Cart{}.TableName()).Delete(&cartmodel.Cart{}, "user_id = ? AND food_id = ?", userId, cartId).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
