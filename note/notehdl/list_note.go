package notehdl

import (
	"fooddlv/common"
	"fooddlv/note/notemodel"
	"fooddlv/note/notestorage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type NoteStorage interface {
	List() ([]notemodel.Note, error)
}

func ListNote(appCtx AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()

		store := notestorage.NewMysqlStore(db)
		notes, err := store.List()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(notes))
	}
}
