package cartmodel

import (
	"fooddlv/common"
)

type CreateCart struct {
	common.SQLModel `json:",inline"`

	CartItems []CartItem `json:"cart_items" gorm:"-"`
}

func (CreateCart) TableName() string {
	return Cart{}.TableName()
}
