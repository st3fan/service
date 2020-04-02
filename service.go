package service

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ServiceFunc func(context.Context)

func Run(ctx context.Context, fn ServiceFunc) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer wg.Done()
		fn(ctx)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	wg.Wait()
}
