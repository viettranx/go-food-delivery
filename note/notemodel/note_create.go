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
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
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
