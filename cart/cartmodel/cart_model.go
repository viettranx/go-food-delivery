package cartmodel

import "fooddlv/common"

const EntityName string = "cart"

type CartItem struct {
	foods    []string
	quantity int
}

type Cart struct {
	//Foods *foodmodel.Foods `json:",inline,foods"`
	common.SQLModel `json:",inline"`
	UserId          string
	items           []CartItem
}

type CartCreation struct {
	common.SQLModel `json:",inline"`
	FoodId          int
	Quantity        int
}

type CartUpdate struct {
	common.SQLModel `json:",inline"`
	FoodId          int
	Quantity        int
}

func (Cart) TableName() string {
	return EntityName
}
