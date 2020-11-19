package notemodel

import "fooddlv/common"

const EntityName = "Note"

type Note struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;"`
	Content         string `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}
