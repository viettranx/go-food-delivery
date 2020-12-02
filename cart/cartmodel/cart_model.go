package cartmodel

import "fooddlv/common"

const EntityName string = "cart"

type CartItem struct {
	UserId   int `json:"user_id" gorm:"column:user_id"`
	FoodId   int
	Quantity int
}

type Cart struct {
	common.SQLModel `json:",inline"`
	UserId          []int      `json:"user_id" gorm:"column:user_id"`
	CartItems       []CartItem `json:"cart_items" gorm:"cart_items"`
}

func (Cart) TableName() string {
	return EntityName
}
