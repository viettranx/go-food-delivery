package noterepo

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/note/notemodel"
)

type GetImageStorage interface {
	GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error)
	DeleteImages(ctx context.Context, ids []int) error
}

type CreateNoteStorage interface {
	Create(ctx context.Context, data *notemodel.NoteCreate) error
}

type createNoteRepo struct {
	imgStore GetImageStorage
	store    CreateNoteStorage
}

func NewCreateNoteRepo(imgStore GetImageStorage, store CreateNoteStorage) *createNoteRepo {
	return &createNoteRepo{imgStore: imgStore, store: store}
}

func (repo *createNoteRepo) CreateNote(ctx context.Context, data *notemodel.NoteCreate) error {
	imgs, err := repo.imgStore.GetImages(ctx, nil, data.ImageIds)

	if err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	if len(data.ImageIds) != len(imgs) {
		return common.ErrCannotCreateEntity(notemodel.EntityName, errors.New("images not enough"))
	}

	imgsTemp := make([]*common.Image, len(imgs))
	for i := range imgs {
		imgsTemp[i] = &imgs[i]
	}
	t := common.Images(imgsTemp)
	data.Images = &t

	err = repo.store.Create(ctx, data)

	if err != nil {
		return common.ErrCannotCreateEntity(notemodel.EntityName, err)
	}

	go func() {
		// All images with status 0 is used, otherwise is unused
		repo.imgStore.DeleteImages(ctx, data.ImageIds) // side effect
	}()

	return nil
}
