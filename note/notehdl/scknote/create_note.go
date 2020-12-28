package scknote

import (
	"context"
	"fooddlv/common"
	"fooddlv/note/notemodel"
	"fooddlv/note/noterepo"
	"fooddlv/note/notestorage"
	socketio "github.com/googollee/go-socket.io"
)

// Admin delete a note (REST API)
// that note belongs user id 5
// the system should notify realtime to that user

func CreateNote(appCtx common.AppContext) func(socketio.Conn, notemodel.NoteCreate) {
	return func(conn socketio.Conn, data notemodel.NoteCreate) {
		db := appCtx.GetDBConnection()
		store := notestorage.NewMysqlStore(db)
		repo := noterepo.NewCreateNoteRepo(nil, store, appCtx.GetPubsub())
		if err := repo.CreateNote(context.Background(), &data); err != nil {
			panic(err)
		}
	}
}
