package appctx

import socketio "github.com/googollee/go-socket.io"

type appSocket struct {
	UserId     int
	Connection socketio.Conn
}
