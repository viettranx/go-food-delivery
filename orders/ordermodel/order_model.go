package ordermodel

import "fooddlv/common"

const EntityName = "Order"

type Order struct {
	common.SQLModel `json:",inline"`
	TotalPrice      int     `json:"total" gorm:"column:total_price;"`
	Status          int     `json:"status" gorm:"column: status;"`
	CreateAt        float32 `json:"create_at" gorm:"column:create_at"`
}

func (Order) TableName() string {
	return "orders"
}
