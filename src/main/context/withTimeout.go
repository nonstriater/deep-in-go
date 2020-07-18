package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	go handle(ctx, 5000 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("fiish main", ctx.Err())

	}

}

func handle(ctx context.Context, timeout time.Duration)  {

	//do something

	select {
	case <- ctx.Done():
		fmt.Println("handle done" , ctx.Err())
	case <- time.After(timeout):
		fmt.Println("handle timeout with", timeout)
	}

}