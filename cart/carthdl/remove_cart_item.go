package carthdl

import (
	"context"
	"fooddlv/cart/cartrepo"
	"fooddlv/cart/cartstorage"
	"fooddlv/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (FakeUserIdStore) GetUserId(ctx context.Context) (int, error) {
	return 1, nil
}
func RemoveCartItem(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("food-id"))
		db := appCtx.GetDBConnection()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		// TODO: get userId

		store := cartstorage.NewCartMysql(db)
		repo := cartrepo.NewDeleteCartRepo(store, &FakeUserIdStore{})

		err = repo.DeleteCart(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
