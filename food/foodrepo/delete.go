package foodrepo

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

type DeleteFoodStorage interface {
	Delete(id int) error
	Find(id int, moreInfos ...string) (*foodmodel.Food, error)
}

type deleteFoodRepo struct {
	store DeleteFoodStorage
}

func NewDeleteFoodRepo(storage DeleteFoodStorage) *deleteFoodRepo {
	return &deleteFoodRepo{store: storage}
}

func (repo *deleteFoodRepo) Delete(id int) error {
	_, err := repo.store.Find(id)

	if err != nil {
		return common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	err = repo.store.Delete(id)

	if err != nil {
		return common.ErrCannotDeleteEntity(foodmodel.EntityName, err)
	}

	return nil
}
