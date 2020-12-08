package detailsmodel

import "fooddlv/common"

const EntityName = "OrderDetail"

type Order struct {
	common.SQLModel `json:",inline"`
	OrderID         int                    `json:"order_id" gorm:"column:order_id"`
	Food            string                 `json:"food" gorm:"column:food_origin"`
	Price           float32                `json:"total" gorm:"column:price"`
	Quantity        int                    `json:"quantity" gorm:"column:quantity"`
	Discount        float32                `json:"discount" gorm:"column:discount"`
	Status          common.OrderStatusEnum `json:"status" gorm:"column:status"`
	CratedAt        float32                `json:"create_at" gorm:"column:create_at"`
	UpdatedAt       float32                `json:"update_at" gorm:"column:update_at"`
}

func (Order) TableName() string {
	return "order_details"
}
