package main

import "fmt"

var ch chan int = make(chan int)

func main()  {

	go funA() //写到下面会死锁
	ch <- 1
	fmt.Println("hello")
}

func funA()  {

	a := <- ch
	println(a)

}



