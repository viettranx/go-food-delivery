package cartrepo

import (
	"context"
	"fooddlv/cart/cartmodel"
	"fooddlv/common"
)

type GetUserIdStore interface {
	GetUserId(ctx context.Context) (int, error)
}

type CreateCartStorage interface {
	Create(ctx context.Context, createCartsData []*cartmodel.Cart) error
}

type createCartRepo struct {
	userIDStore GetUserIdStore
	store       CreateCartStorage
}

func NewCreateCartRepo(store CreateCartStorage, userIDStore GetUserIdStore) *createCartRepo {
	return &createCartRepo{store: store, userIDStore: userIDStore}
}

func (repo *createCartRepo) CreateCart(ctx context.Context, createCartsData []*cartmodel.Cart) (result []int, err error) {

	// Get User id,
	userId, err := repo.userIDStore.GetUserId(ctx)
	result = make([]int, len(createCartsData))
	if err != nil {
		return nil, common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}
	for i := range createCartsData {
		createCartsData[i].UserID = userId
		result[i] = createCartsData[i].FoodID
	}
	// otherwise handle the add to the cart
	if err := repo.store.Create(ctx, createCartsData); err != nil {
		return nil, common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}
	return result, nil
}
