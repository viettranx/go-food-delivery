package detailrepo

import (
	"context"
	"fooddlv/cart/cartrepo"
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
)

type CancelOrderStorage interface {
	FindOrder(ctx context.Context, orderId map[string]interface{}, moreInfos ...string) (*detailsmodel.Order, error)
	Update(ctx context.Context, order *detailsmodel.Order, orderId int) error
}

type cancelOrderRepo struct {
	userIDStore cartrepo.GetUserIdStore
	store       CancelOrderStorage
}

func NewCancelOrderRepo(store CancelOrderStorage, userIDStore cartrepo.GetUserIdStore) *cancelOrderRepo {
	return &cancelOrderRepo{store: store, userIDStore: userIDStore}
}

func (repo *cancelOrderRepo) CancelOrder(ctx context.Context, orderId int) error {
	var updateData *detailsmodel.Order
	userId, err := repo.userIDStore.GetUserId(ctx)
	if err != nil {
		return common.ErrCannotGetEntity("user", err)
	}
	updateData, err = repo.store.FindOrder(ctx, map[string]interface{}{"order_id": orderId})

	if err != nil {
		return common.ErrCannotGetEntity(detailsmodel.EntityName, err)
	}

	updateData.Status = common.OrderCanceled

	if err := repo.store.Update(ctx, updateData, userId); err != nil {
		return err
	}
	return nil
}
