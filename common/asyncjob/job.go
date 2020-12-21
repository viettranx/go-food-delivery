package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

// Some APIs have side effect (async method/job). We have to design a job can configurable (timeout, retry count
// and time), support concurrent and maintainable.

const (
	defaultMaxTimeout    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed"}[js]
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
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

	//if err != nil {
	//	j.retryIndex = 0
	//
	//	for j.retryIndex < len(j.config.Retries) {
	//
	//		time.Sleep(time.Second * time.Duration(j.config.Retries[j.retryIndex]))
	//		err = j.handler(ctx)
	//
	//		if err == nil {
	//			break
	//		}
	//
	//		j.retryIndex += 1
	//	}
	//}

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	j.state = StateRunning
	err := j.Execute(ctx)

	if err == nil {
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}

//
//func (j *job) Stop() {
//
//}
