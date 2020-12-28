package authhdl

import (
	"fooddlv/common"
	"fooddlv/module/auth/authmodel"
	"fooddlv/module/auth/authrepo"
	"fooddlv/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx common.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		var user authmodel.CreateUser

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := user.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := userstorage.NewUserMysql(db)
		repo := authrepo.NewAuthRepo(store)

		userId, err := repo.Register(c.Request.Context(), &user)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userId))
	}
}
