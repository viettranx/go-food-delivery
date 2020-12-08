package common

import (
	"context"
	"time"
)

// Some APIs have side effect (async method/job). We have to design a job can configurable (timeout, retry count
// and time), support concurrent and maintainable.

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []int{1, 5, 15}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []int
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed"}[js]
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	errChan    chan error
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		errChan:    make(chan error, 1),
		stopChan:   make(chan bool),
	}

	return &j
}

//
//func (j *job) Execute() error {
//	return j.handler(context.Background())
//}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning
	var err error

	err = j.handler(ctx)

	if err != nil {
		j.retryIndex = 0

		for j.retryIndex < len(j.config.Retries) {

			time.Sleep(time.Second * time.Duration(j.config.Retries[j.retryIndex]))
			err = j.handler(ctx)

			if err == nil {
				break
			}

			j.retryIndex += 1
		}
	}

	if err != nil {
		j.state = StateFailed
		j.errChan <- err
		return err
	}

	j.state = StateCompleted
	j.errChan <- nil

	return nil
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }
func (j *job) GetError() error { return <-j.errChan }

//
//func (j *job) Stop() {
//
//}
