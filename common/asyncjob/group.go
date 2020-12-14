package asyncjob

import (
	"context"
	"sync"
)

type group struct {
	isParallel bool
	jobs       []*job
	wg         *sync.WaitGroup
}

func NewGroup(isParallel bool, jobs ...*job) *group {
	g := &group{
		isParallel: isParallel,
		jobs:       jobs,
		wg:         new(sync.WaitGroup),
	}

	return g
}

func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))

	errChan := make(chan error, len(g.jobs))

	for i, _ := range g.jobs {
		if g.isParallel {
			go func(aj *job) {
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])

			continue
		}

		errChan <- g.runJob(ctx, g.jobs[i])
		g.wg.Done()
	}

	var err error

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}

	g.wg.Wait()
	return err
}

// Retry if needed
func (g *group) runJob(ctx context.Context, j *job) error {
	if err := j.Execute(ctx); err != nil {
		for {
			if j.State() == StateRetryFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}
