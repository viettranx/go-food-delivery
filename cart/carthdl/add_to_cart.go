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
		db := appCtx.GetDBConnection()
		userID := cartmodel.CartCreation{
			FoodId:   12,
			Quantity: 1,
		}
		if err := c.ShouldBind(&userID); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewCreateCartRepo(store)

		result, err := repo.AddToCart(c.Request.Context(), &userID)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result > 0))
	}
}
