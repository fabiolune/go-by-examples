package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	minLatency  = 10
	maxLatency  = 5000
	maxDuration = 2500
)

func main() {
	rootCtx := context.Background()
	var cancel context.CancelFunc

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("gracefully aborting due to interrupt...")
		cancel()
		os.Exit(0)
	}()

	for i := 0; i < 50; i++ {
		cancel = invokeCancellableGetData(rootCtx)
	}
}

func invokeCancellableGetData(ctx context.Context) context.CancelFunc {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(maxDuration)*time.Millisecond)
	defer cancel()
	fmt.Println(getDataCtxAware(ctx))
	return cancel
}

func getDataCtxAware(ctx context.Context) (string, error) {
	res := make(chan string)
	go func() {
		res <- getData()
		close(res)
	}()

	// The for loop runs until one of the two events is detected and the function returns
	for {
		select {
		case dst := <-res:
			return dst, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// slow function simultaing some IO operation
func getData() string {
	rand.Seed(time.Now().Unix())
	latency := minLatency + rand.Intn(maxLatency-minLatency)
	time.Sleep(time.Millisecond * time.Duration(latency))
	return fmt.Sprintf("result with latency %v", latency)
}
