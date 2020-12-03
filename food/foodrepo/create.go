package foodrepo

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

type GetImageStorage interface {
	GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error)
	DeleteImages(ctx context.Context, ids []int) error
}

type CreateFoodStorage interface {
	Create(ctx context.Context, data *foodmodel.Food) error
}

type createFoodRepo struct {
	createFoodStorage CreateFoodStorage
	imgStore          GetImageStorage
}

func NewCreateFoodRepo(createFoodStorage CreateFoodStorage, imgStore GetImageStorage) *createFoodRepo {
	return &createFoodRepo{
		createFoodStorage: createFoodStorage,
		imgStore:          imgStore,
	}
}

func (repo *createFoodRepo) CreateFood(ctx context.Context, data *foodmodel.Food) error {
	imgs, err := repo.imgStore.GetImages(ctx, nil, data.ImageIds)

	if err != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	if len(data.ImageIds) != len(imgs) {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, errors.New("images not enough"))
	}

	t := common.Images(imgs)
	data.Images = &t

	error := repo.createFoodStorage.Create(ctx, data)

	if error != nil {
		return common.ErrCannotCreateEntity(foodmodel.EntityName, err)
	}

	go func() {
		// All images with status 0 is used, otherwise is unused
		repo.imgStore.DeleteImages(ctx, data.ImageIds) // side effect
	}()

	return nil
}
