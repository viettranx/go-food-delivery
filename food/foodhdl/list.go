package foodhdl

import (
	"fooddlv/common"
	"fooddlv/food/foodmodel"
	"fooddlv/food/foodrepo"
	"fooddlv/food/foodstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFood(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var p foodmodel.ListParam

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		p.Fulfill()

		db := appCtx.GetDBConnection()

		store := foodstorage.NewMysqlStore(db)
		repo := foodrepo.NewListFoodRepo(store)

		result, err := repo.ListFood(c.Request.Context(), &p.Paging, p.ListFilter)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, p.Paging, p.ListFilter))
	}
}

