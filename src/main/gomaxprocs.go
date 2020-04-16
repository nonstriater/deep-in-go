package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(1)
	go func(){
		fmt.Println("hello world")
		// panic("hello world")  // 强制观察输出
	}()
	go func(){
		for {
			//fmt.Println("aaa")  // 非内联函数，这行注释打开，将导致 hello world 的输出
		}
	}()

	select {}
}

