package handl

import (
	"fooddlv/common"
	"fooddlv/food_likes/model"
	"fooddlv/food_likes/repo"
	"fooddlv/food_likes/storage"
	"fooddlv/user/userrepo"
	"fooddlv/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateFoodLikes(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.FoodLikes

		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		foodStorage := storage.NewMysqlStore(db)
		userStorage := userstorage.NewUserMysql(db)

		userRepo := userrepo.NewFindUserStorage(userStorage)

		user, exists := c.Get(common.KeyCurrentUser)

		if !exists {
			c.JSON(http.StatusUnauthorized, exists)
			return
		}
		userId := user.(*common.SimpleUser).GetUserId()
		_, err := userRepo.FindUserByCondition(c.Request.Context(), map[string]interface{}{"id": userId})

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		data.UserId = int32(userId)
		foodId := c.Param("food-id")
		i, err := strconv.ParseInt(foodId, 10, 32)
		if err != nil {
			panic(err)
		}
		data.FoodId = int32(i)
		err = repo.NewCreateFoodLikesRepo(foodStorage).CreateFoodLikes(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
