package main

import (
	"fmt"
	"time"
)

func cal2(a int , b int ,Exitchan chan bool)  {
	c := a+b
	fmt.Printf("%d + %d = %d\n",a,b,c)
	time.Sleep(time.Second*2)
	Exitchan <- true
	fmt.Printf("write chan %d\n", a)
}

func main() {

	Exitchan := make(chan bool,10)  //声明并分配管道内存
	for i :=0 ; i<10 ;i++{
		go cal2(i,i+1,Exitchan)
	}
	for j :=0; j<10; j++{
		<- Exitchan  //取信号数据，如果取不到则会阻塞
		fmt.Printf("read chan %d\n", j)
	}
	close(Exitchan) // 关闭管道
}