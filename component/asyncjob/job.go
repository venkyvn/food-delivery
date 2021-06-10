package asyncjob

import (
	"context"
	"time"
)

// Job requirements
// 1. Job can do something (handler)
// 2. Job can retry (config retries time and duration)
// 3. Job should have state
// 4. Job should be managed by job manager
type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

type JobState int

func (j JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[j]
}

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

const (
	defaultMaxTimeout = time.Second * 10
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type JobHandler func(ctx context.Context) error

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func (j *job) RetryIndex() int {
	return j.retryIndex
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handler(ctx)

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

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times
}

func NewJob(handler JobHandler) *job {
	return &job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}
