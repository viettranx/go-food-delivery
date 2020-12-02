package notemodel

import (
	"errors"
	"fooddlv/common"
	"strings"
)

var (
	ErrTitleCannotEmpty = common.NewCustomError(errors.New("ErrNoteTitleEmpty"), "note title cannot not be empty", "ErrNoteTitleEmpty")
)

type NoteCreate struct {
	Id      int           `json:"id" gorm:"column:id;"`
	Title   string        `json:"title" gorm:"column:title;"`
	Content string        `json:"content" gorm:"column:content;"`
	Image   *common.Image `json:"image" gorm:"column:image;"`

	// Case upload image ids
	ImageIds []int          `json:"image_ids" gorm:"-"`
	Images   *common.Images `json:"images" gorm:"column:images;"`
}

func (NoteCreate) TableName() string { return Note{}.TableName() }

func (n *NoteCreate) Validate() error {
	n.Title = strings.TrimSpace(n.Title)
	n.Content = strings.TrimSpace(n.Content)

	if len(n.Title) == 0 {
		return ErrTitleCannotEmpty
	}

	return nil
}
