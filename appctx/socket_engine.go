package appctx

type AppSocket interface {
	Emit(event string, msg interface{}) error
}

type socketEngine struct {
	db map[int][]AppSocket
}

func NewSocketEngine() *socketEngine {
	sckEngine := &socketEngine{
		db: make(map[int][]AppSocket),
	}

	return sckEngine
}

func (se *socketEngine) AddSocket(userId int, s AppSocket) {
	if v, ok := se.db[userId]; ok {
		se.db[userId] = append(v, s)
		return
	}

	se.db[userId] = []AppSocket{s}
}
