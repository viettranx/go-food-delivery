package appctx

import (
	"fooddlv/pubsub"
	"fooddlv/pubsub/pblocal"
	"gorm.io/gorm"
)

type appContext struct {
	db *gorm.DB
	ps pubsub.Pubsub
	rt RealtimeEngine
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db, ps: pblocal.NewPubSub()}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	return ctx.db.Session(&gorm.Session{WithConditions: false})
}

func (ctx *appContext) GetPubsub() pubsub.Pubsub {
	return ctx.ps
}
