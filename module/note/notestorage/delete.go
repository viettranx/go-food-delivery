package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

func (s *storeMysql) Delete(
	ctx context.Context,
	id int,
) error {
	if err := s.db.Table(notemodel.Note{}.
		TableName()).
		Where("id = ?", id).
		UpdateColumns(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
