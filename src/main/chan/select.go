package main

import (
	"fmt"
	"runtime"
)

func  main()  {
	select1()
}

func select1()  {
	runtime.GOMAXPROCS(10)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 100
	string_chan <- "hello"

	select {
	case value := <-string_chan:
		//fmt.Println("string chan")
		panic(value)
	case value := <-int_chan:
		fmt.Println(value)
	}
}
