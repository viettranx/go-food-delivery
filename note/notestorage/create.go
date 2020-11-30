package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/note/notemodel"
)

func (s *storeMysql) Create(ctx context.Context, data *notemodel.NoteCreate) error {
	if err := s.db.Table(notemodel.Note{}.
		TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
