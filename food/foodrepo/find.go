package foodrepo

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

type FindFoodStorage interface {
	Find(id int, moreInfos ...string) (*foodmodel.Food, error)
}

type findFoodRepo struct {
	store FindFoodStorage
}

func NewFindFoodStorage(storage DeleteFoodStorage) *findFoodRepo {
	return &findFoodRepo{store: storage}
}

func (repo *findFoodRepo) Find(id int) (*foodmodel.Food, error) {
	food, err := repo.store.Find(id)

	if err != nil {
		return nil, common.ErrCannotGetEntity(foodmodel.EntityName, err)
	}

	return food, nil
}
