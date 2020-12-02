package carthdl

import (
	"fooddlv/cart/cartmodel"
	"fooddlv/cart/cartrepo"
	"fooddlv/cart/cartstorage"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateCart(appCtx common.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data cartmodel.Cart
		db := appCtx.GetDBConnection()
		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewUpdateCartRepo(store)
		//if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
		//	panic(err)
		//}
		if err := ctx.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}
		_, err := repo.UpdateCart(ctx.Request.Context(), &data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
