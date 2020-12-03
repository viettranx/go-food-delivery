package foodhdl

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
	"fooddlv/food/foodrepo"
	"fooddlv/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type fakeImageStore struct{}

func (fakeImageStore) GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error) {
	if len(ids) == 0 {
		return nil, errors.New("image ids can not be empty")
	}

	return []common.Image{
		{
			Id:     1,
			Url:    "https://",
			Width:  100,
			Height: 100,
		},
		{
			Id:     2,
			Url:    "https://",
			Width:  200,
			Height: 200,
		},
	}, nil
}

func (fakeImageStore) DeleteImages(ctx context.Context, ids []int) error {
	return nil
}

func CreateFood(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data foodmodel.Food
		data.Status = 1

		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := foodstorage.NewMysqlStore(db)

		repo := foodrepo.NewCreateFoodRepo(store, &fakeImageStore{})
		if err := repo.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
