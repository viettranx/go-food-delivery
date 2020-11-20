package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/note/notemodel"
)

func (s *storeMysql) Find(
	ctx context.Context,
	cond map[string]interface{},
	moreInfos ...string,
) (*notemodel.Note, error) {
	var rs notemodel.Note

	db := s.db.Table(notemodel.Note{}.TableName()).Where(cond)

	if err := db.First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &rs, nil
}
