package noterepo

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
	"fooddlv/pubsub"
)

type GetImageStorage interface {
	GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error)
}

type CreateNoteStorage interface {
	Create(ctx context.Context, data *notemodel.NoteCreate) error
}

type createNoteRepo struct {
	imgStore GetImageStorage
	store    CreateNoteStorage
	pb       pubsub.Pubsub
}

func NewCreateNoteRepo(imgStore GetImageStorage, store CreateNoteStorage, pb pubsub.Pubsub) *createNoteRepo {
	return &createNoteRepo{imgStore: imgStore, store: store, pb: pb}
}

func (repo *createNoteRepo) CreateNote(ctx context.Context, data *notemodel.NoteCreate) error {
	imgs, err := repo.imgStore.GetImages(ctx, nil, data.ImageIds)

	if err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	if len(data.ImageIds) != len(imgs) {
		return common.ErrCannotCreateEntity(notemodel.EntityName, errors.New("images not enough"))
	}

	t := common.Images(imgs)
	data.Images = &t

	err = repo.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	repo.pb.Publish(ctx, common.ChanNoteCreated, pubsub.NewMessage(data))

	//go func() {
	//
	//
	//	// All images with status 0 is used, otherwise is unused
	//	if err := common.NewJob(a).Execute(ctx); err != nil {
	//		repo.imgStore.DeleteImages(ctx, data.ImageIds)
	//	}
	//
	//}()

	return nil
}
