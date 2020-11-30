package foodmodel

import "fooddlv/common"

const EntityName = "Food"

type Food struct {
	common.SQLModel 		   		`json:",inline"`
	Name           		string 		`json:"name" gorm:"column:name;"`
	Description         string 		`json:"description" gorm:"column:description;"`
	Price           	float32 	`json:"price" gorm:"column:price;"`
	Images          *common.JSON    `json:"images,omitempty" gorm:"column:images;type:json"`
}

func (Food) TableName() string {
	return "foods"
}
