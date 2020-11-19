package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/note/notemodel"
)

func (s *storeMysql) List(ctx context.Context, paging *common.Paging) ([]notemodel.Note, error) {
	var rs []notemodel.Note

	db := s.db.Table(notemodel.Note{}.TableName()).Where("status = 1")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.
		Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).
		Find(&rs).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return rs, nil
}
