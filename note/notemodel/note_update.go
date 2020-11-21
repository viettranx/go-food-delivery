package notemodel

type NoteUpdate struct {
	Title   *string `json:"title" gorm:"column:title;"`
	Content *string `json:"content" gorm:"column:content;"`
}

func (NoteUpdate) TableName() string { return Note{}.TableName() }
