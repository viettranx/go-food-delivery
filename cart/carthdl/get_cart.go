package carthdl

import (
	"fooddlv/cart/cartmodel"
	"fooddlv/cart/cartrepo"
	"fooddlv/cart/cartstorage"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowCart(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var p cartmodel.Cart

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewCartDetailRepo(store)

		result, err := repo.ShowCartDetailFromUser(c.Request.Context(), p.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrCannotGetEntity(cartmodel.EntityName, err))
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}
