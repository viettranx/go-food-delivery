package main

import (
	"fmt"
	"fooddlv/appctx"
	"fooddlv/consumers"
	"fooddlv/food_likes/handl"
	"fooddlv/middleware"
	"fooddlv/module/auth/authhdl"
	"fooddlv/module/note/notehdl/ginnote"
	"fooddlv/module/upload/imghdl"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
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
	// setup all consumers
	//consumers.Setup(appCtx)
	consumers.NewEngine(appCtx).Start()

	//go func() {
	//	time.Sleep(time.Second * 3)
	//	appCtx.GetPubsub().Publish(context.Background(), common.ChanNoteCreated, pubsub.NewMessage(10))
	//}()

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	notes.GET("", ginnote.ListNote(appCtx))
	notes.POST("", ginnote.CreateNote(appCtx))
	notes.DELETE("/:note-id", ginnote.DeleteNote(appCtx))

	notes.GET("/:note-id", func(c *gin.Context) {
		noteId := c.Param("note-id")
		c.String(http.StatusOK, "Hello %s", noteId)
	})

	auth := v1.Group("/auth")
	auth.POST("/register", authhdl.Register(appCtx))
	auth.POST("/login", authhdl.Login(appCtx, secretKey))

	foodLikes := v1.Group("/food", middleware.RequiredAuth(appCtx, secretKey))
	foodLikes.POST("/:food-id/like", handl.CreateFoodLikes(appCtx))
	foodLikes.DELETE("/:food-id/dis-like", handl.DeleteFoodLike(appCtx))

	v1.Static("/file", "./public")
	upload := v1.Group("/upload")
	upload.POST("", imghdl.UploadImg(appCtx))

	//v1.GET("my-profile", ParseToken, GetProfile)
	//users := v1.Group("users", ParseToken)
	//users.GET("/:user-id")

	//job := common.NewJob(func(ctx context.Context) error {
	//	fmt.Println("Hahaha")
	//	return errors.New("something went wrong")
	//})
	//
	//log.Println(job.State())
	//
	//timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*10)
	//go job.Execute(timeoutCtx)
	//cancelFn()

	//log.Println(job.State(), job.GetError())
	//checkClosure()
	startSocketIOServer(r)
	r.Run()
}

func startSocketIOServer(engine *gin.Engine) {
	server, _ := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{websocket.Default},
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())

		//go func() {
		//	i := 0
		//	for {
		//		i++
		//		s.Emit("test", i)
		//		time.Sleep(time.Second)
		//	}
		//}()
		return nil
	})

	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
		// Validate token
		// If false: s.Close(), and return

		// If true
		// => UserId
		// Fetch db find user by Id
		// Here: s belongs to who? (user_id)
		// We need a map[user_id][]socketio.Conn
		log.Println(s.ID(), token)
	})

	type A struct {
		Age int `json:"age"`
	}

	server.OnEvent("/", "notice", func(s socketio.Conn, msg A) {
		fmt.Println("notice:", msg.Age)
		s.Emit("reply", msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
		// Remove socket from socket engine (from app context)
	})

	go server.Serve()

	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))

	engine.StaticFile("/demo/", "./demo.html")
}

type Requester interface {
	UserId() int
	Role() string
	FirstName() string
	LastName() string
}

//
//func AuthUser(c *gin.Context) {
//	c.Set("user", &u{int: 10})
//	c.Next()
//}
//
//func GetProfile(c *gin.Context) {
//	u, ok := c.Get("user")
//	userId := u.(Requester)
//}

// Flow: register
// 1. Parse input form
// 2. Gen salt (random string)
// 3. Gen password hashMD5(password + salt)
// 4. Return: true | token JWT

// Flow: Login
// 1. Parse login form
// 2. Find user with email
// 3. Gen password hashMD5(password (input) + salt (db))
// 4. Compare password (input) == password MD5
// 4. Return: token JWT

// Flow: Get user info (GET /v1/users/:id, /v1/my-profile) | Auth (Header: {token}
// 1. [Middleware] Get token from header, JWT parser -> user_id
// 1.1 Get user by user id
// 2. Handler get user, repo
// 3. Return user info

// API has image upload:
// Set avatar user, Create Food, Create Restaurant,
// API Upload Image

// API List/Get Food:
// We need a full object food_likes, within restaurant object (simple form):
// Ex: {"id": 1, "title": "abc", "restaurant": {...}}
// Done.

// Create Food (security enhancement) flow:
// 1. User upload images to upload API
// 1.1 Backend store image, insert to images db
// 1.2 Backend return array of image ids to the client
// 2. User create food_likes with food_likes json include image ids: {"title": "...", "img_ids": [1,2,3]}
// 2.1 Backend fetch image objects by ids
// 2.2 Insert new food_likes with request body data and image objects from 2.1
// 3. Return inserted id to client.
// Side effect: Delete image record with ids (async)
// Done.

// Some APIs have side effect (async method/job). We have to design a job can configurable (timeout, retry count
// and time), support concurrent and maintainable.
// TODO: how to implement

// [url img, .....] (100) => save local storage
// [job, ...] (100)
// Who control ? => Group
// [[j1,j2,j3], [j5,j6]] => [j1,j2,j3] serial, [j5,j6] concurrent

// Caching: Only owner resource can update/invalid cache

// Socket Connection (struct - lib provide)
// => App Socket Connection (current user, permission)

func checkClosure() {
	arr := make([]func(), 10)

	for i := 0; i <= 9; i++ {
		f := func(y int) func() {
			// y is value of i
			return func() {
				log.Println(y) // pointer to y, because closure capture all variable outside as a pointer
			}
		}

		arr[i] = f(i + 2)
	}

	for i := range arr {
		arr[i]()
	}
}
