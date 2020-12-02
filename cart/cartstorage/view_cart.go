package cartstorage

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

func (store *cartMysql) ViewCart(
	ctx context.Context,
	userId int,
) ([]cartmodel.Cart, error) {
	// define
	var cart []cartmodel.Cart
	// find the cart with where status = 1
	// SHOULD update to fine where user_id == userid :)
	db := store.db.Table(cartmodel.Cart{}.TableName()).Where("status = 1")

	// handle error
	if err := db.Find(&cart).Where("user_id = ?", userId).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//return
	return cart, nil
}
