package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(500 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(),d)

	defer cancelFunc()

	select {
	case <-ctx.Done():
		fmt.Println("ctx done")
	case <-time.After(1 * time.Second):
		fmt.Println("after 1 second")
	}
}
