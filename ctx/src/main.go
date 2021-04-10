package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	minLatency  = 10
	maxLatency  = 5000
	maxDuration = 2500
)

func main() {
	rootCtx := context.Background()

	for i := 0; i < 50; i++ {
		invokeCancellableGetData(rootCtx)
	}

}

func invokeCancellableGetData(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(maxDuration)*time.Millisecond)
	defer cancel()
	fmt.Println(getDataCtxAware(ctx))
}

func getDataCtxAware(ctx context.Context) (string, error) {
	res := make(chan string)
	go func() {
		res <- getData()
		close(res)
	}()

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
