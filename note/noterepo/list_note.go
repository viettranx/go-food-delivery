package noterepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/note/notemodel"
	"go.opencensus.io/trace"
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
	ctx2, span1 := trace.StartSpan(ctx, "ListNote.repo")
	span1.AddAttributes(
		trace.BoolAttribute("a", true),
		trace.StringAttribute("filter", "user_id = 10"),
	)

	defer span1.End()

	//ctx2, span2 := trace.StartSpan(context.Background(), "test")
	//defer span2.End()

	//time.Sleep(time.Millisecond * 20)
	notes, err := repo.store.List(ctx2, paging, filter)

	span1.End()
	if err != nil {
		return nil, common.ErrCannotListEntity(notemodel.EntityName, err)
	}

	return notes, nil
}
