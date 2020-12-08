package ordermodel

import (
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
)

const EntityName = "Order"

type Order struct {
	common.SQLModel `json:",inline"`
	UserID          int                  `json:"user_id" gorm:"column:user_id;"`
	TotalPrice      int                  `json:"total" gorm:"column:total_price;"`
	Status          int                  `json:"status" gorm:"column: status;"`
	CreateAt        float32              `json:"create_at" gorm:"column:create_at"`
	OrderDetails    []detailsmodel.Order `json:"order_details"`
}

func (Order) TableName() string {
	return "orders"
}
