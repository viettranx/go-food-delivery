package detailstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
)

func (s *orderDetailStorage) ListOrderDetail(
	ctx context.Context,
	paging *common.Paging,
	filter *detailsmodel.ListFilter,
) ([]detailsmodel.Order, error) {
	var rs []detailsmodel.Order

	db := s.db.Table(detailsmodel.Order{}.TableName()).Where("status = 1")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return rs, nil
}
