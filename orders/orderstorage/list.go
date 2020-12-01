package orderstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/orders/ordermodel"
)

func (s *orderStorage) ListOrder(
	ctx context.Context,
	paging *common.Paging,
	filter *ordermodel.ListFilter,
) ([]ordermodel.Order, error) {
	var rs []ordermodel.Order
	db := s.db.Table(ordermodel.Order{}.TableName()).Where("status = 1")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return rs, nil
}
