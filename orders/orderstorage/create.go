package orderstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/orders/ordermodel"
)

func (s *orderStorage) CreateOrder(ctx context.Context, orderCreate ordermodel.Order) (orderId int, err error) {
	if err := s.db.Table(ordermodel.Order{}.TableName()).Create(&orderCreate).Error; err != nil {
		return 0, common.ErrDB(err)
	}
	return orderCreate.ID, nil
}
