package main

import (
	"fooddlv/appctx"
	"fooddlv/auth/authhdl"
	"fooddlv/middleware"
	"fooddlv/note/notehdl"
	"fooddlv/order_details/detailshdl"
	"fooddlv/orders/orderhdl"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

// Tier layer

// Repo (business logic) -----> Storage

func main() {
	dbConStr := os.Getenv("DBConnStr")
	secretKey := os.Getenv("SECRET_KEY")
	db, err := gorm.Open(mysql.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db.Debug())

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		var a []int
		log.Println(a[1])

		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	notes.GET("", notehdl.ListNote(appCtx))
	notes.DELETE("/:note-id", notehdl.DeleteNote(appCtx))

	notes.GET("/:note-id", func(c *gin.Context) {
		noteId := c.Param("note-id")
		c.String(http.StatusOK, "Hello %s", noteId)
	})

	auth := v1.Group("/auth")
	auth.POST("/register", authhdl.Register(appCtx))
	auth.POST("/login", authhdl.Login(appCtx, secretKey))

	//v1.GET("my-profile", ParseToken, GetProfile)
	users := v1.Group("users")
	users.GET("/:user-id")

	orders := users.Group("/:user-id/orders")
	orders.GET("", orderhdl.ListOrder(appCtx))
	orders.GET("/:order-id", detailshdl.ListOrderDetail(appCtx))

	r.Run()
}

type Requester interface {
	UserId() int
	Role() string
	FirstName() string
	LastName() string
}
