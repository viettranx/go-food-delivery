package foodhdl

import (
	"fooddlv/common"
	"fooddlv/food/foodrepo"
	"fooddlv/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteFood(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("food-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()

		store := foodstorage.NewMysqlStore(db)
		repo := foodrepo.NewDeleteFoodRepo(store)

		err = repo.Delete(id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
