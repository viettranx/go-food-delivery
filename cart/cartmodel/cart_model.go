package cartmodel

import "time"

const EntityName string = "carts"

type Cart struct {
	UserId    int `json:"user_id" gorm:"column:user_id"`
	FoodId    int `json:"food_id" gorm:"column:food_id"`
	Quantity  int
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type CartCreate struct {
	Carts []Cart
}

func (Cart) TableName() string {
	return EntityName
}

func (data *CartCreate) GetCarts() *[]Cart {
	return &data.Carts
}
