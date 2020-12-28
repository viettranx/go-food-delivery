package noterepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type GetNoteStorage interface {
	Find(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*notemodel.Note, error)
}

type getNoteRepo struct {
	store GetNoteStorage
}

func NewGetNoteRepo(store GetNoteStorage) *getNoteRepo {
	return &getNoteRepo{store: store}
}

func (repo *getNoteRepo) GetNote(ctx context.Context, id int) (*notemodel.Note, error) {
	note, err := repo.store.Find(ctx, map[string]interface{}{"id": id, "status": 1}, "User")

	if note == nil {
		return nil, common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	return note, nil
}
