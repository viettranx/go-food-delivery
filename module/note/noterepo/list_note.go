package noterepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

type ListNoteStorage interface {
	List(ctx context.Context, paging *common.Paging, filter *notemodel.ListFilter) ([]notemodel.Note, error)
}

type listNoteRepo struct {
	store ListNoteStorage
}

func NewListNoteRepo(store ListNoteStorage) *listNoteRepo {
	return &listNoteRepo{store: store}
}

func (repo *listNoteRepo) ListNote(ctx context.Context, paging *common.Paging, filter *notemodel.ListFilter) ([]notemodel.Note, error) {
	notes, err := repo.store.List(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(notemodel.EntityName, err)
	}

	return notes, nil
}
