package ginnote

import (
	"fooddlv/common"
	"fooddlv/note/noterepo"
	"fooddlv/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DELETE v1/notes/:note-id

func DeleteNote(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("note-id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()

		store := notestorage.NewMysqlStore(db)
		repo := noterepo.NewDeleteNoteRepo(store)

		_, err = repo.DeleteNote(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
