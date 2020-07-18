package main

import (
	"context"
	"fmt"
)

func sleepRandomCtx(ctx context.Context, ch chan bool)  {
	defer func() {
		fmt.Println("func sleepRandomCtx finish")
	}()

	select {
	case <- ctx.Done():
		fmt.Println("time to done")

	}
}

func sleepRandom(ch chan bool)  {

}

func main() {
	ctx := context.Background()
	_,cancelFunc := context.WithCancel(ctx)

	defer func() {
		fmt.Println("main func finish")
		cancelFunc()
	}()

	go func() {
		sleepRandom(nil)
		cancelFunc()

	}()

	//do work



}
