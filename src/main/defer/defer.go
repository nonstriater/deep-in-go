package main

import (
	"fmt"
)

func main() {
	defer_call()
	defer_call2()
	defer_embed()

	func_b_0()
	func_b_1()
	func_b_2()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	//panic("触发异常")
}

func defer_embed()  {
	defer func() {
		fmt.Println("1")
		defer func(){
			fmt.Println("2")
		}()
	}()
}

func defer_call2()  {
	defer func() {
		if ok := recover(); ok != nil {
			fmt.Println("recover," + ok.(string))//recover 返回值就是 panic()传入的参数
		}
	}()
	panic("error")
}


func func_b_0() {
	fmt.Println("func_b_0...top")
	a := 5
	defer fmt.Println("a=", a)//打印值: 5, 原因: 在a := 5之后, 这句话就已经编译好, 只是放到最后才执行, 所以为5
	a++
}
func func_b_1() {
	fmt.Println("func_b_1...top")
	a := 5
	defer func() {
		fmt.Println("a=", a) //打印值: 6, 原因: 在a := 5之后 , 此句就已经编译好, 但没有执行, 放到最后才执行, 最后执行时才传入当时的参数6
	}()
	a++
}
func func_b_2() {
	fmt.Println("func_b_2...top")
	a := 5
	defer func(a int) {
		fmt.Println("a=", a)//打印值: 5, 原因: 在a := 5之后, 此句话就已经编译好, 并传入了当时的参数5, 只是放到了最后才执行
	}(a)
	a++
}


