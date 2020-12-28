package repo

import (
	"context"
	"fooddlv/common"
	"fooddlv/food_likes/model"
)

type DeleteFoodLikeStorage interface {
	Delete(ctx context.Context, foodLike *model.FoodLikes) error
	FindByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		relations ...string,
	) (*model.FoodLikes, error)
}

type deleteFoodLikeRepo struct {
	deleteFoodLikeStorage DeleteFoodLikeStorage
}

func NewDeleteFoodRepo(deleteFoodLikeStorage DeleteFoodLikeStorage) *deleteFoodLikeRepo {
	return &deleteFoodLikeRepo{
		deleteFoodLikeStorage: deleteFoodLikeStorage,
	}
}

func (repo *deleteFoodLikeRepo) DeleteFoodLike(ctx context.Context, userId int, foodId int) error {

	foodLike, err := repo.deleteFoodLikeStorage.FindByCondition(ctx, map[string]interface{}{"user_id": userId, "food_id": foodId}, "FoodLikes")

	if err != nil {
		return err
	}
	if err := repo.deleteFoodLikeStorage.Delete(ctx, foodLike); err != nil {
		return common.ErrCannotDeleteEntity(model.FoodLikes{}.TableName(), err)
	}
	return nil
}
