package imgmodel

type Image struct {
	ID     int    `json:"id" gorm:"column:id"`
	Url    string `json:"url" gorm:"column:url"`
	Width  int    `json:"width" gorm:"column:width"`
	Height int    `json:"height" gorm:"column:height"`
}

func (Image) TableName() string {
	return "images"
}
