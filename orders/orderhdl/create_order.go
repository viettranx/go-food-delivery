package orderhdl

import (
	"fooddlv/cart/carthdl"
	"fooddlv/common"
	"fooddlv/orders/ordermodel"
	"fooddlv/orders/orderrepo"
	"fooddlv/orders/orderstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order ordermodel.Order
		db := appCtx.GetDBConnection()

		if err := c.ShouldBind(&order); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		// TODO: get userId

		store := orderstorage.NewOrderSQLStore(db)
		repo := orderrepo.NewCreateOrderRepo(store, &carthdl.FakeUserIdStore{})

		result, err := repo.CreateOrder(c.Request.Context(), order)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
