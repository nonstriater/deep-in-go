package main

import "fmt"

var ch chan int = make(chan int)

func main()  {
	//test111()
	test123()
}

func test111()  {
	go funA() //写到下面会死锁
	ch <- 1
	fmt.Println("hello")
}

func test123()  {
	go funA()
	closeCh()
}

func funA()  {
	a := <- ch
	fmt.Printf("exit %d\n" , a)
}

func closeCh(){
	close(ch)
}

