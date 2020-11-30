package foodhdl

import (
	"fooddlv/common"
	"fooddlv/food/foodrepo"
	"fooddlv/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindFood(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("food-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()

		store := foodstorage.NewMysqlStore(db)
		repo := foodrepo.NewFindFoodStorage(store)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		food, err := repo.Find(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(food, nil, nil))
	}
}

