package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
)

type UpdateCartStorage interface {
	UpdateCart(ctx context.Context, updateData cartmodel.CartUpdate) (*cartmodel.Cart, error)
}

type updateCartRepo struct {
	store UpdateCartStorage
}

func NewUpdateCartRepo(store UpdateCartStorage) *updateCartRepo {
	return &updateCartRepo{store: store}
}

func (repo *updateCartRepo) UpdateCartFromUser(ctx context.Context, updateData cartmodel.CartUpdate) (*cartmodel.Cart, error) {
	cartDetail, err := repo.store.UpdateCart(ctx, updateData)

	if err != nil {
		return nil, err
	}

	return cartDetail, nil
}
