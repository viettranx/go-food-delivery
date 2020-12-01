package carthdl

import (
	"fooddlv/cart/cartrepo"
	"fooddlv/cart/cartstorage"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCart(appCtx common.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		userId := `ds1`
		db := appCtx.GetDBConnection()
		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewUpdateCartRepo(store)

		_, err := repo.UpdateCartFromUser(ctx.Request.Context(), userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
