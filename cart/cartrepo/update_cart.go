package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

type UpdateCartStorage interface {
	UpdateCart(ctx context.Context, updateData *cartmodel.Cart, userId int) error
}

type updateCartRepo struct {
	userIdStore GetUserIdStore
	store       UpdateCartStorage
}

func NewUpdateCartRepo(store UpdateCartStorage, userIdStore GetUserIdStore) *updateCartRepo {
	return &updateCartRepo{store: store, userIdStore: userIdStore}
}

func (repo *updateCartRepo) UpdateCart(ctx context.Context, updateData *cartmodel.Cart) error {
	userId, err := repo.userIdStore.GetUserId(ctx)
	if err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}
	if err := repo.store.UpdateCart(ctx, updateData, userId); err != nil {
		return err
	}

	return nil
}
