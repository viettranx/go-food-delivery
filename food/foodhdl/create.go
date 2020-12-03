package foodhdl

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/food/foodmodel"
	"fooddlv/food/foodrepo"
	"fooddlv/food/foodstorage"
	"fooddlv/upload/imgstorage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type fakeImageStore struct {
	db *gorm.DB
}

func NewFakeImageStore(db *gorm.DB) *fakeImageStore {
	return &fakeImageStore{
		db: db,
	}
}

func (imageStore fakeImageStore) GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error) {
	if len(ids) == 0 {
		return nil, errors.New("image ids can not be empty")
	}

	imageStorage := imgstorage.NewImgSqlStorage(imageStore.db)

	return imageStorage.GetImages(ctx, cond, ids)
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

		repo := foodrepo.NewCreateFoodRepo(store, NewFakeImageStore(db))
		if err := repo.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
