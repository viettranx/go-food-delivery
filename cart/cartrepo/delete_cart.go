package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

type DeleteCartStorage interface {
	Delete(ctx context.Context, cartId int, userId int) error
}

type deleteCartRepo struct {
	userIDStore GetUserIdStore
	store       DeleteCartStorage
}

func NewDeleteCartRepo(store DeleteCartStorage, userIDStore GetUserIdStore) *deleteCartRepo {
	return &deleteCartRepo{store: store, userIDStore: userIDStore}
}

func (repo *deleteCartRepo) DeleteCart(ctx context.Context, cartId int) error {

	// Get User id,
	userId, err := repo.userIDStore.GetUserId(ctx)

	if err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	// otherwise handle the add to the cart
	if err := repo.store.Delete(ctx, cartId, userId); err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}
	return nil
}
