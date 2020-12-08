package orderrepo

import (
	"context"
	"fooddlv/cart/cartrepo"
	"fooddlv/common"
	"fooddlv/orders/ordermodel"
)

type CreateOrderStorage interface {
	CreateOrder(ctx context.Context, orderCreate ordermodel.Order) (orderId int, err error)
}

type createOrderRepo struct {
	store       CreateOrderStorage
	userIDStore cartrepo.GetUserIdStore
}

func NewCreateOrderRepo(store CreateOrderStorage, userIDStore cartrepo.GetUserIdStore) *createOrderRepo {
	return &createOrderRepo{store: store, userIDStore: userIDStore}
}

func (repo *createOrderRepo) CreateOrder(ctx context.Context, orderCreate ordermodel.Order) (orderId int, err error) {
	userId, err := repo.userIDStore.GetUserId(ctx)

	if err != nil {
		return 0, common.ErrCannotGetEntity("user", err)
	}
	orderCreate.UserID = userId

	if orderId, err = repo.store.CreateOrder(ctx, orderCreate); err != nil {
		return 0, common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}
	return orderId, nil
}
