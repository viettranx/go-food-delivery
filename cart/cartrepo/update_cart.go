package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
)

type UpdateCartStorage interface {
	UpdateCart(ctx context.Context, updateData *cartmodel.Cart) (int, error)
}

type updateCartRepo struct {
	store UpdateCartStorage
}

func NewUpdateCartRepo(store UpdateCartStorage) *updateCartRepo {
	return &updateCartRepo{store: store}
}

func (repo *updateCartRepo) UpdateCart(ctx context.Context, updateData *cartmodel.Cart) (int, error) {

	if _, err := repo.store.UpdateCart(ctx, updateData); err != nil {
		return 0, err
	}

	return 1, nil
}
