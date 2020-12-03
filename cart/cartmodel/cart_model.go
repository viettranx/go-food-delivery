package cartmodel

import (
	"fooddlv/common"
	"time"
)

const EntityName string = "carts"

type Cart struct {
	UserID    int        `json:"user_id" gorm:"column:user_id;"`
	FoodID    int        `json:"food_id" gorm:"column:food_id;"`
	Food      Food       `gorm:"foreignKey:FoodID"`
	Quantity  int        `json:"quantity" gorm:"column:quantity;"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Cart) TableName() string {
	return EntityName
}

//func (data *CartData) GetCarts() *[]Cart {
//	return &data.Carts
//}

type Food struct {
	common.SQLModel `json:",inline"`
	Name            string       `json:"name" gorm:"column:name;"`
	Description     string       `json:"description" gorm:"column:description;"`
	Price           int          `json:"price" gorm:"column:price;"`
	Images          *common.JSON `json:"images,omitempty" gorm:"-"`
	Status          int          `json:"status" gorm:"column:status;default:1"`
}

func (Food) TableName() string {
	return "foods"
}

type SimpleUser struct {
	common.SQLModel `json:",inline"`
	LastName        string          `json:"last_name" gorm:"column:last_name;"`
	FirstName       string          `json:"first_name" gorm:"column:first_name;"`
	Roles           common.RoleEnum `json:"-" gorm:"column:roles;type:ENUM('user', 'admin')"`
	Avatar          *common.JSON    `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}
