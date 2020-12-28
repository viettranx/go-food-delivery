package noterepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type DeleteNoteStorage interface {
	Find(
		ctx context.Context,
		cond map[string]interface{},
		moreInfos ...string,
	) (*notemodel.Note, error)
	Delete(ctx context.Context, id int) error
}

type deleteNoteRepo struct {
	store DeleteNoteStorage
}

func NewDeleteNoteRepo(store DeleteNoteStorage) *deleteNoteRepo {
	return &deleteNoteRepo{store: store}
}

func (repo *deleteNoteRepo) DeleteNote(ctx context.Context, id int) (*notemodel.Note, error) {
	note, err := repo.store.Find(ctx, map[string]interface{}{"id": id, "status": 1})

	if note == nil {
		return nil, common.ErrCannotGetEntity(notemodel.EntityName, err)
	}

	if err := repo.store.Delete(ctx, id); err != nil {
		return nil, common.ErrCannotDeleteEntity(notemodel.EntityName, err)
	}

	return note, nil
}
