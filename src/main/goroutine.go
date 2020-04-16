package main

import (
	"fmt"
	"runtime"
	"time"
)

func cal(a int , b int )  {
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
}

func main() {

	runtime.GOMAXPROCS(1)

	for i :=0 ; i<9 ;i++{
		go cal(i,i+1)  //启动10个goroutine 来计算
	}
	time.Sleep(time.Second * 2) // sleep作用是为了等待所有任务完成
}
