package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

func (s *storeMysql) Find(
	ctx context.Context,
	cond map[string]interface{},
	moreInfos ...string,
) (*notemodel.Note, error) {
	db := s.db.Table(notemodel.Note{}.TableName()).Where(cond)

	for i := range moreInfos {
		db = db.Preload(moreInfos[i])
	}

	var rs notemodel.Note

	if err := db.First(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &rs, nil
}
