package notemodel

type NoteCreate struct {
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (NoteCreate) TableName() string { return Note{}.TableName() }
