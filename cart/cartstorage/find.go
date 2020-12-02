package cartstorage

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

func (store *cartMysql) FindCartByCondition(
	ctx context.Context,
	conditions []interface{},
	_ ...string) (*cartmodel.Cart, error) {
	// define
	var cart cartmodel.Cart
	// find the cart with where status = 1
	// SHOULD update to fine where user_id == userid :)
	db := store.db.Table(cartmodel.Cart{}.TableName()).Where(conditions)

	// handle error
	if err := db.Where(conditions).First(&cart).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//return
	return &cart, nil
}
