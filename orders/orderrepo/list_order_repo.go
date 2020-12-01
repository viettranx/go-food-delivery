package orderrepo

import (
	"context"
	"fmt"
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
	"fooddlv/orders/ordermodel"
)

type ListOrderStorage interface {
	ListOrder(ctx context.Context, paging *common.Paging, filter *ordermodel.ListFilter) ([]ordermodel.Order, error)
}

type ListOrderDetailStorage interface {
	ListOrderDetail(ctx context.Context, paging *common.Paging, filter *detailsmodel.ListFilter) ([]detailsmodel.Order, error)
}

type listOrderRepo struct {
	store            ListOrderStorage
	orderDetailStore ListOrderDetailStorage
}

func NewListOrderRepo(store ListOrderStorage, orderDetailStore ListOrderDetailStorage) *listOrderRepo {
	return &listOrderRepo{store: store, orderDetailStore: orderDetailStore}
}

func (repo *listOrderRepo) ListOrder(
	ctx context.Context,
	paging *common.Paging,
	filter *ordermodel.ListFilter,
) ([]ordermodel.Order, error) {
	orders, err := repo.store.ListOrder(ctx, paging, filter)
	ordersDetail, err := repo.orderDetailStore.ListOrderDetail(ctx, paging, nil)

	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}
	fmt.Print(orders, ordersDetail)
	return orders, nil
}
