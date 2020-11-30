package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
)

type CartDetailStorage interface {
	ShowCartFromUser(ctx context.Context, userId string) (*cartmodel.Cart, error)
}

type cartDetailRepo struct {
	store CartDetailStorage
}

func NewCartDetailRepo(store CartDetailStorage) *cartDetailRepo {
	return &cartDetailRepo{store: store}
}

func (repo *cartDetailRepo) ShowCartDetailFromUser(ctx context.Context, userId string) (*cartmodel.Cart, error) {
	cartDetail, err := repo.store.ShowCartFromUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	return cartDetail, nil
}
