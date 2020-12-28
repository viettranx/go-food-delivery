package appctx

import (
	"fooddlv/common"
	"fooddlv/pubsub"
	"fooddlv/pubsub/pblocal"
	"gorm.io/gorm"
)

type appContext struct {
	db *gorm.DB
	ps pubsub.Pubsub
	rt common.RealtimeEngine
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

func (ctx *appContext) RealtimeEngine() common.RealtimeEngine { return ctx.rt }
