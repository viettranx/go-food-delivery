package imgrepo

import (
	"context"
	"fooddlv/common"
)

type CreateImgStorage interface {
	Create(ctx context.Context, data []common.Image) error
}

type createImgRepo struct {
	store CreateImgStorage
}

func NewCreateImgRepo(store CreateImgStorage) *createImgRepo {
	return &createImgRepo{
		store: store,
	}
}

func (repo *createImgRepo) Create(ctx context.Context, data []common.Image) error {
	return repo.store.Create(ctx, data)
}
