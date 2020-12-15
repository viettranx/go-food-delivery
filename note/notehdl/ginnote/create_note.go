package ginnote

import (
	"context"
	"errors"
	"fooddlv/common"
	"fooddlv/note/notemodel"
	"fooddlv/note/noterepo"
	"fooddlv/note/notestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type fakeCreateNoteStore struct{}

func (fakeCreateNoteStore) Create(ctx context.Context, data *notemodel.NoteCreate) error {
	data.Id = 10
	return nil
}

type fakeImageStore struct{}

func (fakeImageStore) GetImages(ctx context.Context, cond map[string]interface{}, ids []int) ([]common.Image, error) {
	if len(ids) == 0 {
		return nil, errors.New("image ids can not be empty")
	}

	return []common.Image{
		{
			Id:     1,
			Url:    "https://",
			Width:  100,
			Height: 100,
		},
		{
			Id:     2,
			Url:    "https://",
			Width:  200,
			Height: 200,
		},
	}, nil
}

func (fakeImageStore) DeleteImages(ctx context.Context, ids []int) error {
	return nil
}

func CreateNote(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.NoteCreate

		if err := c.ShouldBind(&data); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := notestorage.NewMysqlStore(db)
		repo := noterepo.NewCreateNoteRepo(&fakeImageStore{}, store, appCtx.GetPubsub())
		if err := repo.CreateNote(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		//repo := noterepo.NewDeleteNoteRepo(store)
		//
		//_, err = repo.DeleteNote(c.Request.Context(), id)
		//
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, err)
		//	return
		//}
		//
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
