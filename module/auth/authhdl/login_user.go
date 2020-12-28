package authhdl

import (
	"fooddlv/common"
	"fooddlv/module/auth/authmodel"
	"fooddlv/module/auth/authrepo"
	"fooddlv/module/user/userstorage"
	"fooddlv/token"
	"fooddlv/token/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx common.AppContext, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		var loginUserData authmodel.LoginUser

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		tokenProvider := jwt.NewTokenProvider(token.WithSecretKey([]byte(secretKey)))

		store := userstorage.NewUserMysql(db)
		repo := authrepo.NewLoginUserRepo(store, tokenProvider)
		account, err := repo.LoginUser(c.Request.Context(), &loginUserData)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"access_token": account.AccessToken,
		})
	}
}
