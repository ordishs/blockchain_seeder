package main

import (
	"blockchain_seeder/blocklistener"
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for interupt signals
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		cancel()
	}()

	bl := blocklistener.NewBlockListener()

	if err := bl.Start(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		panic(err)
	}

	<-ctx.Done()
}
