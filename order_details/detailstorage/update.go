package detailstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
)

func (s *orderDetailStorage) Update(
	ctx context.Context,
	orderUpdateData *detailsmodel.Order,
	id int,
) error {
	db := s.db.Begin()

	if err := db.Table(detailsmodel.Order{}.TableName()).
		Where("id = ?", id).
		Updates(orderUpdateData).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
