package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	data := "abc";
	longRunningThingManager(ctx, data)
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err error
	}

	ch := make(chan wrapper, 1)

	go func() {
		result, err := longRunningThing(ctx, data)
		ch <- wrapper{result, err}
	}()

	select {
	case data := <- ch:
		return data.result, data.err
        case <-ctx.Done():
		return "", ctx.Err()
	}
}

func longRunningThing(ctx context.Context, data string) (string, error) {
	fmt.Println("2秒Sleep")
	time.Sleep(2*time.Second)
	return data, nil
}
