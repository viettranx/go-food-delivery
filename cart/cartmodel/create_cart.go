package cartmodel

import (
	"fooddlv/common"
)

type CreateCart struct {
	common.SQLModel `json:",inline"`
}

func (CreateCart) TableName() string {
	return Cart{}.TableName()
}
