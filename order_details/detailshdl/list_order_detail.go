package detailshdl

import (
	"fooddlv/common"
	"fooddlv/order_details/detailsmodel"
	"github.com/gin-gonic/gin"
)

func ListOrderDetail(appCtx common.AppContext) func(ctx *gin.Context) {
	// return func(c *gin.Context) {}
	return func(c *gin.Context) {
		// define param
		var param detailsmodel.ListParam

		// if err => err.Error()
		if err := c.ShouldBind(&param); err != nil && err.Error() != "EOF" {
			panic(err)
		}
		// param Fulfill
		param.Fulfill()

		// get connection
		// new order store
		// create repo list
		// get result
		// return response

	}
}
