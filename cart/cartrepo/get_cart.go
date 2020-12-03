package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
)

type CartDetailStorage interface {
	List(ctx context.Context, userId int) ([]cartmodel.Cart, error)
}

type cartDetailRepo struct {
	store CartDetailStorage
}

func NewCartDetailRepo(store CartDetailStorage) *cartDetailRepo {
	return &cartDetailRepo{store: store}
}

func (repo *cartDetailRepo) GetCart(ctx context.Context, userId int) ([]cartmodel.Cart, error) {
	cartDetail, err := repo.store.List(ctx, userId)

	if err != nil {
		return nil, err
	}

	return cartDetail, nil
}
