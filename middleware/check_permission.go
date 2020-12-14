package middleware

import (
	"fooddlv/common"
	"github.com/gin-gonic/gin"
)

type PermissionStore interface {
	GetPermission() ([]interface{}, error)
}

func CheckPermission(sc common.AppContext, resourceName string, store PermissionStore) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
