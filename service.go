package service

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ServiceRunner interface {
	Run(context.Context)
}

type ServiceFunc func(context.Context)

func (fn ServiceFunc) Run(ctx context.Context) {
	fn(ctx)
}

func Run(ctx context.Context, runners ...ServiceRunner) {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for _, runner := range runners {
		wg.Add(1)
		go func(runner ServiceRunner) {
			defer wg.Done()
			runner.Run(ctx)
		}(runner)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	wg.Wait()
}
