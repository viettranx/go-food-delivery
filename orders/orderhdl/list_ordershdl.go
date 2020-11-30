package orderhdl

import (
	"fooddlv/common"
	"fooddlv/order_details/detailstorage"
	"fooddlv/orders/ordermodel"
	"fooddlv/orders/orderrepo"
	"fooddlv/orders/orderstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOrder(appCtx common.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var param ordermodel.ListParam

		if err := ctx.ShouldBind(&param); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		param.Fulfill()
		db := appCtx.GetDBConnection()
		store := orderstorage.NewOrderSQLStore(db)
		detailStore := detailstorage.NewOrderDetailSQLStorage(db)
		repo := orderrepo.NewListOrderRepo(store, detailStore)

		result, err := repo.ListOrder(ctx.Request.Context(), &param.Paging, param.ListFilter)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, param.Paging, param.ListFilter))
	}
}

// define param
// if err => err.Error()
// param Fulfill
// get connection
// new order store
// create repo list
// get result
// return response
