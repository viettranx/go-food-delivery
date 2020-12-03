package foodrepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
)

type ListFoodStorage interface {
	List(ctx context.Context, paging *common.Paging, filter *foodmodel.ListFilter) ([]foodmodel.Food, error)
}

type listFoodRepo struct {
	store ListFoodStorage
}

func NewListFoodRepo(store ListFoodStorage) *listFoodRepo {
	return &listFoodRepo{store: store}
}

func (repo *listFoodRepo) ListFood(ctx context.Context, paging *common.Paging, filter *foodmodel.ListFilter) ([]foodmodel.Food, error) {
	foods, err := repo.store.List(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return foods, nil
}

