package common

const maxQueue = 10000

type Message struct {
	Name string
	Data interface{}
}

type jobQueue struct {
	queue chan Message
}

func NewJobQueue() *jobQueue {
	jq := jobQueue{
		queue: make(chan Message, maxQueue),
	}

	return &jq
}

func (jq *jobQueue) Emit(msg Message) {
	go func() { jq.queue <- msg }()
}

func (jq *jobQueue) Listen() <-chan Message {
	return jq.queue
}
