package detailshdl

import (
	"fooddlv/cart/carthdl"
	"fooddlv/common"
	"fooddlv/order_details/detailrepo"
	"fooddlv/order_details/detailstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PUT v1/notes/:note-id

func CancelOrder(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("order-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		store := detailstorage.NewOrderDetailSQLStorage(db)
		repo := detailrepo.NewCancelOrderRepo(store, carthdl.FakeUserIdStore{})

		err = repo.CancelOrder(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
