package main

import (
	"context"
	"fmt"
	"time"
)

var msg chan int

func main()  {

	msg = make(chan int , 10)
	for i := 0; i<10 ; i++  {
		msg <- i
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)

	go subRoutine(ctx)

	defer cancel()//可以控制 子 goroutine 的结束
	defer close(msg)

	select {
	case <- ctx.Done():
		fmt.Println("main done" , ctx.Err())
	}

}

func subRoutine(ctx context.Context)  {

	ticker := time.NewTicker(1*time.Second)
	for {
		select {
		case <- ticker.C:
			fmt.Println("consumer message: ", <-msg )
		case <- ctx.Done():
			fmt.Println("sub routine done" , ctx.Err())
		}
	}

}

