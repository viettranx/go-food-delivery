package carthdl

import (
	"fooddlv/cart/cartmodel"
	"fooddlv/cart/cartrepo"
	"fooddlv/cart/cartstorage"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddToCart(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var carts []cartmodel.Cart
		db := appCtx.GetDBConnection()

		if err := c.ShouldBind(&carts); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// TODO: get userId

		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewCreateCartRepo(store)

		result, err := repo.AddToCart(c.Request.Context(), &carts)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result > 0))
	}
}
