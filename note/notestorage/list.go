package notestorage

import "fooddlv/note/notemodel"

func (s *storeMysql) List() ([]notemodel.Note, error) {
	var rs []notemodel.Note

	if err := s.db.Limit(10).Find(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}
