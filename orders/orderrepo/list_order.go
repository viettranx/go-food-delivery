package orderrepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/orders/ordermodel"
)

type ListOrderStorage interface {
	List(ctx context.Context, paging *common.Paging, filter *ordermodel.ListFilter) ([]ordermodel.Order, error)
}

type listOrderRepo struct {
	store ListOrderStorage
}

func NewListOrderRepo(store ListOrderStorage) *listOrderRepo {
	return &listOrderRepo{store: store}
}

func (repo *listOrderRepo) ListOrder(ctx context.Context, paging *common.Paging, filter *ordermodel.ListFilter) ([]ordermodel.Order, error) {
	orders, err := repo.store.List(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return orders, nil
}
