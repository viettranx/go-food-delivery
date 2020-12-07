package detailstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
)

func (s *orderDetailStorage) FindOrder(ctx context.Context, cond map[string]interface{}, moreInfos ...string) (*detailsmodel.Order, error) {
	db := s.db.Table(detailsmodel.Order{}.TableName()).Where(cond)

	for i := range moreInfos {
		db = db.Preload(moreInfos[i])
	}

	var rs detailsmodel.Order

	if err := db.First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &rs, nil
}

//(s *orderStorage) FindOrder(ctx context.Context, orderId map[string]interface{}, moreInfos ...string) (*detailsmodel.Order, error)
